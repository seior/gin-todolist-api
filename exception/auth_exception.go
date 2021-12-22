package exception

type AuthException struct {
	Error string
}

func NewAuthException(error string) AuthException {
	return AuthException{Error: error}
}
