package crypto

// HashFunc represents a function that creates a password-hash from body and salt.
type HashFunc func(body string, salt string) (digest []byte, err error)
