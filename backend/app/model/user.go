package model

import (
	"github.com/uptrace/bun"
)

type User struct {
	bun.BaseModel `bun:"table:users"`

	ID       string `bun:",pk,type:varchar(36),default:uuid()" json:"id"`
	UserName string `bun:"username,unique,notnull" json:"username"`
	Email    string `bun:"email,unique,notnull" json:"email"`
	Password string `bun:"password,notnull" json:"password"`

	CreateUpdateUnixTimestamp
	SoftDelete
}
