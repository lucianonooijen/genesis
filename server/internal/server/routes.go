package server

import (
	"git.bytecode.nl/bytecode/genesis/internal/server/handlers"
	"github.com/gin-gonic/gin"
)

func registerRoutes(r *gin.RouterGroup, h handlers.Handlers) {
	// STATUS
	r.GET("/status", h.Status)
}
