package authorize

import (
	"errors"
)

// GenerateAuthorize is a factory to generate Authorize
func GenerateAuthorize(class string, data ...string) (auth Authorize, err error) {
	switch class {
	case "PassWord":
		if len(data) != 3 {
			err = errors.New("Generate authorize param lenth error")
			return auth, err
		}
		auth = &PassWord{data[0], data[1], data[2], ""}
	case "AuthorizationCode":
		// TODO
	}

	return auth, nil
}
