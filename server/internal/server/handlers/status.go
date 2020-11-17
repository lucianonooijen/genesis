package handlers

import "github.com/gin-gonic/gin"

func (h Handlers) Status(c *gin.Context) {
	// TODO: Check DB connection before sending 200
	h.sendSuccess(c, nil)
}
