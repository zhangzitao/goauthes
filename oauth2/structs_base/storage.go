package base

import "time"

// StorageData is storage data handler
type StorageData struct {
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
