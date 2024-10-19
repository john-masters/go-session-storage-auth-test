package utils

import (
	"crypto/rand"
	"crypto/sha256"
	"encoding/hex"
)

func GenerateRandomHash() (string, error) {
	b := make([]byte, 512)
	_, err := rand.Read(b)

	if err != nil {
		return "", err
	}

	hasher := sha256.New()
	hasher.Write(b)
	sha := hex.EncodeToString(hasher.Sum(nil))

	return sha, nil
}
