package arvancloud

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

var (
	// HTTP request multiplexer
	mux *http.ServeMux

	// API client for test
	client *API

	// HTTP server used for mocking
	server *httptest.Server
)

func setup(opts ...Option) {
	// Create a test server
	mux = http.NewServeMux()
	server = httptest.NewServer(mux)

	// Disable rate limits and retries for test purpose
	opts = append([]Option{UsingRateLimit(100000), UsingRetryPolicy(0, 0, 0)}, opts...)

	// Configure client
	client, _ = New("deadbeef", opts...)
	client.BaseURL = server.URL
}

func teardown() {
	server.Close()
}

func TestHeaders(t *testing.T) {
	// Default
	setup()
	mux.HandleFunc("/domains/"+testDomain+"/dns-records", func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, http.MethodGet, r.Method, "Expected method 'GET', got %s", r.Method)
		assert.Equal(t, "application/json", r.Header.Get("Content-Type"))
	})
	teardown()

	// Override defaults
	headers := make(http.Header)
	headers.Set("Content-Type", "application/graphql")
	headers.Add("H1", "V1")
	headers.Add("H2", "V2")
	setup(Headers(headers))
	mux.HandleFunc("/domains/"+testDomain+"/dns-records", func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, http.MethodGet, r.Method, "Expected method 'GET', got %s", r.Method)
		assert.Equal(t, "application/graphql", r.Header.Get("Content-Type"))
		assert.Equal(t, "V1", r.Header.Get("H1"))
		assert.Equal(t, "V2", r.Header.Get("H2"))
	})
	teardown()

	// Authentication
	setup()
	client, err := New("TEST_TOKEN")
	assert.NoError(t, err)
	client.BaseURL = server.URL
	mux.HandleFunc("/domains/"+testDomain+"/dns-records", func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, http.MethodGet, r.Method, "Expected method 'GET', got %s", r.Method)
		assert.Equal(t, "Bearer TEST_TOKEN", r.Header.Get("Authorization"))
		assert.Equal(t, "application/json", r.Header.Get("Content-Type"))
	})
	teardown()
}
