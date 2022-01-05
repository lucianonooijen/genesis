package server

import (
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
	"strings"

	"go.uber.org/zap"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/sqreen/go-agent/sdk/middleware/sqgin"

	"git.bytecode.nl/bytecode/genesis/internal/constants"
	"git.bytecode.nl/bytecode/genesis/internal/interactors"
	"git.bytecode.nl/bytecode/genesis/internal/server/handlers"
	"git.bytecode.nl/bytecode/genesis/internal/server/middleware"
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

	gin.DefaultWriter = ioutil.Discard // Not so clean way to don't get the "you are using Gin debug" error // FIXME: Still get it
	server := GinServer{
		Router: gin.New(),
		port:   port,
	}
	gin.DefaultWriter = os.Stdout // Reset to the default writer

	registerMiddleware(services.BaseLogger, server.Router, debug)

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

func registerMiddleware(logger *zap.Logger, router *gin.Engine, devMode bool) {
	if !devMode { // Run Sqreen in production
		router.Use(sqgin.Middleware())
	}

	router.Use(gin.Recovery())
	router.Use(middleware.GinLogger(logger))

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
		"Authorization"}

	router.Use(cors.New(config))

	if err := router.SetTrustedProxies([]string{}); err != nil { // TODO: Set this
		logger.Fatal("could not set trusted proxies in Gin", zap.Error(err))
	}
}

// TODO: Add to registerMiddleware end
// nolint:gocritic
// router.Use(middleware.EnsureKeysMap())
// router.Use(middleware.JwtAuth(r.User.CheckUserJwt))

func setGinRouteLogger(logger *zap.Logger) {
	gin.DebugPrintRouteFunc = func(httpMethod, absolutePath, handlerName string, nuHandlers int) {
		handlerShort := strings.ReplaceAll(handlerName, "git.bytecode.nl/bytecode/genesis/internal/server/handlers.", "") // TODO: cleaner
		handlerShort = strings.ReplaceAll(handlerShort, "github.com/gin-gonic/", "")                                      // TODO: cleaner
		logger.Debug(fmt.Sprintf("Route registered: %-6s %-25s --> %s (%d handlers)", httpMethod, absolutePath, handlerShort, nuHandlers))
	}
}
