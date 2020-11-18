package server

import (
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
	"strings"

	"git.bytecode.nl/bytecode/genesis/internal/constants"

	"git.bytecode.nl/bytecode/genesis/internal/utils/logger"

	"git.bytecode.nl/bytecode/genesis/internal/server/handlers"
	"git.bytecode.nl/bytecode/genesis/internal/server/middleware"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/sqreen/go-agent/sdk/middleware/sqgin"
)

var log = logger.New("server_init")

// GinServer is the Server instance struct
type GinServer struct {
	port   int
	Router *gin.Engine
}

// Requirements are the requirements for creating a new Server instance
type Requirements struct {
	Debug bool
	Port  int `validate:"required"`
}

// Start the server instance
func (s GinServer) Start() error {
	return s.Router.Run(":" + strconv.Itoa(s.port))
}

// New creates a new Server instance with middleware and handlers added.
// Use Server.Start() to run the server.
func New(r Requirements) (GinServer, error) {
	validate := validator.New()
	err := validate.Struct(r)
	if err != nil {
		return GinServer{}, err
	}

	if r.Debug {
		log.Debug("Detected debug mode for Gin")
		gin.SetMode(gin.DebugMode)
	} else {
		log.Debug("Detected production mode for Gin")
		gin.SetMode(gin.ReleaseMode)
	}

	gin.DefaultWriter = ioutil.Discard // Hacky way to don't get the "you are using Gin debug" error
	server := GinServer{
		Router: gin.New(),
		port:   r.Port,
	}
	gin.DefaultWriter = os.Stdout // Reset to the default writer

	registerMiddleware(server.Router, r.Debug)

	initializedHandlers, err := handlers.New()
	if err != nil {
		return server, err
	}

	setGinRouteLogger() // Print the Gin routes using our own logger
	log.Debug("Registering routes")
	registerRoutes(server.Router.Group(constants.BasePathAPI), initializedHandlers)
	log.Debug("Routes registered")
	return server, nil
}

// Registers middleware
func registerMiddleware(router *gin.Engine, devMode bool) {
	if !devMode { // Run Sqreen in production
		router.Use(sqgin.Middleware())
	}
	router.Use(gin.Recovery())
	router.Use(middleware.GinLogger())
	//router.Use(middleware.ActivityTableLog(r.SaveActivity))
	config := cors.DefaultConfig()
	config.AllowAllOrigins = true
	config.AllowHeaders = []string{ // TODO: Remove unused
		"Access-Control-Allow-Headers",
		"Content-Type", "Content-Length",
		"Accept-Encoding", "X-CSRF-Token",
		"accept", "origin",
		"Cache-Control",
		"Authorization"}
	router.Use(cors.New(config))
	//router.Use(middleware.EnsureKeysMap())
	//router.Use(middleware.JwtAuth(r.User.CheckUserJwt))
}

func setGinRouteLogger() {
	gin.DebugPrintRouteFunc = func(httpMethod, absolutePath, handlerName string, nuHandlers int) {
		handlerShort := strings.ReplaceAll(handlerName, "git.bytecode.nl/bytecode/genesis/internal/server/handlers.", "") // TODO: cleaner
		handlerShort = strings.ReplaceAll(handlerShort, "github.com/gin-gonic/", "")                                      // TODO: cleaner
		log.Debug(fmt.Sprintf("Route registered: %-6s %-25s --> %s (%d handlers)", httpMethod, absolutePath, handlerShort, nuHandlers))
	}
}
