package storage

import (
	"errors"
	"os"
	"strconv"

	"github.com/zhangzitao/goauthes/pkg/authorize"
	"github.com/zhangzitao/goauthes/pkg/token"
)

// GenerateFromAuthorize save the data to storage format
func GenerateFromAuthorize(auth authorize.Authorize) (sto Storage, err error) {
	expireTime, err := strconv.Atoi(os.Getenv("GOAUTHES_TOKEN_EXPIRE"))
	newToken, err := token.GenerateToken(os.Getenv("GOAUTHES_TOKEN_TYPE"), int64(expireTime))
	sto, err = GenerateStorage(os.Getenv("GOAUTHES_STORAGE_TYPE"), newToken, auth)
	return sto, err
}

// GetRefreshStorage is
func GetRefreshStorage(refreshTokey string) (sto Storage, err error) {
	switch os.Getenv("GOAUTHES_STORAGE_TYPE") {
	case "Memory":
		sto, err = memoryGetRefreshStorage(refreshTokey)
	}
	return sto, err
}

func memoryGetRefreshStorage(refreshTokey string) (sto Storage, err error) {
	if arr, ok := MemoryDB["goauthes:refresh:"+refreshTokey]; ok {
		sto = &arr[0]
		return sto, nil
	}
	return sto, errors.New("get refresh storage failed")
}
