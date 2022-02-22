package handlers

import (
	"github.com/gin-gonic/gin"

	"git.bytecode.nl/bytecode/genesis/server/internal/domains/user"
	"git.bytecode.nl/bytecode/genesis/server/internal/entities"
)

// CreateUser returns a 200 status code if everything is fine.
// @Summary 	Create user account
// @Tags        User_Auth
// @Description	Create a new user account
// @Accept      json
// @Produce     json
// @Param		user body entities.NewUserRequest true "New account data"
// @Success		201 {object} entities.JwtResponse
// @Failure		400 {object} responses.ErrorBody
// @Failure		401 {object} responses.ErrorBody
// @Failure		404 {object} responses.ErrorBody
// @Failure		409 {object} responses.ErrorBody
// @Failure		426 {object} responses.ErrorBody
// @Failure     500 {object} responses.ErrorBody
// @Router 		/user/register [post]
func (h Handlers) CreateUser(c *gin.Context) {
	var reqBody entities.NewUserRequest
	if failed := h.extractBody(c, &reqBody); failed {
		return
	}

	res, err := user.CreateUser(h.services, reqBody)
	if err != nil {
		h.handleDomainError(c, err)
		return
	}

	h.sendCreated(c, res)
}

// LoginUser checks the credentials in the request and sends a JWT if they are valid.
// @Summary 	Log into user account
// @Tags        User_Auth
// @Description	Log into user account
// @Accept      json
// @Produce     json
// @Param		user body entities.LoginRequest true "Username and password"
// @Success		200 {object} entities.JwtResponse
// @Failure		400 {object} responses.ErrorBody
// @Failure		401 {object} responses.ErrorBody
// @Failure		404 {object} responses.ErrorBody
// @Failure		409 {object} responses.ErrorBody
// @Failure		426 {object} responses.ErrorBody
// @Failure     500 {object} responses.ErrorBody
// @Router 		/user/login [post]
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
