package interactors

// PasswordHasher is the interface returned by data/passhash
type PasswordHasher interface {
	PlainTextToHash(plaintextPassword string) (hash string, err error)
	ComparePassToHash(plaintextPassword string, hash string) (err error)
}
