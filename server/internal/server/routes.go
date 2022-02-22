package server

import (
	"io/fs"
	"net/http"

	"github.com/gin-gonic/gin"

	"git.bytecode.nl/bytecode/genesis/server/internal/constants"
	"git.bytecode.nl/bytecode/genesis/server/internal/server/handlers"
	"git.bytecode.nl/bytecode/genesis/server/static"
)

// NOTE: All final handlers should have Swagger docs present,
// see: https://github.com/swaggo/swag

func registerRoutes(r *gin.RouterGroup, h handlers.Handlers) {
	// STATIC FILES
	r.StaticFS(constants.APIStaticPath, mustFS())

	// STATUS AND VERSION CHECK ROUTES
	r.GET("/status", h.Status)
	r.GET("/app-version", h.CheckAppVersion)

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

	// PUSH NOTIFICATIONS
	r.POST("/user/push-notifications", checkLoggedInAsUser, h.RegisterPushNotificationToken)
}

func mustFS() http.FileSystem {
	sub, err := fs.Sub(static.FS, ".")

	if err != nil {
		panic(err)
	}

	return http.FS(sub)
}
