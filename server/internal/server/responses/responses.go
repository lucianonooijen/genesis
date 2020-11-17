package responses

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// Contains methods for returning data in the correct format
type Responses struct{}

func New() Responses {
	return Responses{}
}

// 2XX: Success
func (r Responses) Success(c *gin.Context, code SuccessCode, data interface{}) {
	c.JSON(int(code), generateResponseBody(true, nil, data))
}

// 4XX: Client errors
func (r Responses) ClientError(c *gin.Context, code ClientErrorCode, message string) {
	c.JSON(int(code), generateResponseBody(false, &message, nil))
}

// 5XX: Server errors
func (r Responses) ServerError(c *gin.Context, message string) {
	c.JSON(http.StatusInternalServerError, generateResponseBody(false, &message, nil))
}
func (r Responses) NotImplemented(c *gin.Context) {
	msg := "not implemented"
	c.JSON(http.StatusNotImplemented, generateResponseBody(false, &msg, nil))
}
