package server

import (
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"

	sentrygin "github.com/getsentry/sentry-go/gin"
	"go.uber.org/zap"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/sqreen/go-agent/sdk/middleware/sqgin"

	"git.bytecode.nl/bytecode/genesis/server/internal/constants"
	"git.bytecode.nl/bytecode/genesis/server/internal/domains/user"
	"git.bytecode.nl/bytecode/genesis/server/internal/interactors"
	"git.bytecode.nl/bytecode/genesis/server/internal/server/handlers"
	"git.bytecode.nl/bytecode/genesis/server/internal/server/middleware"
)

// GinServer is the Server instance struct.
type GinServer struct {
	port   int
	Router *gin.Engine
}

// Start the server instance.
func (s GinServer) Start() error {
	return s.Router.Run(":" + strconv.Itoa(s.port))
}

// New creates a new Server instance with middleware and handlers added.
// Use Server.Start() to run the server.
func New(services *interactors.Services) (GinServer, error) {
	debug := services.Config.IsDevMode
	port := services.Config.ServerPort
	logger := services.BaseLogger.Named("gin_init")

	if debug {
		logger.Debug("Detected debug mode for Gin")
		gin.SetMode(gin.DebugMode)
	} else {
		logger.Debug("Detected production mode for Gin")
		gin.SetMode(gin.ReleaseMode)
	}

	gin.DefaultWriter = io.Discard // Not so clean way to don't get the "you are using Gin debug" error
	server := GinServer{
		Router: gin.New(),
		port:   port,
	}
	gin.DefaultWriter = os.Stdout // Reset to the default writer

	registerMiddleware(services, server.Router, debug)

	initializedHandlers, err := handlers.New(services)
	if err != nil {
		return server, err
	}

	setGinRouteLogger(logger) // Print the Gin routes using our own logger
	logger.Debug("Registering routes")
	registerRoutes(server.Router.Group(constants.BasePathAPI), initializedHandlers)
	logger.Debug("Routes registered")

	return server, nil
}

func registerMiddleware(services *interactors.Services, router *gin.Engine, devMode bool) {
	if !devMode { // Run Sqreen in production
		router.Use(sqgin.Middleware())
	}

	router.Use(gin.Recovery())
	router.Use(middleware.GinLogger(services.BaseLogger))

	config := cors.DefaultConfig()
	config.AllowAllOrigins = true
	config.AllowHeaders = []string{ // TODO: Remove unused
		"Access-Control-Allow-Headers",
		"Content-Type",
		"Content-Length",
		"Accept-Encoding",
		"X-CSRF-Token",
		"accept",
		"origin",
		"Cache-Control",
		"Authorization",
		"X-Genesis-App-Version",
	}

	router.Use(sentrygin.New(sentrygin.Options{}))
	router.Use(cors.New(config))
	router.Use(middleware.EnsureKeysMap())
	router.Use(middleware.JwtAuth(services.BaseLogger, user.GenerateUserJwtMiddleware(services)))
	router.Use(middleware.VersionCheck(services.BaseLogger))

	if err := router.SetTrustedProxies([]string{}); err != nil { // TODO: Set this
		services.BaseLogger.Named("server/registerMiddleware").Fatal("could not set trusted proxies in Gin", zap.Error(err))
	}
}

func setGinRouteLogger(logger *zap.Logger) {
	gin.DebugPrintRouteFunc = func(httpMethod, absolutePath, handlerName string, nuHandlers int) {
		handlerShort := strings.ReplaceAll(handlerName, "git.bytecode.nl/bytecode/genesis/server/internal/server/handlers.", "") // TODO: cleaner
		handlerShort = strings.ReplaceAll(handlerShort, "github.com/gin-gonic/", "")                                             // TODO: cleaner
		logger.Debug(fmt.Sprintf("Route registered: %-6s %-25s --> %s (%d handlers)", httpMethod, absolutePath, handlerShort, nuHandlers))
	}
}
