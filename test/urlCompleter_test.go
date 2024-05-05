package helpers

import (
	"testing"
	"time"

	"github.com/afurgapil/go-url-shortener/pkg/helpers"
)

func TestUrlCompleter(t *testing.T) {
	shortURL := "example"
	expectedCompletedURL := "https://example.com/example"

	start := time.Now()
	completedURL, err := helpers.UrlCompleter(shortURL)
	duration := time.Since(start)

	if err != nil {
		t.Errorf("Got error: %v", err)
	}

	if completedURL != expectedCompletedURL {
		t.Errorf("Expected completed URL: %s, got: %s", expectedCompletedURL, completedURL)
	}

	t.Logf("Test completed in %s", duration)
}
