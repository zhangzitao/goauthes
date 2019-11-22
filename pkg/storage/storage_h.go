package storage

import (
	"os"
	"strconv"
	"time"

	"github.com/zhangzitao/goauthes/pkg/authorize"
	"github.com/zhangzitao/goauthes/pkg/token"
)

// GenerateFromAuthorize save the data to storage format
func GenerateFromAuthorize(class string, auth authorize.Authorize) (sto Storage, err error) {
	expireTime, err := strconv.Atoi(os.Getenv("GOAUTHES_TOKEN_EXPIRE"))
	newToken, err := token.GenerateToken(os.Getenv("GOAUTHES_TOKEN_TYPE"), int64(expireTime))
	sto, err = GenerateStorage(class, newToken, auth)
	return sto, err
}

// GenerateStorage is a facotry produce storage
func GenerateStorage(class string, tok token.Token, auth authorize.Authorize) (sto Storage, err error) {
	switch class {
	case "PassWord":
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
