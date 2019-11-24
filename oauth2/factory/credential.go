package factory

import (
	"errors"
	"log"

	"github.com/zhangzitao/goauthes/oauth2/implementation/credential"
	"github.com/zhangzitao/goauthes/oauth2/interfaces"
)

// GenerateCredential is factory function
func GenerateCredential(class string, data ...string) (cred interfaces.Credential, err error) {
	switch class {
	case "PassWord":
		if len(data) != 4 {
			err = errors.New("Generate Credential param lenth error")
			log.Println("info:", err)
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
