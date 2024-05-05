package helpers

import (
	"regexp"
	"testing"

	"github.com/afurgapil/go-url-shortener/pkg/helpers"
)

func TestPassKeyGenerator(t *testing.T) {
	length := 8
	passKey := helpers.PassKeyGenerator(length)

	if len(passKey) != length {
		t.Errorf("Expected pass key length: %d, got: %d", length, len(passKey))
	}

	match, err := regexp.MatchString("^[a-zA-Z0-9]*$", passKey)
	if err != nil {
		t.Errorf("Error matching pass key format: %v", err)
	}

	if !match {
		t.Errorf("Pass key contains invalid characters")
	}
}
