package handlers

import (
	"database/sql"
	"errors"
	"fmt"

	jwtLib "github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/lib/pq"
	"golang.org/x/crypto/bcrypt"

	"git.bytecode.nl/bytecode/genesis/internal/server/responses"
)

func (h Handlers) handleDomainError(c *gin.Context, err error) {
	// Not found in database => 404
	if errors.Is(err, sql.ErrNoRows) {
		h.sendNotFound(c, err)
		return
	}

	// Postgres errors in separate handler
	if pqErr, ok := err.(*pq.Error); ok {
		handlePostgresError(c, pqErr)
		return
	}

	// Error: "signature is invalid", is unauthorized response
	if errors.Is(err, jwtLib.ErrSignatureInvalid) {
		r.ClientError(c, responses.StatusCodes.UnauthorizedRequest, "Invalid authentication token signature", "The JWT signature is invalid", err)
		return
	}

	// Bcrypt errors
	if errors.Is(err, bcrypt.ErrMismatchedHashAndPassword) {
		r.ClientError(c, responses.StatusCodes.UnauthorizedRequest, "Password and password hash mismatch", "incorrect password: hashed password is not the hash of the given password\"", err)
		return
	}

	// Default to server error (500)
	r.ServerError(c, err)
}

func handlePostgresError(c *gin.Context, pqErr *pq.Error) {
	formatErr := fmt.Errorf("%s (code %s, name %s)", pqErr.Detail, pqErr.Code, pqErr.Code.Name())

	switch pqErr.Code.Name() {
	case "unique_violation":
		detail := fmt.Sprintf("Data unique violation: %s", pqErr.Detail)
		r.ClientError(c, responses.StatusCodes.Conflict, "Duplicate entry in database", detail, formatErr)
	case "invalid_text_representation":
		err := fmt.Errorf("invalid input for enum type: %s (%s)", pqErr.Message, formatErr)
		r.ServerError(c, err)
	case "foreign_key_violation":
		detail := fmt.Sprintf("Foreign key database error: fkey = %s, message = %s", pqErr.Constraint, pqErr.Message)
		r.ClientError(c, responses.StatusCodes.BadRequest, "Invalid data", detail, formatErr)
	case "check_violation":
		detail := fmt.Sprintf("Input does not satisfy database data check: check_name = %s, message = %s", pqErr.Constraint, pqErr.Message)
		r.ClientError(c, responses.StatusCodes.BadRequest, "Invalid data", detail, formatErr)
	default:
		r.ServerError(c, formatErr)
	}
}
