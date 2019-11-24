package factory

import (
	"errors"
	"os"
	"strconv"
	"time"

	"github.com/zhangzitao/goauthes/oauth2/implementation/storage"
	"github.com/zhangzitao/goauthes/oauth2/interfaces"
	base "github.com/zhangzitao/goauthes/oauth2/structs_base"
	pure "github.com/zhangzitao/goauthes/oauth2/structs_pure"
)

// GetRefreshStorage is
func GetRefreshStorage(refreshTokey string) (sto interfaces.Storage, err error) {
	switch os.Getenv("GOAUTHES_STORAGE_TYPE") {
	case "Memory":
		sto, err = memoryGetRefreshStorage(refreshTokey)
	}
	return sto, err
}

func memoryGetRefreshStorage(refreshTokey string) (sto interfaces.Storage, err error) {
	if arr, ok := storage.MemoryDB["goauthes:refresh:"+refreshTokey]; ok {
		sto = &arr[0]
		return sto, nil
	}
	return sto, errors.New("get refresh storage failed")
}

// GenerateFromAuthorize save the data to storage format
func GenerateFromAuthorize(auth interfaces.Authorize) (sto interfaces.Storage, err error) {
	expireTime, err := strconv.Atoi(os.Getenv("GOAUTHES_TOKEN_EXPIRE"))
	newToken, err := GenerateToken(os.Getenv("GOAUTHES_TOKEN_TYPE"), int64(expireTime))
	sto, err = GenerateStorage(newToken, auth)
	return sto, err
}

// GenerateStorage is a facotry produce storage
func GenerateStorage(tok pure.Token, auth interfaces.Authorize) (sto interfaces.Storage, err error) {
	var data base.StorageData
	switch auth.GetType() {
	case "PassWord":
		au, _ := auth.ToBasePassWord()
		data = base.StorageData{
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
		sto = &storage.Memory{Data: data}
	}

	return sto, err
}

// MakeNewStorageByOld is pick the auth data
func MakeNewStorageByOld(s interfaces.Storage) (m interfaces.Storage, err error) {
	expireTime, err := strconv.Atoi(os.Getenv("GOAUTHES_TOKEN_EXPIRE"))
	tok, err := GenerateToken("Bearer", int64(expireTime))
	if err != nil {
		return m, err
	}
	data := s.ToBase()
	base := base.StorageData{
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
