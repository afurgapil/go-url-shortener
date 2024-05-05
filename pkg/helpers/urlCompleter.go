package helpers

import (
	"log"

	"github.com/afurgapil/go-url-shortener/configs"
)

func UrlCompleter(shortURL string) (string, error) {
    urlTemplate, err := configs.ReturnURL()
    if err != nil {
        log.Println("Error loading config:", err)
        return "", err
    }
    completedURL := urlTemplate + shortURL
    return completedURL, nil
}
