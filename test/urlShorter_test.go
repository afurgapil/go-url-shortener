package helpers

import (
	"strings"
	"testing"

	"github.com/afurgapil/go-url-shortener/pkg/helpers"
)


func TestURLShorter(t *testing.T) {
    longURL := "https://www.example.com/test"
    shortURL := helpers.URLShorter(longURL)
	println(shortURL)

    if !strings.HasPrefix(shortURL, "https://example.com/") {
        t.Errorf("Expected shortURL to start with 'https://example.com/', got %s", shortURL)
    }

    lastIndex := strings.LastIndex(shortURL, "/")
    randomString := shortURL[lastIndex+1:]
    if len(randomString) != 6 {
        t.Errorf("Expected random string length to be 6, got %d", len(randomString))
    }
}




