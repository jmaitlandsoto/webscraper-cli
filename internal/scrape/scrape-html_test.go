package scrape

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestScrapeHTML(t *testing.T) {
	// Start a local HTTP server for testing
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte(`<html><body>Hello, World!</body></html>`))
	}))
	defer server.Close() // Close the server when test finishes

	// Use the URL of the test server
	result, err := ScrapeHTML(server.URL)
	if err != nil {
		t.Fatalf("ScrapeHTML returned an error: %v", err)
	}
	fmt.Println(result)

	// expecting it to return the contents of the body tag
	expected := "Hello, World!"
	if result != expected {
		t.Errorf("Expected %q, got %q", expected, result)
	}

	// TODO: add some more test cases with more complex html inputs
}
