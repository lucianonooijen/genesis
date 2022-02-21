package handlers

import (
	"github.com/gin-gonic/gin"
)

// Status returns a 200 status code if everything is fine, first pings database to see if there is connection.
// @Summary 	Status handler
// @Tags        Common
// @Description	To be used with for status pings
// @Produce     json
// @Success		200
// @Failure     500
// @Router 		/status [get]
func (h Handlers) Status(c *gin.Context) {
	pingErr := h.services.DBConn.Ping()
	if pingErr != nil {
		h.sendServerError(c, pingErr)
	}

	h.sendSuccess(c, nil)
}

// CheckAppVersion is used to check the version of the mobile app using HTTP headers
// @Summary 	App version check handler
// @Tags        Common
// @Description	To be used with for checking app status based on HTTP version headers
// @Param		X-Genesis-Client-Version header string false "Client version code (semver)"
// @Produce     json
// @Success		200
// @Failure     500
// @Router 		/app-version [get]
func (h Handlers) CheckAppVersion(c *gin.Context) {
	pingErr := h.services.DBConn.Ping()
	if pingErr != nil {
		h.sendServerError(c, pingErr)
	}

	h.sendSuccess(c, nil)
}
