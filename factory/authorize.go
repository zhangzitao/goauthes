package factory

import (
	"errors"

	"github.com/zhangzitao/goauthes/pkg"
	"github.com/zhangzitao/goauthes/pkg/authorize"
)

// GenerateAuthorize is a factory to generate Authorize
func GenerateAuthorize(class string, data ...string) (auth pkg.Authorize, err error) {
	switch class {
	case "PassWord":
		if len(data) != 3 {
			err := errors.New("Generate authorize param lenth error")
			return auth, err
		}
		auth = &authorize.PassWord{
			UserName: data[0],
			PassWord: data[1],
			Scope:    data[2],
			UserID:   "",
		}
	case "AuthorizationCode":
		// TODO
	}

	return auth, nil
}

// GenerateAuthorizeFromCredential , as it's name
func GenerateAuthorizeFromCredential(s pkg.Credential) (auth pkg.Authorize, err error) {
	arr := s.ToArray()
	switch arr[0] {
	case "PassWord":
		auth, err = GenerateAuthorize("PassWord", arr[1], arr[2], arr[3])
	}

	return auth, err
}
