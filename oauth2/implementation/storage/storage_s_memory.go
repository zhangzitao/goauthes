package storage

import (
	"os"
	"strconv"

	base "github.com/zhangzitao/goauthes/oauth2/structs_base"
)

// MemoryDB is memory store
var MemoryDB = map[string][]Memory{}

// Memory is a storage specific implemention
type Memory struct {
	Data base.StorageData
}

func makeUserSessionReady(key string) error {
	size, err := strconv.Atoi(os.Getenv("GOAUTHES_STORAGE_USER_SESSION_LENGTH"))
	if err != nil {
		return err
	}
	// Delete old storage
	if len(MemoryDB[key]) >= size {
		for i := 0; i <= len(MemoryDB[key])-size; i++ {
			_, err = MemoryDB[key][i].Delete()
			if err != nil {
				return err
			}
		}
		// Delete old session item
		MemoryDB[key] = MemoryDB[key][len(MemoryDB[key])-size+1:]
	}
	return nil
}
func insertUserSession(s *Memory) (err error) {
	key := "goauthes:userid:" + s.Data.UserID
	err = makeUserSessionReady(key)
	if err != nil {
		return err
	}
	// insert user session item
	MemoryDB[key] = append(MemoryDB[key], *s)
	// set access token item
	MemoryDB["goauthes:access:"+s.Data.Access] = []Memory{*s}
	// set refresh token item
	MemoryDB["goauthes:refresh:"+s.Data.Refresh] = []Memory{*s}
	return nil
}

// Create is to create data in db
func (s *Memory) Create() (bool, error) {
	err := insertUserSession(s)
	if err != nil {
		return false, err
	}
	// log.Printf("%v", MemoryDB)
	return true, nil
}

// Refresh is to delete and create data in db
func (s *Memory) Refresh() (bool, error) {
	// delete old keys
	// find storage in user session array
	key := "goauthes:userid:" + s.Data.UserID
	flag1 := -1
	for i, v := range MemoryDB[key] {
		if v.Data.Refresh == s.Data.Refresh {
			flag1 = i
			break
		}
	}
	if flag1 != -1 {
		MemoryDB[key] = append(MemoryDB[key][:flag1], MemoryDB[key][flag1+1:]...)
	}
	// find storage hit refresh
	if _, ok := MemoryDB["goauthes:refresh:"+s.Data.Refresh]; ok {
		delete(MemoryDB, "goauthes:refresh:"+s.Data.Refresh)
	}
	// find storage hit access
	if _, ok := MemoryDB["goauthes:access:"+s.Data.Access]; ok {
		delete(MemoryDB, "goauthes:access:"+s.Data.Access)
	}
	return true, nil
}

// Delete is to delete data in db
func (s *Memory) Delete() (bool, error) {
	delete(MemoryDB, "goauthes:access:"+s.Data.Access)
	delete(MemoryDB, "goauthes:refresh:"+s.Data.Refresh)
	return true, nil
}

// ToBase is
func (s *Memory) ToBase() base.StorageData {
	return s.Data
}
