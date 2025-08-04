package auth

import (
	"app/app/helper"
	"app/app/message"
	authdto "app/app/modules/auth/dto"
	"app/app/response"

	"github.com/gin-gonic/gin"
)

type Controller struct {
	Service *Service
}

func NewController(svc *Service) *Controller {
	return &Controller{
		Service: svc,
	}
}

func (ctl *Controller) Register(ctx *gin.Context) {
	req := authdto.RegisterRequest{}
	if err := ctx.Bind(&req); err != nil {
		response.BadRequest(ctx, err.Error(), nil)
		return
	}

	_, mserr, err := ctl.Service.Register(ctx, &req)
	if err != nil {
		ms := message.InternalServerError
		if mserr {
			ms = err.Error()
		}
		response.InternalError(ctx, ms, nil)
		return
	}
	response.Success(ctx, nil)
}

func (ctl *Controller) Login(ctx *gin.Context) {
	req := authdto.LoginRequest{}
	if err := ctx.Bind(&req); err != nil {
		response.BadRequest(ctx, err.Error(), nil)
		return
	}
	token, claim, err := ctl.Service.Login(ctx, &req)
	if err != nil {
		response.InternalError(ctx, err.Error(), nil)
		return
	}
	// set user in claims
	if claim != nil {
		helper.SetUserInClaims(ctx, claim)

	}
	response.Success(ctx, token)
}

func (ctl *Controller) ForgotPassword(ctx *gin.Context) {
	req := authdto.ForgotPasswordRequest{}
	if err := ctx.Bind(&req); err != nil {
		response.BadRequest(ctx, err.Error(), nil)
		return
	}

	data, err := ctl.Service.ForgotPassword(ctx, req.Email)
	if err != nil {
		response.InternalError(ctx, err.Error(), nil)
		return
	}
	response.Success(ctx, data)
}

func (ctl *Controller) VerifyOtp(ctx *gin.Context) {
	req := authdto.VerifyOtpRequest{}
	if err := ctx.Bind(&req); err != nil {
		response.BadRequest(ctx, err.Error(), nil)
		return
	}

	if req.NewPassword != req.ConfirmPassword {
		response.BadRequest(ctx, message.PasswordNotMatch, nil)
		return
	}

	err := ctl.Service.VerifyOtp(ctx, &req)
	if err != nil {
		response.InternalError(ctx, err.Error(), nil)
		return
	}
	response.Success(ctx, nil)
}
