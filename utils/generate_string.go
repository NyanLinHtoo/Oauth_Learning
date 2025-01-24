package utils

import (
	"crypto/rand"
	"encoding/base64"
)

// GenerateRandomState generates a secure random string to be used as a state parameter
func GenerateRandomState(length int) (string, error) {
	bytes := make([]byte, length)
	_, err := rand.Read(bytes)
	if err != nil {
		return "", err
	}
	return base64.URLEncoding.EncodeToString(bytes), nil
}
