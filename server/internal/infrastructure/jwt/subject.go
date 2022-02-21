// nolint:dupl // duplication in this case is cleaner than having a shared extractor (simpler)
package jwt

import (
	"fmt"

	jwtLib "github.com/dgrijalva/jwt-go"
)

// ExtractSubject parses the JWT token and ONLY returns the subject.
// Use this with caution, because the TOKEN SIGNATURE IS NOT VALIDATED!
// This is useful in cases where you have multiple JWT verifying instances
// and need to select which one base on the subject.
// Never use this function for validating JWTs!
func ExtractSubject(token string) (subject string, err error) {
	tok, err := jwtLib.Parse(token, nil)
	if err != nil {
		// This error is expected.
		if err.Error() != noKeyFuncJwtLibErr {
			return "", err
		}
	}

	if tok == nil {
		return "", fmt.Errorf("could not parse token claims")
	}

	claims, ok := tok.Claims.(jwtLib.MapClaims)
	if !ok {
		return "", fmt.Errorf("could not extract token claims")
	}

	subRaw := claims["sub"]
	if subRaw == nil {
		return "", fmt.Errorf("subject in map is nil")
	}

	sub, ok := subRaw.(string)
	if !ok {
		return "", fmt.Errorf("subject could not be converted to string")
	}

	if sub == "" {
		return "", fmt.Errorf("subject is not set")
	}

	return sub, nil
}
