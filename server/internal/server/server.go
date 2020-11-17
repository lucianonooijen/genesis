package server

import (
	"strconv"

	"git.bytecode.nl/bytecode/genesis/internal/server/handlers"
	"git.bytecode.nl/bytecode/genesis/internal/server/middleware"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/sqreen/go-agent/sdk/middleware/sqgin"
)

const (
	// BasePathAPI is the base used for the base of all requests
	// Example: [server url]:[port]/[BasePathAPI]/[all other urls]
	BasePathAPI = "/v1"
)

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
		//r.Logger.Debug("Detected debug mode for Gin")
		gin.SetMode(gin.DebugMode)
	} else {
		//r.Logger.Debug("Detected production mode for Gin")
		gin.SetMode(gin.ReleaseMode)
	}

	server := GinServer{
		Router: gin.New(),
		port:   r.Port,
		//Logger: r.Logger,
	}

	registerMiddleware(server.Router, r.Debug)

	initializedHandlers, err := handlers.New()
	if err != nil {
		return server, err
	}

	registerRoutes(server.Router.Group(BasePathAPI), initializedHandlers)
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
	config.AllowHeaders = []string{
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
