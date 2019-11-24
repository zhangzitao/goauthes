package authorize

import (
	"encoding/json"
	"errors"
	"net/http"
	"net/url"
	"os"

	"github.com/zhangzitao/goauthes/oauth2/interfaces"
	base "github.com/zhangzitao/goauthes/oauth2/structs_base"
)

// PassWord is used for password mode of oauth2, an implemention of Authorize
type PassWord struct {
	Data base.PassWordData
	Type string
}

// Verify is switcher
func (s *PassWord) Verify(class string) (interfaces.Authorize, error) {
	switch class {
	case "Remote":
		return s.VerifyRemote()
	case "Local":
		return s.VerifyLocal()
	}
	return s, nil
}

// VerifyRemote the password, and wait for userid return
func (s *PassWord) VerifyRemote() (auth interfaces.Authorize, err error) {
	values := url.Values{"username": {s.Data.UserName}, "password": {s.Data.PassWord}, "scope": {s.Data.Scope}}
	// make request
	resp, err := http.PostForm(os.Getenv("GOAUTHES_AUTHORIZE_REMOTE_VERIFY_URL"), values)
	if err != nil {
		return auth, err
	}
	defer resp.Body.Close()
	if resp.StatusCode != 200 {
		return auth, errors.New("remote respond error:" + resp.Status)
	}
	// read respond
	type dataRespond struct {
		Success bool
		Message string
	}
	data := dataRespond{false, ""}
	err = json.NewDecoder(resp.Body).Decode(&data)
	if err != nil {
		return auth, err
	}
	if !data.Success || data.Message == "" {
		return auth, errors.New("read respond error:" + data.Message)
	}
	// get the id return from remote server
	auth = &PassWord{Data: base.PassWordData{
		UserName: s.Data.UserName,
		PassWord: s.Data.PassWord,
		Scope:    s.Data.Scope,
		UserID:   data.Message},
		Type: "PassWord",
	}
	return auth, nil
}

// VerifyLocal the password, local, modify this if you want to use single server
func (s *PassWord) VerifyLocal() (interfaces.Authorize, error) {
	return s, nil
}

// Verified is Check if auth verified yet
func (s *PassWord) Verified() bool {
	if s.Data.UserID == "" {
		return false
	}
	return true
}

// ToBasePassWord is
func (s *PassWord) ToBasePassWord() (base.PassWordData, bool) {
	return s.Data, true
}

// GetType is
func (s *PassWord) GetType() string {
	return s.Type
}
