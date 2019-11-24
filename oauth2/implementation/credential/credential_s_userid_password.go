package credential

// InputPassWord is input data
type InputPassWord struct {
	GrantType string
	Username  string
	Password  string
	Scope     string
}

//ToArray is
func (s *InputPassWord) ToArray() [4]string {
	arr := [4]string{
		s.GrantType,
		s.Username,
		s.Password,
		s.Scope,
	}
	return arr
}
