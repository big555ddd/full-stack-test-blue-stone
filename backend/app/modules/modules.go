package modules

import (
	"app/app/modules/auth"
	"app/app/modules/user"
	"app/app/modules/userotp"
	"app/config"
)

type Controller struct {
	Auth    *auth.Module
	User    *user.Module
	UserOtp *userotp.Module
}

func New() *Controller {

	db := config.GetDB()

	user := user.NewModule(db)
	userOtp := userotp.NewModule(db)
	auth := auth.NewModule(db, user, userOtp)

	return &Controller{

		User:    user,
		Auth:    auth,
		UserOtp: userOtp,
	}
}
