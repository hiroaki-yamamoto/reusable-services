package crypto

// PasswordHasher indicates the password hashing algorithm.
type PasswordHasher interface {
	Encode(body string, salt string) (digest []byte, err error)
	GetAlgoName() string
}
