package storage

import (
	"log"
	"os"
	"strconv"
	"time"

	"github.com/zhangzitao/goauthes/pkg/token"
)

// MemoryDB is memory store
var MemoryDB = map[string][]Memory{}

// Memory is a storage specific implemention
type Memory struct {
	Data Base
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
	log.Printf("%v", MemoryDB)
	return true, nil
}

// Refresh is to delete and create data in db
func (s *Memory) Refresh() (Storage, error) {
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
	// rebuild storage
	sto, err := s.makeNewByOld()
	if err != nil {
		return &sto, err
	}
	return &sto, nil
}

// Delete is to delete data in db
func (s *Memory) Delete() (bool, error) {
	delete(MemoryDB, "goauthes:access:"+s.Data.Access)
	delete(MemoryDB, "goauthes:refresh:"+s.Data.Refresh)
	return true, nil
}

// ToToken is pick the token data
func (s *Memory) ToToken() (token.Token, error) {
	tok := token.Token{
		AccessToken:  s.Data.Access,
		TokenType:    s.Data.TokenType,
		ExpiresIn:    s.Data.AccessExpiresIn,
		RefreshToken: s.Data.Refresh,
	}
	return tok, nil
}

// ToAuth is pick the auth data
func (s *Memory) makeNewByOld() (m Memory, err error) {
	expireTime, err := strconv.Atoi(os.Getenv("GOAUTHES_TOKEN_EXPIRE"))
	tok, err := token.GenerateToken("Bearer", int64(expireTime))
	m.Data = Base{
		ClientID:         s.Data.ClientID,
		UserID:           s.Data.UserID,
		RedirectURI:      s.Data.RedirectURI,
		Scope:            s.Data.Scope,
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
	if err != nil {
		return m, err
	}
	return m, nil
}
