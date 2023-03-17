package arvancloud

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log"
	"math"
	"net/http"
	"net/http/httputil"
	"regexp"
	"strings"
	"time"

	"golang.org/x/time/rate"
)

type API struct {
	APIKey      string
	APIEmail    string
	BaseURL     string
	UserAgent   string
	headers     http.Header
	httpClient  *http.Client
	rateLimiter *rate.Limiter
	retryPolicy RetryPolicy
	logger      Logger
	Debug       bool
}

type RetryPolicy struct {
	MaxRetries    int
	MinRetryDelay time.Duration
	MaxRetryDelay time.Duration
}

type Logger interface {
	Printf(format string, v ...interface{})
}

func newClient(opts ...Option) (*API, error) {
	silentLogger := log.New(io.Discard, "", log.LstdFlags)

	api := &API{
		BaseURL:     fmt.Sprintf("%s://%s%s", SCHEME, HOST_NAME, BASE_PATH),
		UserAgent:   UA + "/" + VERSION,
		headers:     make(http.Header),
		rateLimiter: rate.NewLimiter(rate.Limit(4), 1),
		retryPolicy: RetryPolicy{
			MaxRetries:    3,
			MinRetryDelay: 1 * time.Second,
			MaxRetryDelay: 30 * time.Second,
		},
		logger: silentLogger,
	}

	err := api.parseOptions(opts...)
	if err != nil {
		return nil, fmt.Errorf("options parsing failed: %w", err)
	}

	api.httpClient = http.DefaultClient

	return api, nil
}

func New(token string, opts ...Option) (*API, error) {
	if token == "" {
		return nil, errors.New(errEmptyAPIToken)
	}

	api, err := newClient(opts...)
	if err != nil {
		return nil, err
	}

	api.APIKey = token

	return api, nil
}

func (api *API) makeRequestContext(ctx context.Context, method, uri string, record interface{}) (*DNSRecord_Response, error) {
	res, err := api.makeRequest(ctx, method, uri, record)

	if err != nil {
		return nil, err
	}

	return res, err
}

func (api *API) makeRequest(ctx context.Context, method, uri string, record interface{}) (*DNSRecord_Response, error) {
	var err error
	var resp *http.Response
	var respErr error
	var respBody []byte

	for i := 0; i <= api.retryPolicy.MaxRetries; i++ {
		var reqBody io.Reader

		if record != nil {
			if r, ok := record.(io.Reader); ok {
				reqBody = r
			} else if paramBytes, ok := record.([]byte); ok {
				reqBody = bytes.NewReader(paramBytes)
			} else {
				var jsonBody []byte
				jsonBody, err = json.Marshal(record)

				if err != nil {
					return nil, fmt.Errorf("error marshalling record data to JSON: %w", err)
				}

				reqBody = bytes.NewReader(jsonBody)
			}
		}

		if i > 0 {
			sleepDuration := time.Duration(math.Pow(2, float64(i-1)) * float64(api.retryPolicy.MinRetryDelay))

			if sleepDuration > api.retryPolicy.MaxRetryDelay {
				sleepDuration = api.retryPolicy.MaxRetryDelay
			}

			api.logger.Printf("Sleeping %s before retry attempt number %d for request %s %s", sleepDuration.String(), i, method, uri)

			select {
			case <-time.After(sleepDuration):
			case <-ctx.Done():
				return nil, fmt.Errorf("operation aborted during backoff: %w", ctx.Err())
			}
		}

		err = api.rateLimiter.Wait(ctx)
		if err != nil {
			return nil, fmt.Errorf("error caused by request rate limiting: %w", err)
		}

		resp, respErr = api.request(ctx, method, uri, reqBody)

		// short circuit processing on context timeouts
		if respErr != nil && errors.Is(respErr, context.DeadlineExceeded) {
			return nil, respErr
		}

		// retry if the server is rate limiting us or if it failed
		// assumes server operations are rolled back on failure
		if respErr != nil || resp.StatusCode == http.StatusTooManyRequests || resp.StatusCode >= 500 {
			if resp != nil && resp.StatusCode == http.StatusTooManyRequests {
				respErr = errors.New("exceeded available rate limit retries")
			}

			if respErr == nil {
				respErr = fmt.Errorf("received %s response (HTTP %d), please try again later", strings.ToLower(http.StatusText(resp.StatusCode)), resp.StatusCode)
			}
			continue
		} else {
			respBody, err = io.ReadAll(resp.Body)
			defer resp.Body.Close()
			if err != nil {
				return nil, fmt.Errorf("could not read response body: %w", err)
			}

			break
		}
	}

	// still had an error after all retries
	if respErr != nil {
		return nil, respErr
	}

	if resp.StatusCode >= http.StatusBadRequest {
		if resp.StatusCode >= http.StatusInternalServerError {

			return nil, &ServiceError{arvancloudError: &Error{
				Code:    resp.StatusCode,
				Message: errInternalServiceError,
			}}
		}

		errBody := &DNSRecord_Response{}
		err = json.Unmarshal(respBody, &errBody)
		if err != nil {

			return nil, fmt.Errorf(errUnmarshalErrorBody+": %w", err)
		}

		err := &Error{
			Code:    resp.StatusCode,
			Message: errBody.Message,
			Errors:  errBody.Errors,
			Type:    ErrorTypeRequest,
		}

		return nil, &RequestError{arvancloudError: err}
	}

	response := &DNSRecord_Response{}
	err = json.Unmarshal(respBody, &response)
	if err != nil {
		return nil, fmt.Errorf(errUnmarshalError+": %w", err)
	}

	return response, nil
}

func (api *API) request(ctx context.Context, method, uri string, reqBody io.Reader) (*http.Response, error) {
	req, err := http.NewRequestWithContext(ctx, method, api.BaseURL+uri, reqBody)
	if err != nil {
		return nil, fmt.Errorf("HTTP request creation failed: %w", err)
	}

	combinedHeaders := make(http.Header)
	copyHeader(combinedHeaders, api.headers)
	req.Header = combinedHeaders

	req.Header.Set("Authorization", api.APIKey)
	req.Header.Set("User-Agent", api.UserAgent)
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Accept", "application/json")

	if api.Debug {
		dump, err := httputil.DumpRequestOut(req, true)
		if err != nil {
			return nil, err
		}

		// Strip out sensitive information
		sensitiveKeys := []string{api.APIKey}
		for _, key := range sensitiveKeys {
			if key != "" {
				valueRegex := regexp.MustCompile(fmt.Sprintf("(?m)%s", key))
				dump = valueRegex.ReplaceAll(dump, []byte("[redacted]"))
			}
		}
		log.Printf("\n%s", string(dump))
	}

	resp, err := api.httpClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("HTTP request failed: %w", err)
	}

	if api.Debug {
		dump, err := httputil.DumpResponse(resp, true)
		if err != nil {
			return resp, err
		}
		log.Printf("\n%s", string(dump))
	}

	return resp, nil
}

func copyHeader(target, source http.Header) {
	for k, vs := range source {
		target[k] = vs
	}
}
