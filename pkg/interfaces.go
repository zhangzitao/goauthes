package pkg

type (
	// Authorize is used building & verifying the data that to be authorized
	Authorize interface {
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
		Create() (bool, error)
		Refresh() (bool, error)
		Delete() (bool, error)
		ToBase() BaseStorage
	}

	// Token is
	Token interface {
		Access()
		Refresh()
	}
)
