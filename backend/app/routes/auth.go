package routes

import (
	"app/app/middleware"
	"app/app/modules"

	"github.com/gin-gonic/gin"
)

func Auth(router *gin.RouterGroup) {
	module := modules.New()
	logger := middleware.NewLogResponse()
	auth := router.Group("")
	{
		auth.POST("/register", module.Auth.Ctl.Register)
		auth.POST("/login", logger, module.Auth.Ctl.Login)
		auth.POST("/forgot-password", module.Auth.Ctl.ForgotPassword)
		auth.POST("/verify-otp", module.Auth.Ctl.VerifyOtp)
	}
}
