package auth

import (
	"app/app/helper"
	"app/app/message"
	"app/app/model"
	authdto "app/app/modules/auth/dto"
	"app/app/modules/user"
	userdto "app/app/modules/user/dto"
	"app/app/modules/userotp"
	userotpdto "app/app/modules/userotp/dto"
	"app/app/util/jwt"
	"app/config"
	"context"
	"database/sql"
	"errors"
	"time"

	"github.com/uptrace/bun"
	"golang.org/x/crypto/bcrypt"
)

type Service struct {
	db      *bun.DB
	user    *user.Module
	userOtp *userotp.Module
}

func NewService(db *bun.DB, user *user.Module, userOtp *userotp.Module) *Service {
	return &Service{
		db:      db,
		user:    user,
		userOtp: userOtp,
	}
}

func (s *Service) Register(ctx context.Context, req *authdto.RegisterRequest) (*model.User, bool, error) {
	emailExists, err := s.user.Svc.ExistEmail(ctx, req.Email)
	if err != nil {
		return nil, false, err
	}
	if emailExists {
		return nil, true, errors.New(message.EmailAlreadyExists)
	}
	usernameExists, err := s.user.Svc.ExistUserName(ctx, req.Username)
	if err != nil {
		return nil, false, err
	}
	if usernameExists {
		return nil, true, errors.New(message.UserAlreadyExists)
	}
	user, mserr, err := s.user.Svc.Create(ctx, userdto.CreateUser{
		Username: req.Username,
		Email:    req.Email,
		Password: req.Password,
	})

	return user, mserr, err
}

func (s *Service) Login(ctx context.Context, req *authdto.LoginRequest) (string, *jwt.Claims, error) {
	user, err := s.user.Svc.GetByUsername(ctx, req.Username)
	if err != nil {
		return "", nil, errors.New(message.InvalidCredentials)
	}

	data := jwt.ClaimData{
		ID:       user.ID,
		Username: user.UserName,
		Email:    user.Email,
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.Password)); err != nil {
		return "", nil, errors.New(message.InvalidCredentials)
	}

	token, claim, err := jwt.CreateToken(data)
	if err != nil {
		return "", nil, err
	}

	return token, claim, nil
}

func (s *Service) ForgotPassword(ctx context.Context, email string) (*model.UserOtp, error) {
	user, err := s.user.Svc.GetByEmail(ctx, email)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, errors.New(message.UserNotFound)
		}
		return nil, err
	}
	otp, err := helper.GenerateOTPCode(6)
	if err != nil {
		return nil, err
	}
	userotp := userotpdto.CreateUserOtp{
		UserID: user.ID,
		Otp:    otp,
	}

	data, err := s.userOtp.Svc.Create(ctx, userotp)
	if err != nil {
		return nil, err
	}

	//send OTP to user's email
	htmlBody := `
<!DOCTYPE html>
<html>
<head>
    <meta charset="UTF-8">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <title>BlueStone - Password Reset</title>
    <style>
        body { font-family: Arial, sans-serif; margin: 0; padding: 0; background-color: #f0f8ff; }
        .container { max-width: 600px; margin: 0 auto; background-color: #ffffff; }
        .header { background: linear-gradient(135deg, #1e3a8a, #3b82f6); padding: 30px; text-align: center; }
        .header h1 { color: #ffffff; margin: 0; font-size: 28px; font-weight: bold; }
        .content { padding: 40px 30px; }
        .otp-box { background: linear-gradient(135deg, #dbeafe, #bfdbfe); border: 2px solid #3b82f6; border-radius: 8px; padding: 20px; text-align: center; margin: 20px 0; }
        .otp-code { font-size: 32px; font-weight: bold; color: #1e3a8a; letter-spacing: 3px; margin: 10px 0; }
        .footer { background-color: #1e3a8a; color: #ffffff; text-align: center; padding: 20px; font-size: 14px; }
        .blue-text { color: #1e3a8a; }
    </style>
</head>
<body>
    <div class="container">
        <div class="header">
            <h1>BlueStone</h1>
        </div>
        <div class="content">
            <h2 class="blue-text">Password Reset Request</h2>
            <p>Hello,</p>
            <p>You have requested to reset your password. Please use the following OTP code to proceed:</p>
            <div class="otp-box">
                <p style="margin: 0; font-size: 16px; color: #1e3a8a;">Your OTP Code:</p>
                <div class="otp-code">` + otp + `</div>
                <p style="margin: 0; font-size: 14px; color: #6b7280;">This code will expire in 15 minutes</p>
            </div>
            <p>If you didn't request this password reset, please ignore this email.</p>
            <p>Best regards,<br><strong class="blue-text">BlueStone Team</strong></p>
        </div>
        <div class="footer">
            <p>&copy; 2025 BlueStone. All rights reserved.</p>
        </div>
    </div>
</body>
</html>`

	err = config.SendEmail(user.Email, "BlueStone - Password Reset", "Your OTP Code", htmlBody)
	if err != nil {
		return nil, err
	}

	return data, nil
}

func (s *Service) VerifyOtp(ctx context.Context, req *authdto.VerifyOtpRequest) error {
	otp, err := s.userOtp.Svc.Get(ctx, req.ID)
	if err != nil {
		return err
	}

	if otp.ExpiresAt < time.Now().Unix() {
		return errors.New(message.OTPExpired)
	}

	if otp.Used {
		return errors.New(message.OTPAlreadyUsed)
	}

	if otp.Otp != req.Otp {
		return errors.New(message.OTPInvalid)
	}

	err = s.user.Svc.UpdatePassword(ctx, otp.UserID, req.NewPassword)
	if err != nil {
		return err
	}

	err = s.userOtp.Svc.UpdateUsed(ctx, req.ID)
	if err != nil {
		return err
	}

	return nil
}
