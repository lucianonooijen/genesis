package jwt_test

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"

	"git.bytecode.nl/bytecode/genesis/server/internal/infrastructure/jwt"
)

var (
	testSubject = "this_is_a_subject"
)

func TestExtractSubject_Extract(t *testing.T) {
	j, err := jwt.New("key1", testSubject, Year)
	assert.NoError(t, err)

	// Create key
	key, err := j.CreateJWT(testUser, testUUID)
	assert.NoError(t, err)
	assert.NotEmpty(t, key)

	// Extract JWT and check if it's the test subject
	sub, err := jwt.ExtractSubject(key)
	assert.NoError(t, err)
	assert.Equal(t, testSubject, sub)
}

func TestExtractSubject_NoValidateSignature(t *testing.T) {
	j, err := jwt.New("key1", testSubject, Year)
	assert.NoError(t, err)

	// Create key
	key, err := j.CreateJWT(testUser, testUUID)
	assert.NoError(t, err)
	assert.NotEmpty(t, key)

	// Modify the key to invalidate the signature
	keyParts := strings.Split(key, ".")
	exampleSignature := "VKPicz1jQzeysLyvjPxAJAJYzc0zHFVuMqabop9ovXc"
	keyParts[2] = exampleSignature
	keyWithInvalidSignature := strings.Join(keyParts, ".")

	// Extract JWT and check if it's the test subject
	sub, err := jwt.ExtractSubject(keyWithInvalidSignature)
	assert.NoError(t, err)
	assert.Equal(t, testSubject, sub)
}
