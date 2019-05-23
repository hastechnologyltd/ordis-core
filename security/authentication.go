package security

type credentials struct {
	username []byte
	password []byte
}

func NewCredentials(username []byte, password []byte) credentials {
	return credentials{
		username: username,
		password: password,
	}
}

func (auth *credentials) Authenticate() bool {
	return false
}
