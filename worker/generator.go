package worker

import (
	"crypto/rand"
	"encoding/hex"
)

func GenerateValue(size int) (string, error) {
	value := make([]byte, size)
	_, err := rand.Read(value)
	if err != nil {
		return "", err
	}
	return hex.EncodeToString(value), nil
}
