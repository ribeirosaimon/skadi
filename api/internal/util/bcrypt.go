package util

import (
	"encoding/hex"
	"math/rand"

	"golang.org/x/crypto/bcrypt"
)

func EncryptPassword(pass string) (string, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(pass), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(hash), nil
}

func CompareEncryptPassword(pass, encryptPass string) bool {
	if err := bcrypt.CompareHashAndPassword([]byte(pass), []byte(encryptPass)); err != nil {
		return false
	}
	return true
}

func GenerateUUIDToken() string {
	uuidBytes := make([]byte, 32)
	_, err := rand.Read(uuidBytes)
	if err != nil {
		return ""
	}
	return hex.EncodeToString(uuidBytes)
}
