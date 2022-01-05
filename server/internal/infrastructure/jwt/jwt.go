package jwt

import (
	"errors"
	"fmt"
	"time"

	jwtLib "github.com/dgrijalva/jwt-go"
)

// Util contains methods for creating and validating JSON Web Tokens (jwt)
type Util struct {
	jwtSecret []byte        `validate:"required,min=10"`
	subject   string        `validate:"required"`
	validity  time.Duration `validate:"required"`
}

// New creates a Util instance and returns it if argument validation succeeds
func New(jwtSecret string, subject string, validity time.Duration) (Util, error) {
	if jwtSecret == "" || subject == "" || validity == 0 {
		return Util{}, errors.New("arguments cannot be default values")
	}
	jwtUtil := Util{
		jwtSecret: []byte(jwtSecret),
		subject:   subject,
		validity:  validity,
	}
	return jwtUtil, nil
}

// CreateJWT creates a JWT for a user string
func (jwt Util) CreateJWT(userUniqueIdentifyer string) (token string, err error) {
	claims := &jwtLib.StandardClaims{
		Audience:  userUniqueIdentifyer,
		Issuer:    userUniqueIdentifyer,
		Subject:   jwt.subject,
		ExpiresAt: time.Now().Add(jwt.validity).Unix(),
	}
	tok := jwtLib.NewWithClaims(jwtLib.SigningMethodHS256, claims)
	return tok.SignedString(jwt.jwtSecret)

}

// ValidateJWT validates the JWT and returns the user string
func (jwt Util) ValidateJWT(token string) (user string, err error) {
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

		// Fetch the user from the JWT
		userClaim := claims["aud"]
		user, ok := userClaim.(string)
		if !ok {
			return "", fmt.Errorf("cannot convert aud claim of '%s' to user string", userClaim)
		}
		return user, nil

		// TODO: Add JWT id check
	}
	return "", errors.New("token validation failed")
}
