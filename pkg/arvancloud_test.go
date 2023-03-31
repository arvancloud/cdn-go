package arvancloud

import (
	"net/http"
	"net/http/httptest"
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
