package handlers

import (
	"github.com/gin-gonic/gin"

	"git.bytecode.nl/bytecode/genesis/server/internal/domains/user"
	"git.bytecode.nl/bytecode/genesis/server/internal/entities"
)

// RegisterPushNotificationToken is the handler for registering push notification tokens for users.
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
