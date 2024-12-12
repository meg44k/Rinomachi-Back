package utils

import (
	"crypto/rand"
	"crypto/sha256"
	"encoding/hex"
)

func GenerateUserID() string {
	b := make([]byte, 32)
	_, err := rand.Read(b)
	if err != nil {
		return ""
	}
	hash := sha256.Sum256(b)
	return "U" + hex.EncodeToString(hash[:16])

}

func GenerateBuildingID() string {
	b := make([]byte, 32)
	_, err := rand.Read(b)
	if err != nil {
		return ""
	}
	hash := sha256.Sum256(b)
	return "B" + hex.EncodeToString(hash[:16])

}
