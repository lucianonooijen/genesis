package handlers

import (
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"

	"git.bytecode.nl/bytecode/genesis/server/internal/interactors"
	"git.bytecode.nl/bytecode/genesis/server/internal/server/responses"
)

// Handlers contains Gin request handlers as methods.
type Handlers struct {
	services *interactors.Services
}

// New returns Handlers instance.
func New(services *interactors.Services) (Handlers, error) {
	handlers := Handlers{
		services: services,
	}
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
func (h Handlers) sendInvalidPostBody(c *gin.Context, err error) { //nolint:unused // this will be used when adding new features
	r.ClientError(c, s.BadRequest, "Invalid post body", "The received post body does not conform to the required structure", err)
}
func (h Handlers) sendUnauthorized(c *gin.Context, err error) { //nolint:unused // this will be used when adding new features
	r.ClientError(c, s.UnauthorizedRequest, "Unauthorized", "This endpoint requires you to be authenticated using the correct role", err)
}
func (h Handlers) sendForbidden(c *gin.Context, err error) { //nolint:unused // this will be used when adding new features
	r.ClientError(c, s.ForbiddenRequest, "Forbidden", "You don't have access to this resource", err)
}
func (h Handlers) sendNotFound(c *gin.Context, err error) { //nolint:unused // this will be used when adding new features
	r.ClientError(c, s.NotFoundResponse, "Not found", "The specified resource has not been found", err)
}
func (h Handlers) sendConflict(c *gin.Context, err error) { //nolint:unused // this will be used when adding new features
	r.ClientError(c, s.Conflict, "Conflict", "There is a conflict with the current state with the given resource", err)
}
func (h Handlers) sendServerError(c *gin.Context, err error) { //nolint:unused // this will be used when adding new features
	r.InternalServerError(c, err)
}
