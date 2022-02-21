package jwt_test

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"

	"git.bytecode.nl/bytecode/genesis/server/internal/infrastructure/jwt"
)

func TestExtractAudience_Extract(t *testing.T) {
	j, err := jwt.New("key1", testSubject, Year)
	assert.NoError(t, err)

	// Create key
	key, err := j.CreateJWT(testUser, testUUID)
	assert.NoError(t, err)
	assert.NotEmpty(t, key)

	// Extract JWT and check if it's the test audience
	aud, err := jwt.ExtractAudience(key)
	assert.NoError(t, err)
	assert.Equal(t, testUser, aud)
}

func TestExtractAudience_NoValidateSignature(t *testing.T) {
	j, err := jwt.New("key1", testSubject, Year)
	assert.NoError(t, err)

	// Create key
	key, err := j.CreateJWT(testUser, testUUID)
	assert.NoError(t, err)
	assert.NotEmpty(t, key)

	// Modify the key to invalidate the signature
	keyParts := strings.Split(key, ".")
	exampleSignature := "VKPicz1jQzeysLyvjPxAJAJYzc0zHFVuMqabop9ovxc"
	keyParts[2] = exampleSignature
	keyWithInvalidSignature := strings.Join(keyParts, ".")

	// Extract JWT and check if it's the test audience
	aud, err := jwt.ExtractAudience(keyWithInvalidSignature)
	assert.NoError(t, err)
	assert.Equal(t, testUser, aud)
}
