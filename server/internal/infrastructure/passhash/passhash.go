package passhash

import (
	"golang.org/x/crypto/bcrypt"
)

// Util contains methods for generating and validating password hashes.
type Util struct {
	cost int
}

// New returns a PasswordHasher instance.
func New() *Util {
	return &Util{
		cost: bcrypt.DefaultCost,
	}
}

// PlainTextToHash creates a password hash using bcrypt.
func (ph Util) PlainTextToHash(plaintextPassword string) (hash string, err error) {
	hashByteArr, err := bcrypt.GenerateFromPassword([]byte(plaintextPassword), ph.cost)
	if err != nil {
		return "", err
	}

	return string(hashByteArr), err
}

// ComparePassToHash checks if a plain text password is valid for a password hash (check if password is correct).
func (ph Util) ComparePassToHash(plaintextPassword, hash string) (err error) {
	return bcrypt.CompareHashAndPassword([]byte(hash), []byte(plaintextPassword))
}
