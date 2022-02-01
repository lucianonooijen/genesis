package server

import (
	"github.com/gin-gonic/gin"

	"git.bytecode.nl/bytecode/genesis/server/internal/constants"
	"git.bytecode.nl/bytecode/genesis/server/internal/server/handlers"
)

func registerRoutes(r *gin.RouterGroup, h handlers.Handlers) {
	// TODO: Use Embed.FS for this
	// https://github.com/gin-contrib/static/issues/19
	// STATIC FILES
	r.Static(constants.APIStaticPath, "./static")

	// STATUS
	r.GET("/status", h.Status)

	// USER LOGIN AND REGISTER
	r.POST("/user/register", h.CreateUser)
	r.POST("/user/login", h.LoginUser)

	// USER PASSWORD RESET
	r.POST("/user/password-reset/start", h.PasswordResetStart)
	r.POST("/user/password-reset/complete", h.PasswordResetComplete)

	// USER PROFILE
	r.GET("/user/profile", checkLoggedInAsUser, h.GetUserProfile)
	r.PUT("/user/profile", checkLoggedInAsUser, h.UpdateUserProfile)
	r.DELETE("/user/profile", checkLoggedInAsUser, h.DeleteAccount)
}
