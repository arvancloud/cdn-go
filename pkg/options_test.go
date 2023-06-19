package arvancloud

import (
	"net/http"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"golang.org/x/time/rate"
)

func TestParseOptions(t *testing.T) {
	headers := make(http.Header)
	headers.Add("H1", "V1")
	headers.Add("H2", "V2")
	headers.Add("H3", "V3")

	api, err := New(
		"TEST_TOKEN",
		Debug(true),
		UserAgent("UA-test"),
		Headers(headers),
		UsingRateLimit(1.5),
		UsingRetryPolicy(1, 2, 3),
	)

	if err != nil {
		t.Errorf("Error ocurred: %v", err.Error())
	}

	assert.Equal(t, true, api.Debug)

	assert.Equal(t, "UA-test", api.UserAgent)

	assert.Equal(t, "V1", api.headers.Get("H1"))
	assert.Equal(t, "V2", api.headers.Get("H2"))
	assert.Equal(t, "V3", api.headers.Get("H3"))

	assert.Equal(t, 1, api.rateLimiter.Burst())
	assert.Equal(t, rate.Limit(1.5), api.rateLimiter.Limit())

	assert.Equal(t, 1, api.retryPolicy.MaxRetries)
	assert.Equal(t, time.Duration(2000000000), api.retryPolicy.MinRetryDelay)
	assert.Equal(t, time.Duration(3000000000), api.retryPolicy.MaxRetryDelay)
}
