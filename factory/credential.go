package factory

import (
	"errors"
	"log"

	"github.com/zhangzitao/goauthes/pkg"
	"github.com/zhangzitao/goauthes/pkg/credential"
)

// GenerateCredential is factory function
func GenerateCredential(class string, data ...string) (cred pkg.Credential, err error) {
	switch class {
	case "PassWord":
		if len(data) != 4 {
			err = errors.New("Generate Credential param lenth error")
			log.Println("error:", err)
			return cred, err
		}
		cred = &credential.InputPassWord{
			GrantType: data[0],
			Username:  data[1],
			Password:  data[2],
			Scope:     data[3],
		}
	case "AuthorizationCode":
	}
	return cred, nil
}
