package factory

import (
	"errors"

	base "github.com/zhangzitao/goauthes/oauth2/structs_base"
	"github.com/zhangzitao/goauthes/oauth2/interfaces"
	"github.com/zhangzitao/goauthes/oauth2/implementation/authorize"
)

// GenerateAuthorize is a factory to generate Authorize
func GenerateAuthorize(class string, data ...string) (auth interfaces.Authorize, err error) {
	switch class {
	case "PassWord":
		if len(data) != 3 {
			err := errors.New("Generate authorize param lenth error")
			return auth, err
		}
		auth = &authorize.PassWord{Data: base.PassWordData{
			UserName: data[0],
			PassWord: data[1],
			Scope:    data[2],
			UserID:   "",
		}, Type: "PassWord"}
	case "AuthorizationCode":
		// TODO
	}

	return auth, nil
}

// GenerateAuthorizeFromCredential , as it's name
func GenerateAuthorizeFromCredential(s interfaces.Credential) (auth interfaces.Authorize, err error) {
	arr := s.ToArray()
	switch arr[0] {
	case "PassWord":
		auth, err = GenerateAuthorize("PassWord", arr[1], arr[2], arr[3])
	}

	return auth, err
}
