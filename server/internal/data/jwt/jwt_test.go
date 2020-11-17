package jwt_test

import (
	"testing"
	"time"

	"git.bytecode.nl/bytecode/genesis/internal/data/jwt"

	"github.com/stretchr/testify/assert"
)

const Year = time.Hour * 24 * 365

func TestNew_Valid(t *testing.T) {
	_, err := jwt.New("the secret", "user auth key", Year)
	assert.NoError(t, err)
}

func TestNew_Invalid(t *testing.T) {
	jwter, err := jwt.New("", "sub", Year)
	assert.Nil(t, jwter)
	assert.Error(t, err)
	jwter, err = jwt.New("key", "", Year)
	assert.Nil(t, jwter)
	assert.Error(t, err)
	jwter, err = jwt.New("key", "sub", 0)
	assert.Nil(t, jwter)
	assert.Error(t, err)
}

func TestJwtUtil_CreateJWT(t *testing.T) {
	j, err := jwt.New("key", "user auth key", Year)
	assert.NoError(t, err)

	token, err := j.CreateJWT("user")
	assert.NoError(t, err)
	assert.NotEmpty(t, token)
}
func TestJwtUtil_ValidateJWT_User(t *testing.T) {
	testUser := "john_doe@gmail.com"
	j, err := jwt.New("key", "user auth key", Year)
	assert.NoError(t, err)

	// Create key
	key, err := j.CreateJWT(testUser)
	assert.NoError(t, err)
	assert.NotEmpty(t, key)

	// Get the user from the generated key and check if it's the testUser
	user, err := j.ValidateJWT(key)
	assert.NoError(t, err)
	assert.Equal(t, testUser, user)
}

func TestJwtUtil_ValidateJWT_ExpiredKey(t *testing.T) {
	testUser := "john_doe@gmail.com"
	j, err := jwt.New("key", "user auth key", -Year) // Note the minus before the year, so it expired one year ago
	assert.NoError(t, err)

	// Create key
	key, err := j.CreateJWT(testUser)
	assert.NoError(t, err)
	assert.NotEmpty(t, key)

	// Validate JWT key that should err because the key is expired
	user, err := j.ValidateJWT(key)
	assert.Error(t, err)
	assert.NotEqual(t, testUser, user)
}

func TestJwtUtil_ValidateJWT_Subject(t *testing.T) {
	testUser := "john_doe@gmail.com"
	jOne, err := jwt.New("key", "test_one", Year)
	assert.NoError(t, err)
	jTwo, err := jwt.New("key", "test_two", Year)
	assert.NoError(t, err)

	// Create key
	key, err := jOne.CreateJWT(testUser)
	assert.NoError(t, err)
	assert.NotEmpty(t, key)

	// Validate JWT key that should err because the key subject is expired
	user, err := jTwo.ValidateJWT(key)
	assert.Error(t, err)
	assert.NotEqual(t, testUser, user)
}

func TestJwtUtil_ValidateJWT_DifferentSecrets(t *testing.T) {
	testUser := "john_doe@gmail.com"
	jOne, err := jwt.New("key1", "subject", Year)
	assert.NoError(t, err)
	jTwo, err := jwt.New("key2", "subject", Year)
	assert.NoError(t, err)

	// Create key
	key, err := jOne.CreateJWT(testUser)
	assert.NoError(t, err)
	assert.NotEmpty(t, key)

	// Validate JWT key that should err because the jwtSecret is not correct
	user, err := jTwo.ValidateJWT(key)
	assert.Error(t, err)
	assert.NotEqual(t, testUser, user)
}
