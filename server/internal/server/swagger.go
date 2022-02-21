package server

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"

	"git.bytecode.nl/bytecode/genesis/server/internal/interactors"
	"git.bytecode.nl/bytecode/genesis/server/swagger"
)

// @title Genesis
// @version 1
// @description Genesis API server

// @contact.name Customer Support
// @contact.url https://genesis.bytecode.nl
// @contact.email support@genesis.bytecode.nl

// @copyright Genesis
// @license.name GPL-3.0-only
// @license.url https://www.gnu.org/licenses/gpl-3.0.txt

// @securityDefinitions.apiKey JWT_User
// @in header
// @name Authorization

// @host api.genesis.bytecode.nl
// @BasePath /v1

func swaggerDoc(s *interactors.Services) func(c *gin.Context) {
	httpsReplacer := strings.NewReplacer(
		"https://", "",
		"http://", "")

	serverHostname := httpsReplacer.Replace(s.Config.ServerHostname)

	swagger.SwaggerInfo_swagger.Host = serverHostname

	return func(c *gin.Context) {
		c.String(http.StatusOK, swagger.SwaggerInfo_swagger.ReadDoc())
	}
}

func swaggerRedirect(c *gin.Context) {
	c.Redirect(http.StatusPermanentRedirect, "/v1/swagger/index.html")
}

func registerSwagger(r *gin.RouterGroup, s *interactors.Services) {
	r.GET("/swagger.json", swaggerDoc(s))
	r.GET("/swagger", swaggerRedirect)
	r.GET("/swagger/*any", ginSwagger.WrapHandler(
		swaggerfiles.Handler,
		ginSwagger.URL("/v1/swagger.json"),
	))
}
