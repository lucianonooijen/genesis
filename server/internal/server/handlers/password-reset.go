package handlers

import (
	"github.com/gin-gonic/gin"

	"git.bytecode.nl/bytecode/genesis/server/internal/domains/user"
	"git.bytecode.nl/bytecode/genesis/server/internal/entities"
)

// PasswordResetStart starts the password reset procedure based on a given email.
// @Summary 	Begin password reset
// @Tags        User_PasswordReset
// @Description	Start password request by sending a reset token to a user's email address
// @Accept 		json
// @Produce     json
// @Param		account body entities.PasswordResetStartRequest true "AccountData"
// @Success		201
// @Failure		400 {object} responses.ErrorBody
// @Failure		401 {object} responses.ErrorBody
// @Failure		404 {object} responses.ErrorBody
// @Failure		409 {object} responses.ErrorBody
// @Failure		426 {object} responses.ErrorBody
// @Failure     500 {object} responses.ErrorBody
// @Router 		/user/password-reset/start [post]
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
// @Summary 	Complete password reset
// @Tags        User_PasswordReset
// @Description	Complete password request using reset token and new password
// @Accept 		json
// @Produce     json
// @Param		account body entities.PasswordResetCompleteRequest true "ResetData"
// @Success		200 {object} entities.JwtResponse
// @Failure		400 {object} responses.ErrorBody
// @Failure		401 {object} responses.ErrorBody
// @Failure		404 {object} responses.ErrorBody
// @Failure		409 {object} responses.ErrorBody
// @Failure		426 {object} responses.ErrorBody
// @Failure     500 {object} responses.ErrorBody
// @Router 		/user/password-reset/complete [post]
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
