package pkg

import "time"

// BasePassWord is a authorize data
type BasePassWord struct {
	UserName string
	PassWord string
	Scope    string
	UserID   string
}

// BaseStorage is storage data handler
type BaseStorage struct {
	ClientID         string
	UserID           string
	RedirectURI      string
	Scope            string
	Code             string
	CodeCreateAt     time.Time
	CodeExpiresIn    int64
	Access           string
	AccessCreateAt   time.Time
	AccessExpiresIn  int64
	Refresh          string
	RefreshCreateAt  time.Time
	RefreshExpiresIn int64
	TokenType        string
}
