package handlers

import "github.com/gin-gonic/gin"

// Status returns a 200 status code if everything is fine, first pings database to see if there is connection.
func (h Handlers) Status(c *gin.Context) {
	pingErr := h.services.DBConn.Ping()
	if pingErr != nil {
		h.sendServerError(c, pingErr)
	}

	h.sendSuccess(c, nil)
}
