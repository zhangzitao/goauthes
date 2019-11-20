package credential

import (
	"github.com/zhangzitao/goauthes/pkg/authorize"
)

// InputPassWord is input data
type InputPassWord struct {
	GrantType string
	Username  string
	Password  string
	Scope     string
}

// GenerateAuthorize , as it's name
func (s *InputPassWord) GenerateAuthorize() (auth authorize.Authorize, err error) {
	auth, err = authorize.GenerateAuthorize("PassWord", s.Username, s.Password, s.Scope)
	return auth, err
}
