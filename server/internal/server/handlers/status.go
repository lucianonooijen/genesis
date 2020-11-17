package handlers

import "github.com/gin-gonic/gin"

// Status returns a 200 status code if everything is fine
func (h Handlers) Status(c *gin.Context) {
	// TODO: Check DB connection before sending 200
	h.sendSuccess(c, nil)
}
