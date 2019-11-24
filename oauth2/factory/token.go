package factory

import (
	pure "github.com/zhangzitao/goauthes/oauth2/structs_pure"
	"github.com/zhangzitao/goauthes/oauth2/interfaces"

	"github.com/google/uuid"
)

// GenerateToken is factory function
func GenerateToken(tokenType string, duration int64) (tok pure.Token, err error) {
	switch tokenType {
	case "Bearer":
		access := uuid.New().String()
		refresh := uuid.New().String()
		tok = pure.Token{
			AccessToken:  access,
			TokenType:    "bearer",
			ExpiresIn:    duration,
			RefreshToken: refresh,
		}
	}

	return tok, nil
}

// GenerateTokenFromStorage is pick the token data
func GenerateTokenFromStorage(s interfaces.Storage) (pure.Token, error) {
	tok := pure.Token{
		AccessToken:  s.ToBase().Access,
		TokenType:    s.ToBase().TokenType,
		ExpiresIn:    s.ToBase().AccessExpiresIn,
		RefreshToken: s.ToBase().Refresh,
	}
	return tok, nil
}
