package jwt

import (
	"errors"
	"fmt"
	"time"

	jwtLib "github.com/dgrijalva/jwt-go"
	"github.com/google/uuid"
)

// Issuer is the issuer used for JWTs.
const Issuer = "genesis_v1" // if this value is changed, all tokens are invalidated

// Util contains methods for creating and validating JSON Web Tokens (jwt).
type Util struct {
	jwtSecret []byte        `validate:"required,min=10"`
	subject   string        `validate:"required"`
	validity  time.Duration `validate:"required"`
}

// ErrDifferentKeyID is the error given when the password uuid (key id) of a jwt is not the expected value.
// This happens when a user has changed his/her password.
var ErrDifferentKeyID = fmt.Errorf("jwt id is not the current password uuid of the user")

// New creates a Util instance and returns it if argument validation succeeds.
func New(jwtSecret, subject string, validity time.Duration) (*Util, error) {
	if jwtSecret == "" || subject == "" || validity == 0 {
		return nil, errors.New("arguments cannot be default values")
	}

	jwtUtil := Util{
		jwtSecret: []byte(jwtSecret),
		subject:   subject,
		validity:  validity,
	}

	return &jwtUtil, nil
}

// CreateJWT creates a JWT for a user string.
func (jwt Util) CreateJWT(userUniqueIdentifyer string, keyUUID uuid.UUID) (token string, err error) {
	claims := &jwtLib.StandardClaims{
		Audience:  userUniqueIdentifyer,
		Issuer:    Issuer,
		Id:        keyUUID.String(),
		Subject:   jwt.subject,
		ExpiresAt: time.Now().Add(jwt.validity).Unix(),
	}
	tok := jwtLib.NewWithClaims(jwtLib.SigningMethodHS256, claims)

	return tok.SignedString(jwt.jwtSecret)
}

// ValidateJWT validates the JWT and returns the user string.
func (jwt Util) ValidateJWT(token string, currentUserPasswordUUID uuid.UUID) (user string, err error) {
	tok, err := jwtLib.Parse(token, func(token *jwtLib.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwtLib.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return jwt.jwtSecret, nil
	})

	if err != nil {
		return "", err
	}

	if claims, ok := tok.Claims.(jwtLib.MapClaims); ok && tok.Valid {
		// Check the jwt subject
		subjectClaim := claims["sub"]
		subject, ok := subjectClaim.(string)

		if !ok {
			return "", fmt.Errorf("cannot convert sub claim of '%s' to string", subjectClaim)
		}

		if subject != jwt.subject {
			return "", fmt.Errorf("JWT subject (%s) does not match Util instance subject (%s)", subject, jwt.subject)
		}

		// Check the key id in case this JWT is invalidated
		keyIDClaim := claims["jti"]

		keyID, ok := keyIDClaim.(string)
		if !ok {
			return "", fmt.Errorf("cannot convert jti claim of '%s' to keyID string", keyID)
		}

		if keyID != currentUserPasswordUUID.String() {
			return "", ErrDifferentKeyID
		}

		// Fetch the user from the JWT
		userClaim := claims["aud"]

		user, ok := userClaim.(string)
		if !ok {
			return "", fmt.Errorf("cannot convert aud claim of '%s' to user string", userClaim)
		}

		return user, nil
	}

	return "", errors.New("token validation failed")
}
