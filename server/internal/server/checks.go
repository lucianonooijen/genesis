package server

import (
	"fmt"

	"github.com/gin-gonic/gin"

	"git.bytecode.nl/bytecode/genesis/server/internal/constants"
	"git.bytecode.nl/bytecode/genesis/server/internal/data/database"
	"git.bytecode.nl/bytecode/genesis/server/internal/server/responses"
)

func checkLoggedInAsUser(c *gin.Context) {
	res := responses.New()
	authHeader := c.Request.Header.Get("Authorization")

	// If the authHeader is not "", is must have been handled by the auth middleware and thus have the user in the context keys
	if authHeader == "" {
		res.ClientError(c, responses.StatusCodes.UnauthorizedRequest, "authenticated required", "this endpoint requires a user account", nil)

		c.Abort()

		return
	}

	hasRole := c.Keys[constants.GinContextKeyRole]
	requireRole := constants.JwtSubjectUsers

	// User must have user role
	if c.Keys[constants.GinContextKeyRole] != constants.JwtSubjectUsers {
		err := fmt.Errorf("endpoint requires role '%s' but has '%s'", hasRole, requireRole)
		res.ClientError(c, responses.StatusCodes.ForbiddenRequest, "incorrect role", "this endpoint is only available for people with a user role", err)

		c.Abort()

		return
	}

	userFromKeys := c.Keys[constants.GinContextKeyUser]

	_, ok := userFromKeys.(*database.User)
	if !ok {
		res.ServerError(c, fmt.Errorf("could not convert user data to user"))

		c.Abort()

		return
	}

	c.Next()
}
