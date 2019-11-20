package credential

import (
	"errors"
	"log"
)

// GenerateCredential is factory function
func GenerateCredential(class string, data ...string) (cred Credential, err error) {
	switch class {
	case "PassWord":
		if len(data) != 4 {
			err = errors.New("Generate Credential param lenth error")
			log.Println("error:", err)
			return cred, err
		}
		cred = &InputPassWord{data[0], data[1], data[2], data[3]}
	case "AuthorizationCode":
	}
	return cred, nil
}
