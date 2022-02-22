package handlers

import (
	"github.com/gin-gonic/gin"

	"git.bytecode.nl/bytecode/genesis/server/internal/domains/user"
	"git.bytecode.nl/bytecode/genesis/server/internal/entities"
)

// GetUserProfile gets the profile for the authenticated user.
// @Summary 	Fetch user account
// @Tags        User_Account
// @Description	Fetches the user account for logged-in user
// @Produce     json
// @Security	JWT_User
// @Success		200 {object} entities.UserProfile
// @Failure		400 {object} responses.ErrorBody
// @Failure		401 {object} responses.ErrorBody
// @Failure		404 {object} responses.ErrorBody
// @Failure		409 {object} responses.ErrorBody
// @Failure		426 {object} responses.ErrorBody
// @Failure     500 {object} responses.ErrorBody
// @Router 		/user/profile [get]
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
// @Summary 	Update user account
// @Tags        User_Account
// @Description	Updates the user account for logged-in user
// @Accept 		json
// @Produce     json
// @Security	JWT_User
// @Param		profile body entities.UserProfile true "Profile"
// @Success		201 {object} entities.UserProfile
// @Failure		400 {object} responses.ErrorBody
// @Failure		401 {object} responses.ErrorBody
// @Failure		404 {object} responses.ErrorBody
// @Failure		409 {object} responses.ErrorBody
// @Failure		426 {object} responses.ErrorBody
// @Failure     500 {object} responses.ErrorBody
// @Router 		/user/profile [put]
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
// @Summary 	Delete user account
// @Tags        User_Account
// @Description	Deletes logged-in user's account
// @Accept 		json
// @Produce     json
// @Security	JWT_User
// @Param		profile body entities.DeleteAccountRequest true "Password"
// @Success		200
// @Failure		400 {object} responses.ErrorBody
// @Failure		401 {object} responses.ErrorBody
// @Failure		404 {object} responses.ErrorBody
// @Failure		409 {object} responses.ErrorBody
// @Failure		426 {object} responses.ErrorBody
// @Failure     500 {object} responses.ErrorBody
// @Router 		/user/profile [delete]
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
