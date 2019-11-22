package storage

import (
	"errors"
	"os"
)

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
