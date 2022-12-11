package utils

import (
	"crypto/rand"
	"encoding/hex"
)

// Creates a random hex token and returns it.
func RandomHexToken(length int) (string, error) {
	bytes := make([]byte, length)

	if _, err := rand.Read(bytes); err != nil {
		return "", err
	}

	return hex.EncodeToString(bytes), nil
}
