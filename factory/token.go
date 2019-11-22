package factory

import (
	"github.com/zhangzitao/goauthes/pkg"
	"github.com/zhangzitao/goauthes/pkg/token"

	"github.com/google/uuid"
)

// GenerateToken is factory function
func GenerateToken(tokenType string, duration int64) (tok token.Token, err error) {
	switch tokenType {
	case "Bearer":
		access := uuid.New().String()
		refresh := uuid.New().String()
		tok = token.Token{
			AccessToken:  access,
			TokenType:    "bearer",
			ExpiresIn:    duration,
			RefreshToken: refresh,
		}
	}

	return tok, nil
}

// GenerateTokenFromStorage is pick the token data
func GenerateTokenFromStorage(s pkg.Storage) (token.Token, error) {
	tok := token.Token{
		AccessToken:  s.ToBase().Access,
		TokenType:    s.ToBase().TokenType,
		ExpiresIn:    s.ToBase().AccessExpiresIn,
		RefreshToken: s.ToBase().Refresh,
	}
	return tok, nil
}
