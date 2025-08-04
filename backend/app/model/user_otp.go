package model

import "github.com/uptrace/bun"

type UserOtp struct {
	bun.BaseModel `bun:"table:user_otps"`

	ID        string `json:"id" bun:",pk,type:varchar(36),default:uuid()" form:"id"`
	UserID    string `json:"user_id" bun:",notnull,type:varchar(36)" form:"user_id"`
	Otp       string `json:"otp" bun:",notnull,type:varchar(60)" form:"otp"`
	ExpiresAt int64  `json:"expires_at" bun:",notnull" form:"expires_at"`
	Used      bool   `json:"used" bun:",notnull,default:false" form:"used"`
	CreateUnixTimestamp
}
