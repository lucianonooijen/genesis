package handlers

import (
	"github.com/gin-gonic/gin"

	"git.bytecode.nl/bytecode/genesis/server/internal/domains/user"
	"git.bytecode.nl/bytecode/genesis/server/internal/entities"
)

// RegisterPushNotificationToken is the handler for registering push notification tokens for users.
// @Summary 	Register push notification token
// @Tags        User_PushNotifications
// @Description	Saves a push notification token for the authenticated user
// @Accept      json
// @Produce     json
// @Param		user body entities.PushNotificationRegister true "Token data"
// @Security	JWT_User
// @Success		201
// @Failure		400 {object} responses.ErrorBody
// @Failure		401 {object} responses.ErrorBody
// @Failure		404 {object} responses.ErrorBody
// @Failure		409 {object} responses.ErrorBody
// @Failure		426 {object} responses.ErrorBody
// @Failure     500 {object} responses.ErrorBody
// @Router 		/user/push-notifications [post]
func (h Handlers) RegisterPushNotificationToken(c *gin.Context) {
	reqUser, ok := extractUserFromRequest(c)
	if !ok {
		return // error was already sent in extractUserFromRequest
	}

	var pushdata entities.PushNotificationRegister
	if failed := h.extractBody(c, &pushdata); failed {
		return
	}

	if err := user.SaveUserPushNotificationToken(h.services, reqUser.ID, pushdata.Platform, pushdata.Token); err != nil {
		h.handleDomainError(c, err)
		return
	}

	h.sendCreated(c, nil)
}
