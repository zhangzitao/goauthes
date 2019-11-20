package token

import (
	"github.com/google/uuid"
)

// GenerateToken is factory function
func GenerateToken(tokenType string, duration int64) (tok Token, err error) {
	switch tokenType {
	case "Bearer":
		access := uuid.New().String()
		refresh := uuid.New().String()
		tok = Token{
			AccessToken:  access,
			TokenType:    "bearer",
			ExpiresIn:    duration,
			RefreshToken: refresh,
		}
	}

	return tok, nil
}
