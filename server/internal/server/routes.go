package server

import (
	"git.bytecode.nl/bytecode/genesis/internal/server/handlers"
	"github.com/gin-gonic/gin"
)

// Registers routes to the Gin RouterGroup, in an ExpressJS-like fashion
func registerRoutes(r *gin.RouterGroup, h handlers.Handlers) {
	// STATUS
	r.GET("/status", h.Status)
}
