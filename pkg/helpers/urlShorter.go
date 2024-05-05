package helpers

import (
	"math/rand"
	"time"
)
func URLShorter(longURL string) string {
	randomString := generateRandomString(6)

	shortURL := "https://example.com/" + randomString

	return shortURL
}

func generateRandomString(length int) string {
	const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	var seededRand *rand.Rand = rand.New(rand.NewSource(time.Now().UnixNano()))

	randomString := make([]byte, length)
	for i := range randomString {
		randomString[i] = charset[seededRand.Intn(len(charset))]
	}

	return string(randomString)
}