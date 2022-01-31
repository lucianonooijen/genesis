package handlers

import (
	"github.com/gin-gonic/gin"

	"git.bytecode.nl/bytecode/genesis/server/internal/domains/user"
	"git.bytecode.nl/bytecode/genesis/server/internal/entities"
)

// GetUserProfile gets the profile for the authenticated user.
func (h Handlers) GetUserProfile(c *gin.Context) {
	currentUser, ok := extractUserFromRequest(c)
	if !ok {
		return // error was already sent in extractUserFromRequest
	}

	profile, err := user.GetUserProfile(h.services, currentUser.ID)
	if err != nil {
		h.handleDomainError(c, err)
		return
	}

	h.sendSuccess(c, profile)
}

// UpdateUserProfile updates the profile for the authenticated user.
func (h Handlers) UpdateUserProfile(c *gin.Context) {
	currentUser, ok := extractUserFromRequest(c)
	if !ok {
		return // error was already sent in extractUserFromRequest
	}

	var reqBody entities.UserProfile
	if failed := h.extractBody(c, &reqBody); failed {
		return
	}

	profile, err := user.UpdateUserProfile(h.services, currentUser.ID, reqBody)
	if err != nil {
		h.handleDomainError(c, err)
		return
	}

	h.sendCreated(c, profile)
}

// DeleteAccount deletes the account.
func (h Handlers) DeleteAccount(c *gin.Context) {
	currentUser, ok := extractUserFromRequest(c)
	if !ok {
		return // error was already sent in extractUserFromRequest
	}

	var reqBody entities.DeleteAccountRequest
	if failed := h.extractBody(c, &reqBody); failed {
		return
	}

	err := user.DeleteAccount(h.services, currentUser.ID, reqBody.Password)
	if err != nil {
		h.handleDomainError(c, err)
		return
	}

	h.sendSuccess(c, nil)
}
