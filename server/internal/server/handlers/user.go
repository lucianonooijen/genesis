package handlers

import (
	"github.com/gin-gonic/gin"

	"git.bytecode.nl/bytecode/genesis/internal/domains/user"
	"git.bytecode.nl/bytecode/genesis/internal/entities"
)

// CreateUser returns a 200 status code if everything is fine
func (h Handlers) CreateUser(c *gin.Context) {
	var reqBody entities.NewUserRequest
	if failed := h.extractBody(c, &reqBody); failed {
		return
	}

	res, err := user.CreateUser(h.services, reqBody) // FIXME sendDomainResult(interface{}, error) helper?
	if err != nil {
		h.handleDomainError(c, err)
		return
	}
	h.sendCreated(c, res)
}
