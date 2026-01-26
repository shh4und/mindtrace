package servicos

import (
	"crypto/rand"
	"encoding/hex"
)

// GenerateSecureToken cria um token aleatório de 32 bytes (64 caracteres hex)
func GenerateSecureToken() (string, error) {
	b := make([]byte, 32)
	_, err := rand.Read(b)
	if err != nil {
		return "", err
	}
	return hex.EncodeToString(b), nil
}
