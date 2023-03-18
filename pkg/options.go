package arvancloud

import (
	"net/http"

	"time"

	"golang.org/x/time/rate"
)

type Option func(*API) error

// Headers will set custom HTTP headers for API calls
func Headers(headers http.Header) Option {
	return func(api *API) error {
		api.headers = headers
		return nil
	}
}

// UsingRateLimit will apply a rate limit policy to API client
func UsingRateLimit(rps float64) Option {
	return func(api *API) error {
		api.rateLimiter = rate.NewLimiter(rate.Limit(rps), 1)
		return nil
	}
}

// UsingRetryPolicy will apply a retry policy to API client
func UsingRetryPolicy(maxRetries int, minRetryDelaySecs int, maxRetryDelaySecs int) Option {
	return func(api *API) error {
		api.retryPolicy = RetryPolicy{
			MaxRetries:    maxRetries,
			MinRetryDelay: time.Duration(minRetryDelaySecs) * time.Second,
			MaxRetryDelay: time.Duration(maxRetryDelaySecs) * time.Second,
		}

		return nil
	}
}

// UsingLogger will apply a custom user agent to API client
func UserAgent(userAgent string) Option {
	return func(api *API) error {
		api.UserAgent = userAgent

		return nil
	}
}

// Debug will handle debugging mode for API client
func Debug(debug bool) Option {
	return func(api *API) error {
		api.Debug = debug

		return nil
	}
}

// parseOptions will parse the supplied options for API client
func (api *API) parseOptions(options ...Option) error {
	for _, option := range options {
		err := option(api)

		if err != nil {
			return err
		}
	}

	return nil
}
