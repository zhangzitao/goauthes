package token

// Token is a
type Token struct {
	AccessToken  string
	TokenType    string
	ExpiresIn    int64
	RefreshToken string
}
