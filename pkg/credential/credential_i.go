package credential

import (
	"github.com/zhangzitao/goauthes/pkg/authorize"
)

// Credential is input data maker
type Credential interface {
	GenerateAuthorize() (authorize.Authorize, error)
}
