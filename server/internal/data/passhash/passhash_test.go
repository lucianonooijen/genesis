package passhash_test

import (
	"testing"

	"git.bytecode.nl/bytecode/genesis/internal/data/passhash"

	"github.com/stretchr/testify/assert"
)

func TestNew(t *testing.T) {
	passhash.New()
}

func TestPasswordHasher_PlainTextToHash(t *testing.T) {
	hasher := passhash.New()
	hash, err := hasher.PlainTextToHash("password")
	assert.NoError(t, err)
	assert.NotEmpty(t, hash)
}

func TestPasswordHasher_ComparePassToHash_Correct(t *testing.T) {
	plaintextPassword := "geitje123"
	hasher := passhash.New()
	hash, err := hasher.PlainTextToHash(plaintextPassword)
	assert.NoError(t, err)
	assert.NotEmpty(t, hash)

	err = hasher.ComparePassToHash(plaintextPassword, hash)
	assert.NoError(t, err)
}

func TestPasswordHasher_ComparePassToHash_Incorrect(t *testing.T) {
	hasher := passhash.New()
	hash, err := hasher.PlainTextToHash("correctpass")
	assert.NoError(t, err)
	assert.NotEmpty(t, hash)

	err = hasher.ComparePassToHash("incorrectpass", hash)
	assert.Error(t, err)
}
