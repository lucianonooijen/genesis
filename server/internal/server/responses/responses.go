package responses

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"git.bytecode.nl/bytecode/genesis/server/internal/constants"
)

// Responses contains methods for returning data in the correct format.
type Responses struct{}

// New returns a Responses instance.
func New() Responses {
	return Responses{}
}

func setResponseHeaders(c *gin.Context) {
	c.Header("X-Genesis-Server-Version", constants.APIVersion)
}

// Success sends 2XX responses.
func (r Responses) Success(c *gin.Context, code SuccessCode, data interface{}) {
	setResponseHeaders(c)
	c.JSON(int(code), data)
}

// ClientError returns a client error.
func (r Responses) ClientError(c *gin.Context, code ClientErrorCode, errTitle, errDetail string, err error) {
	setResponseHeaders(c)

	errorBody := ErrorBody{
		Title:  errTitle,
		Detail: errDetail,
		Status: int(code),
	}

	if err != nil {
		errorBody.RawError = err.Error()
	}

	c.JSON(int(code), errorBody)
}

// ServerError sends 500 responses.
func (r Responses) ServerError(c *gin.Context, err error) {
	setResponseHeaders(c)
	c.JSON(http.StatusInternalServerError, ErrorBody{
		Title:    "Unexpected server side error",
		Detail:   "There has been an unexpected server side error",
		Status:   http.StatusInternalServerError,
		RawError: err.Error(),
	})
}

// NotImplemented sends a http.StatusNotImplemented response.
func (r Responses) NotImplemented(c *gin.Context) {
	setResponseHeaders(c)
	c.JSON(http.StatusNotImplemented, ErrorBody{
		Title:  "Endpoint not implemented",
		Detail: "This API endpoint has not yet been implemented",
		Status: http.StatusNotImplemented,
	})
}
