package server

import (
	"github.com/gin-gonic/gin"

	"git.bytecode.nl/bytecode/genesis/internal/constants"
	"git.bytecode.nl/bytecode/genesis/internal/server/handlers"
)

// Registers routes to the Gin RouterGroup, in an ExpressJS-like fashion
func registerRoutes(r *gin.RouterGroup, h handlers.Handlers) {
	// TODO: Use Embed.FS for this
	// https://github.com/gin-contrib/static/issues/19
	// STATIC FILES
	r.Static(constants.APIStaticPath, "./static")

	// STATUS
	r.GET("/status", h.Status)

	// USER LOGIN AND REGISTER
	r.POST("/user/register", h.CreateUser)
}
