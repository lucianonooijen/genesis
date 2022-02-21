package handlers

import (
	"fmt"

	"github.com/gin-gonic/gin"

	"git.bytecode.nl/bytecode/genesis/server/internal/constants"
	"git.bytecode.nl/bytecode/genesis/server/internal/data/database"
	"git.bytecode.nl/bytecode/genesis/server/internal/server/responses"
)

var (
	errDatabaseUserNotOk = fmt.Errorf("could not convert user data to user")
	errDatabaseUserNil   = fmt.Errorf("user from keys is nil")
)

func extractUserFromRequest(c *gin.Context) (user *database.User, success bool) {
	userFromKeys := c.Keys[constants.GinContextKeyUser]
	usr, ok := userFromKeys.(*database.User)

	if !ok {
		responses.New().InternalServerError(c, errDatabaseUserNotOk)
		c.Abort()
	}

	if usr == nil {
		responses.New().InternalServerError(c, errDatabaseUserNil)
		c.Abort()
	}

	return usr, ok
}
