package auth

import (
	"crypto/rand"
	"fmt"
)

type TokenManager struct{}

func (TokenManager) generate_token() string {
	token := make([]byte, 20)
	rand.Read(token)
	return fmt.Sprintf("%x", token)
}
