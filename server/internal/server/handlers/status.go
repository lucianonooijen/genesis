package handlers

import "github.com/gin-gonic/gin"

// Status returns a 200 status code if everything is fine
func (h Handlers) Status(c *gin.Context) {
	pingErr := h.services.DbConn.Ping()
	if pingErr != nil {
		h.sendServerError(c, pingErr)
	}

	h.sendSuccess(c, nil)
}
