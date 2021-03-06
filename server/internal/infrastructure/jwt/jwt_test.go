package jwt_test

import (
	"testing"
	"time"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"

	"git.bytecode.nl/bytecode/genesis/server/internal/infrastructure/jwt"
)

const Year = time.Hour * 24 * 365

var testUser = "john_doe@protonmail.com"
var testUUID = uuid.New()

func TestNew_Valid(t *testing.T) {
	_, err := jwt.New("the secret", "user auth key", Year)
	assert.NoError(t, err)
}

func TestNew_Invalid(t *testing.T) {
	jwter, err := jwt.New("", "sub", Year)
	assert.Empty(t, jwter)
	assert.Error(t, err)
	jwter, err = jwt.New("key", "", Year)
	assert.Empty(t, jwter)
	assert.Error(t, err)
	jwter, err = jwt.New("key", "sub", 0)
	assert.Empty(t, jwter)
	assert.Error(t, err)
}

func TestJwtUtil_CreateJWT(t *testing.T) {
	j, err := jwt.New("key", "user auth key", Year)
	assert.NoError(t, err)

	token, err := j.CreateJWT("user", testUUID)
	assert.NoError(t, err)
	assert.NotEmpty(t, token)
}
func TestJwtUtil_ValidateJWT_User(t *testing.T) {
	j, err := jwt.New("key", "user auth key", Year)
	assert.NoError(t, err)

	// Create key
	key, err := j.CreateJWT(testUser, testUUID)
	assert.NoError(t, err)
	assert.NotEmpty(t, key)

	// Get the user from the generated key and check if it's the testUser
	user, err := j.ValidateJWT(key, testUUID)
	assert.NoError(t, err)
	assert.Equal(t, testUser, user)
}

func TestJwtUtil_ValidateJWT_ExpiredKey(t *testing.T) {
	j, err := jwt.New("key", "user auth key", -Year) // Note the minus before the year, so it expired one year ago
	assert.NoError(t, err)

	// Create key
	key, err := j.CreateJWT(testUser, testUUID)
	assert.NoError(t, err)
	assert.NotEmpty(t, key)

	// Validate JWT key that should err because the key is expired
	user, err := j.ValidateJWT(key, testUUID)
	assert.Error(t, err)
	assert.NotEqual(t, testUser, user)
}

func TestJwtUtil_ValidateJWT_Subject(t *testing.T) {
	jOne, err := jwt.New("key", "test_one", Year)
	assert.NoError(t, err)
	jTwo, err := jwt.New("key", "test_two", Year)
	assert.NoError(t, err)

	// Create key
	key, err := jOne.CreateJWT(testUser, testUUID)
	assert.NoError(t, err)
	assert.NotEmpty(t, key)

	// Validate JWT key that should err because the key subject is expired
	user, err := jTwo.ValidateJWT(key, testUUID)
	assert.Error(t, err)
	assert.NotEqual(t, testUser, user)
}

func TestJwtUtil_ValidateJWT_DifferentSecrets(t *testing.T) {
	jOne, err := jwt.New("key1", "subject", Year)
	assert.NoError(t, err)
	jTwo, err := jwt.New("key2", "subject", Year)
	assert.NoError(t, err)

	// Create key
	key, err := jOne.CreateJWT(testUser, testUUID)
	assert.NoError(t, err)
	assert.NotEmpty(t, key)

	// Validate JWT key that should err because the jwtSecret is not correct
	user, err := jTwo.ValidateJWT(key, testUUID)
	assert.Error(t, err)
	assert.NotEqual(t, testUser, user)
}

func TestJwtUtil_ValidateJWT_ChangedKeyID(t *testing.T) {
	testUUIDOne := uuid.New()
	testUUIDTwo := uuid.New()
	j, err := jwt.New("key", "user auth key", Year)
	assert.NoError(t, err)

	// Create key
	key, err := j.CreateJWT(testUser, testUUIDOne)
	assert.NoError(t, err)
	assert.NotEmpty(t, key)

	// Get the user from the generated key and check if it's the testUser
	user, err := j.ValidateJWT(key, testUUIDTwo)
	assert.Error(t, err)
	assert.Equal(t, jwt.ErrDifferentKeyID, err)
	assert.Empty(t, user)
}
