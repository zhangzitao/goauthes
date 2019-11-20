package authorize

// Authorize is used building & verifying the data that to be authorized
type Authorize interface {
	Verify(string) (Authorize, error)
	VerifyRemote() (Authorize, error)
	VerifyLocal() (Authorize, error)
	Verified() bool
}
