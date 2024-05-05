package helpers

import (
	"math/rand"
	"time"
)

func PassKeyGenerator(length int) string {
	const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	var seededRand *rand.Rand = rand.New(rand.NewSource(time.Now().UnixNano()))

	passKey := make([]byte, length)
	for i := range passKey {
		passKey[i] = charset[seededRand.Intn(len(charset))]
	}

	return string(passKey)
}
