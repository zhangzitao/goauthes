package storage

import "github.com/zhangzitao/goauthes/pkg/token"

// Storage is data handler
type Storage interface {
	Create() (bool, error)
	Refresh() (Storage, error)
	Delete() (bool, error)
	ToToken() (token.Token, error)
}
