package middleware

import (
	"net/http/httputil"
	"net/url"

	"github.com/gin-gonic/gin"
	"github.com/hashicorp/go-version"
	"go.uber.org/zap"

	"git.bytecode.nl/bytecode/genesis/server/internal/constants"
	"git.bytecode.nl/bytecode/genesis/server/internal/interactors"
	"git.bytecode.nl/bytecode/genesis/server/internal/server/responses"
)

// CanaryProxy is Gin middleware to reverse-proxy requests to a canary server
// if the client version is higher than the server version
// and if the `Config.ServerHostnameCanary` has a value set.
func CanaryProxy(s *interactors.Services) gin.HandlerFunc {
	l := s.BaseLogger.Named("server/middleware/CanaryProxy")
	res := responses.New()
	canaryHost := s.Config.ServerHostnameCanary

	if canaryHost == "" {
		l.Info("no value set for canaryHost, not using CanaryProxy middleware")
		return nil
	}

	l = l.With(zap.String("canaryHost", canaryHost))
	l.Info("using CanaryProxy middleware")

	return func(c *gin.Context) {
		versionHeader := c.GetHeader(constants.GinHeaderNameClientVersion)

		log := l.With(zap.String("app_version", versionHeader))

		if versionHeader == "" {
			log.Debug("no version header found")

			c.Next()

			return
		}

		log.Debug("version header is set")

		shouldProxy, err := checkAppVersionForCanaryProxy(versionHeader)
		if err != nil {
			log.Error("unexpected error from checkAppVersionForCanaryProxy", zap.Error(err))

			res.InternalServerError(c, err)

			c.Abort()

			return
		}

		if !shouldProxy {
			log.Info("client version is lower or equal than server, continuing")

			c.Next()

			return
		}

		log.Info("proxying request to canary server")

		canaryProxyHandler(c, s, l)

		c.Abort() // Request has been handled
	}
}

func canaryProxyHandler(c *gin.Context, s *interactors.Services, logMother *zap.Logger) {
	l := logMother.Named("canaryProxyHandler")

	canaryURL, err := url.Parse(s.Config.ServerHostnameCanary)
	if err != nil {
		l.Error("error parsing canary url", zap.Error(err))

		return
	}

	proxy := httputil.NewSingleHostReverseProxy(canaryURL)

	l.Info("start request proxy")

	proxy.ServeHTTP(c.Writer, c.Request)

	l.Info("complete request proxy")
}

func checkAppVersionForCanaryProxy(clientVersionString string) (bool, error) {
	clientVersion, err := version.NewVersion(clientVersionString)
	if err != nil {
		return false, err
	}

	serverVersion, err := version.NewVersion(constants.APIVersion)
	if err != nil {
		return false, err
	}

	shouldProxy := clientVersion.GreaterThan(serverVersion)

	return shouldProxy, nil
}
