package error

type AuthError struct {
	Code    int
	Message string
}

func (e *AuthError) AuthError() string {
	return e.Message
}
