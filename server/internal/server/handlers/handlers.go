package handlers

import (
	"git.bytecode.nl/bytecode/genesis/internal/server/responses"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

// Handlers contains Gin request handlers as methods
type Handlers struct{}

// New returns Handlers instance
func New() (Handlers, error) {
	handlers := Handlers{}
	validate := validator.New()
	err := validate.Struct(handlers)
	return handlers, err
}

// Add responses to Handlers struct as private methods

var r = responses.New()
var s = responses.StatusCodes

func (h Handlers) sendSuccess(c *gin.Context, data interface{}) {
	h.checkResponseBody(c, data)
	r.Success(c, s.Success, data)
}
func (h Handlers) sendCreated(c *gin.Context, data interface{}) {
	h.checkResponseBody(c, data)
	r.Success(c, s.Created, data)
}
func (h Handlers) sendInvalidPostBody(c *gin.Context, err error) {
	r.ClientError(c, s.BadRequest, err.Error())
}
func (h Handlers) sendUnauthorized(c *gin.Context, err error) {
	r.ClientError(c, s.UnauthorizedRequest, err.Error())
}
func (h Handlers) sendNotFound(c *gin.Context, err error) {
	r.ClientError(c, s.NotFoundResponse, err.Error())
}
func (h Handlers) sendConflict(c *gin.Context, err error) {
	r.ClientError(c, s.Conflict, err.Error())
}
func (h Handlers) sendServerError(c *gin.Context, err error) {
	r.ServerError(c, err.Error())
}
