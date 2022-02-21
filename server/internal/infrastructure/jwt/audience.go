// nolint:dupl // duplication in this case is cleaner than having a shared extractor (simpler)
package jwt

import (
	"fmt"

	jwtLib "github.com/dgrijalva/jwt-go"
)

var noKeyFuncJwtLibErr = "no Keyfunc was provided."

// ExtractAudience parses the JWT token and ONLY returns the audience.
// Use this with caution, because the TOKEN SIGNATURE IS NOT VALIDATED!
// This is useful in cases where you need the audience to fetch data,
// used to validate the JWT later on.
// Never use this function for validating JWTs!
func ExtractAudience(token string) (audience string, err error) {
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

	audRaw := claims["aud"]
	if audRaw == nil {
		return "", fmt.Errorf("audience in map is nil")
	}

	aud, ok := audRaw.(string)
	if !ok {
		return "", fmt.Errorf("audience could not be converted to string")
	}

	if aud == "" {
		return "", fmt.Errorf("audience is not set")
	}

	return aud, nil
}
