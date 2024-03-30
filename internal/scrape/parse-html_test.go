package scrape

import (
	"testing"
)

func TestParseHTML(t *testing.T) {
	htmlContent := `<html><body>Hello, <b>World!</b></body></html>`
	expected := "Hello, World! "
	result, err := ParseHTML(htmlContent)
	if err != nil {
		t.Fatalf("extracTextFromHTML returned an error %v", err)
	}
	if result != expected {
		t.Errorf("Expected %q, got %q", expected, result)
	}
}
