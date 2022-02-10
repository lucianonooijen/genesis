package middleware

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/hashicorp/go-version"
	"go.uber.org/zap"

	"git.bytecode.nl/bytecode/genesis/server/internal/constants"
	"git.bytecode.nl/bytecode/genesis/server/internal/server/responses"
)

// VersionCheck is the middleware used to check if the "X-Genesis-App-Version" header exists, and if so, validate whether the client version is acceptable.
func VersionCheck(loggerParent *zap.Logger) gin.HandlerFunc {
	l := loggerParent.Named("server/middleware/VersionCheck")
	res := responses.New()

	return func(c *gin.Context) {
		versionHeader := c.GetHeader("X-Genesis-App-Version")

		log := l.With(zap.String("app_version", versionHeader))

		if versionHeader == "" {
			log.Debug("no version header found")

			c.Next()

			return
		}

		log.Debug("version header is set")

		shouldUpdate, err := checkAppVersionForForcedUpdate(versionHeader)
		if err != nil {
			log.Error("unexpected error from checkAppVersionForForcedUpdate", zap.Error(err))

			res.ServerError(c, err)

			c.Abort()

			return
		}

		if !shouldUpdate {
			log.Info("app version is ok, continuing")

			c.Next()

			return
		}

		log.Info("rejecting request, client requires update", zap.String("min_version", constants.MinimumClientVersion))

		res.ClientError(c,
			responses.StatusCodes.MustUpgrade,
			"Client update required",
			fmt.Sprintf("Client version %s is too old. Minimum version is %s", versionHeader, constants.MinimumClientVersion), nil)

		c.Abort()
	}
}

func checkAppVersionForForcedUpdate(appVersionString string) (bool, error) {
	appVersion, err := version.NewVersion(appVersionString)
	if err != nil {
		return false, err
	}

	minVersion, err := version.NewVersion(constants.MinimumClientVersion)
	if err != nil {
		return false, err
	}

	shouldUpdate := appVersion.LessThan(minVersion)

	return shouldUpdate, nil
}
