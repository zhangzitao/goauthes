package interfaces

import base "github.com/zhangzitao/goauthes/oauth2/structs_base"

type (
	// Authorize is used building & verifying the data that to be authorized
	Authorize interface {
		GetType() string
		ToBasePassWord() (base.PassWordData, bool)
		Verify(string) (Authorize, error)
		VerifyRemote() (Authorize, error)
		VerifyLocal() (Authorize, error)
		Verified() bool
	}

	// Credential is input data maker
	Credential interface {
		ToArray() [4]string
	}

	// Storage is data handler
	Storage interface {
		ToBase() base.StorageData
		Create() (bool, error)
		Refresh() (bool, error)
		Delete() (bool, error)
	}
)
