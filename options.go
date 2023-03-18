package arvancloud

import (
	"net/http"

	"time"

	"golang.org/x/time/rate"
)

type Option func(*API) error

func Headers(headers http.Header) Option {
	return func(api *API) error {
		api.headers = headers
		return nil
	}
}

func UsingRateLimit(rps float64) Option {
	return func(api *API) error {
		api.rateLimiter = rate.NewLimiter(rate.Limit(rps), 1)
		return nil
	}
}

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

func UserAgent(userAgent string) Option {
	return func(api *API) error {
		api.UserAgent = userAgent

		return nil
	}
}

func Debug(debug bool) Option {
	return func(api *API) error {
		api.Debug = debug

		return nil
	}
}

func (api *API) parseOptions(options ...Option) error {
	for _, option := range options {
		err := option(api)

		if err != nil {
			return err
		}
	}

	return nil
}
