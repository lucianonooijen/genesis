package handlers

import (
	"github.com/gin-gonic/gin"

	"git.bytecode.nl/bytecode/genesis/server/internal/domains/user"
	"git.bytecode.nl/bytecode/genesis/server/internal/entities"
)

// CreateUser returns a 200 status code if everything is fine.
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

// LoginUser checks the credentials in the request and sends a JWT if they are valid.
func (h Handlers) LoginUser(c *gin.Context) {
	var reqBody entities.LoginRequest
	if failed := h.extractBody(c, &reqBody); failed {
		return
	}

	res, err := user.Login(h.services, reqBody)
	if err != nil {
		h.handleDomainError(c, err)
		return
	}

	h.sendSuccess(c, res)
}
