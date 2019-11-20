package authorize

import (
	"encoding/json"
	"errors"
	"net/http"
	"net/url"
	"os"
)

// PassWord is used for password mode of oauth2, an implemention of Authorize
type PassWord struct {
	UserName string
	PassWord string
	Scope    string
	UserID   string
}

// Verify is switcher
func (s *PassWord) Verify(class string) (Authorize, error) {
	switch class {
	case "Remote":
		return s.VerifyRemote()
	case "Local":
		return s.VerifyLocal()
	}
	return s, nil
}

// VerifyRemote the password, and wait for userid return
func (s *PassWord) VerifyRemote() (auth Authorize, err error) {
	values := url.Values{"username": {s.UserName}, "password": {s.PassWord}, "scope": {s.Scope}}
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
	auth = &PassWord{
		UserName: s.UserName,
		PassWord: s.PassWord,
		Scope:    s.Scope,
		UserID:   data.Message,
	}
	return auth, nil
}

// VerifyLocal the password, local, modify this if you want to use single server
func (s *PassWord) VerifyLocal() (Authorize, error) {
	return s, nil
}

// Verified is Check if auth verified yet
func (s *PassWord) Verified() bool {
	if s.UserID == "" {
		return false
	}
	return true
}
