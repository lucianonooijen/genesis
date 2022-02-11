package middleware

import (
	"fmt"
	"strings"

	"github.com/getsentry/sentry-go"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"

	"git.bytecode.nl/bytecode/genesis/server/internal/constants"
	"git.bytecode.nl/bytecode/genesis/server/internal/domains/user"
	"git.bytecode.nl/bytecode/genesis/server/internal/infrastructure/jwt"
	"git.bytecode.nl/bytecode/genesis/server/internal/server/responses"
)

// JwtAuth is the Gin middleware function to check authentication and update the Gin context accordingly.
// nolint:funlen // it's clearer to have this as a long function than to break it up
func JwtAuth(loggerParent *zap.Logger, jwtHandlerUser user.HTTPJwtFunc) gin.HandlerFunc {
	logger := loggerParent.Named("server/middleware/JwtAuth")
	logStep := func(msg string) {
		logger.Debug(msg)
	}

	return func(c *gin.Context) {
		authHeader := c.Request.Header.Get("Authorization")
		if authHeader == "" {
			c.Keys[constants.GinContextKeyUser] = nil

			logStep("no auth header found")

			c.Next() // Start handling request

			return
		} // From here we know that the Authorization header has been set.

		parts := strings.Split(authHeader, " ")
		if len(parts) != 2 { // nolint:gomnd // context is clear enough
			jwtError(c, "the 'Authorization' header should be formatted like 'Bearer [jwt]'", logger)

			return
		} // Now we know that the Authorization header contains two words

		tokType := parts[0]
		jwtString := parts[1]

		if tokType != "Bearer" || jwtString == "" {
			jwtError(c, "jwt was not found or auth token type was not 'Bearer'", logger)

			return
		} // From here we know the Authorization header is formatted correctly

		// First, we recognize the subject of the JWT to use the correct handler.
		// IMPORTANT NOTE: this does NOT mean the JWT is valid, that is done later on.
		sub, err := jwt.ExtractSubject(jwtString)
		c.Keys[constants.GinContextKeyUser] = sub

		logStep(fmt.Sprintf("enriched request with role based on jwt subject: %s", sub))

		if err != nil {
			logStep("error: " + err.Error())

			jwtError(c, err.Error(), logger)

			return
		} // Now we have the subject from the JWT, but it's still not validated!

		switch sub {
		case constants.JwtSubjectUsers:
			logStep("subject is constants.JwtSubjectUsers")

			// If a deleted user tries to use their "old" JWT, the backend will
			// return a 404, because the email address will be modified, thus
			// a user with the old email will not be found.
			// An edge case to this, is if a user (1) deletes his/her account
			// and later on another user (2) creates an account with that same
			// email address. The old JWT for 1 now works for 2's account.
			u, err := jwtHandlerUser(jwtString)
			if err != nil {
				unauthError(c, err.Error(), logger)
				return
			} // From here we know that the user is correctly authenticated

			c.Keys[constants.GinContextKeyUser] = u
			c.Keys[constants.GinContextKeyRole] = constants.JwtSubjectUsers

			sentry.ConfigureScope(func(scope *sentry.Scope) {
				scope.SetUser(sentry.User{
					Email:     u.Email,
					ID:        fmt.Sprintf("%d", u.ID),
					IPAddress: c.ClientIP(),
					Username:  u.UserUuid.String(),
				})
			})

			logStep(fmt.Sprintf("enriched request with user id: %d (%s), email %s, name %s", u.ID, u.UserUuid, u.Email, u.FirstName))

			c.Next()

			return
		default:
			unauthError(c, "jwt subject was not recogized", logger)

			return
		}
	}
}

func jwtError(c *gin.Context, msg string, log *zap.Logger) {
	log.Named("jwtError").Warn(msg)

	res := responses.New()
	res.ClientError(c, responses.StatusCodes.BadRequest, "jwt error", msg, nil)

	c.Abort()
}

func unauthError(c *gin.Context, msg string, log *zap.Logger) {
	log.Named("unauthError").Warn(msg)

	res := responses.New()
	res.ClientError(c, responses.StatusCodes.UnauthorizedRequest, "unauthorized", msg, nil)

	c.Abort()
}
