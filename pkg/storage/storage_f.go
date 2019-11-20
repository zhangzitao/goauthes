package storage

import (
	"time"

	"github.com/zhangzitao/goauthes/pkg/authorize"
	"github.com/zhangzitao/goauthes/pkg/token"
)

// GenerateStorage is a facotry produce storage
func GenerateStorage(class string, tok token.Token, auth authorize.Authorize) (sto Storage, err error) {
	switch class {
	case "Memory":
		au := auth.(*authorize.PassWord)
		base := Base{
			ClientID:         "",
			UserID:           au.UserID,
			RedirectURI:      "",
			Scope:            au.Scope,
			Code:             "",
			CodeCreateAt:     time.Now(),
			CodeExpiresIn:    0,
			Access:           tok.AccessToken,
			AccessCreateAt:   time.Now(),
			AccessExpiresIn:  tok.ExpiresIn,
			Refresh:          tok.RefreshToken,
			RefreshCreateAt:  time.Now(),
			RefreshExpiresIn: 0,
			TokenType:        tok.TokenType,
		}
		sto = &Memory{Data: base}
	}

	return sto, err
}
