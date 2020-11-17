package passhash

import (
	"git.bytecode.nl/bytecode/genesis/internal/interactors"
	"golang.org/x/crypto/bcrypt"
)

// PasswordHasher contains methods for generating and validating password hashes
type PasswordHasher struct {
	cost int
}

// New returns a PasswordHasher instance
func New() interactors.PasswordHasher {
	return PasswordHasher{
		cost: bcrypt.DefaultCost,
	}
}

// PlainTextToHash creates a password hash using bcrypt
func (ph PasswordHasher) PlainTextToHash(plaintextPassword string) (hash string, err error) {
	hashByteArr, err := bcrypt.GenerateFromPassword([]byte(plaintextPassword), ph.cost)
	if err != nil {
		return "", err
	}
	return string(hashByteArr), err
}

// ComparePassToHash checks if a plain text password is valid for a password hash (check if password is correct)
func (ph PasswordHasher) ComparePassToHash(plaintextPassword string, hash string) (err error) {
	return bcrypt.CompareHashAndPassword([]byte(hash), []byte(plaintextPassword))
}
