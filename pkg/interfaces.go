package pkg

type (
	// Authorize is used building & verifying the data that to be authorized
	Authorize interface {
		GetType() string
		ToBasePassWord() (BasePassWord, bool)
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
		ToBase() BaseStorage
		Create() (bool, error)
		Refresh() (bool, error)
		Delete() (bool, error)
	}
)
