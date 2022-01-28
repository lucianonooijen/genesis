package handlers

import (
	"github.com/gin-gonic/gin"

	"git.bytecode.nl/bytecode/genesis/server/internal/domains/user"
	"git.bytecode.nl/bytecode/genesis/server/internal/entities"
)

// PasswordResetStart starts the password reset procedure based on a given email.
func (h Handlers) PasswordResetStart(c *gin.Context) {
	var reqBody entities.PasswordResetStartRequest
	if failed := h.extractBody(c, &reqBody); failed {
		return
	}

	err := user.PasswordResetStart(h.services, reqBody.Email)
	if err != nil {
		h.handleDomainError(c, err)
		return
	}

	h.sendCreated(c, nil)
}

// PasswordResetComplete completes the password reset procedure.
func (h Handlers) PasswordResetComplete(c *gin.Context) {
	var reqBody entities.PasswordResetCompleteRequest
	if failed := h.extractBody(c, &reqBody); failed {
		return
	}

	res, err := user.PasswordResetComplete(h.services, reqBody.ResetToken, reqBody.Password)
	if err != nil {
		h.handleDomainError(c, err)
		return
	}

	h.sendSuccess(c, res)
}
