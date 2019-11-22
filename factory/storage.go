package factory

import (
	"errors"
	"os"
	"strconv"
	"time"

	"github.com/zhangzitao/goauthes/pkg"
	"github.com/zhangzitao/goauthes/pkg/authorize"
	"github.com/zhangzitao/goauthes/pkg/storage"
	"github.com/zhangzitao/goauthes/pkg/token"
)

// GetRefreshStorage is
func GetRefreshStorage(refreshTokey string) (sto pkg.Storage, err error) {
	switch os.Getenv("GOAUTHES_STORAGE_TYPE") {
	case "Memory":
		sto, err = memoryGetRefreshStorage(refreshTokey)
	}
	return sto, err
}

func memoryGetRefreshStorage(refreshTokey string) (sto pkg.Storage, err error) {
	if arr, ok := storage.MemoryDB["goauthes:refresh:"+refreshTokey]; ok {
		sto = &arr[0]
		return sto, nil
	}
	return sto, errors.New("get refresh storage failed")
}

// GenerateFromAuthorize save the data to storage format
func GenerateFromAuthorize(class string, auth pkg.Authorize) (sto pkg.Storage, err error) {
	expireTime, err := strconv.Atoi(os.Getenv("GOAUTHES_TOKEN_EXPIRE"))
	newToken, err := GenerateToken(os.Getenv("GOAUTHES_TOKEN_TYPE"), int64(expireTime))
	sto, err = GenerateStorage(class, newToken, auth)
	return sto, err
}

// GenerateStorage is a facotry produce storage
func GenerateStorage(class string, tok token.Token, auth pkg.Authorize) (sto pkg.Storage, err error) {
	var base pkg.BaseStorage
	switch class {
	case "PassWord":
		au := auth.(*authorize.PassWord)
		base = pkg.BaseStorage{
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
	}
	switch os.Getenv("GOAUTHES_STORAGE_TYPE") {
	case "Memory":
		sto = &storage.Memory{Data: base}
	}

	return sto, err
}

// MakeNewStorageByOld is pick the auth data
func MakeNewStorageByOld(s pkg.Storage) (m pkg.Storage, err error) {
	expireTime, err := strconv.Atoi(os.Getenv("GOAUTHES_TOKEN_EXPIRE"))
	tok, err := GenerateToken("Bearer", int64(expireTime))
	if err != nil {
		return m, err
	}
	data := s.ToBase()
	base := pkg.BaseStorage{
		ClientID:         data.ClientID,
		UserID:           data.UserID,
		RedirectURI:      data.RedirectURI,
		Scope:            data.Scope,
		Code:             data.Code,
		CodeCreateAt:     time.Now(),
		CodeExpiresIn:    data.CodeExpiresIn,
		Access:           tok.AccessToken,
		AccessCreateAt:   time.Now(),
		AccessExpiresIn:  tok.ExpiresIn,
		Refresh:          tok.RefreshToken,
		RefreshCreateAt:  time.Now(),
		RefreshExpiresIn: data.RefreshExpiresIn,
		TokenType:        tok.TokenType,
	}
	switch os.Getenv("GOAUTHES_STORAGE_TYPE") {
	case "Memory":
		m = &storage.Memory{Data: base}
	}
	return m, nil
}
