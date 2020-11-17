package responses

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// Responses contains methods for returning data in the correct format
type Responses struct{}

// New returns a Responses instance
func New() Responses {
	return Responses{}
}

// Success sends 2XX responses
func (r Responses) Success(c *gin.Context, code SuccessCode, data interface{}) {
	c.JSON(int(code), generateResponseBody(true, nil, data))
}

// ClientError sends 4XX responses
func (r Responses) ClientError(c *gin.Context, code ClientErrorCode, message string) {
	c.JSON(int(code), generateResponseBody(false, &message, nil))
}

// ServerError sends 500 responses
func (r Responses) ServerError(c *gin.Context, message string) {
	c.JSON(http.StatusInternalServerError, generateResponseBody(false, &message, nil))
}

// NotImplemented sends a http.StatusNotImplemented reponse
func (r Responses) NotImplemented(c *gin.Context) {
	msg := "not implemented"
	c.JSON(http.StatusNotImplemented, generateResponseBody(false, &msg, nil))
}
