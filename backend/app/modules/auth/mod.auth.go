package auth

import (
	"app/app/modules/user"
	"app/app/modules/userotp"

	"github.com/uptrace/bun"
)

type Module struct {
	Ctl *Controller
	Svc *Service
}

func NewModule(db *bun.DB, user *user.Module, userOtp *userotp.Module) *Module {
	svc := NewService(db, user, userOtp)
	return &Module{
		Ctl: NewController(svc),
		Svc: svc,
	}
}
