package model

import (
	"time"

	"github.com/uptrace/bun"
)

type CreateUpdateUnixTimestamp struct {
	CreateUnixTimestamp
	UpdateUnixTimestamp
}

type CreateUnixTimestamp struct {
	CreatedAt int64 `json:"created_at" bun:",notnull,default:unix_timestamp()"`
}

type UpdateUnixTimestamp struct {
	UpdatedAt int64 `json:"updated_at" bun:",notnull,default:unix_timestamp()"`
}

type SoftDelete struct {
	DeletedAt *time.Time `json:"deleted_at" bun:",soft_delete,nullzero"`
}

// BUN/MariaDB specific base model
type BaseModel struct {
	bun.BaseModel `bun:"table:base_model"`
	ID            int64 `json:"id" bun:",pk,autoincrement"`
	CreateUpdateUnixTimestamp
	SoftDelete
}

func (t *CreateUnixTimestamp) SetCreated(ts int64) {
	t.CreatedAt = ts
}

func (t *CreateUnixTimestamp) SetCreatedNow() {
	t.SetCreated(time.Now().Unix())
}

func (t *UpdateUnixTimestamp) SetUpdate(ts int64) {
	t.UpdatedAt = ts
}

func (t *UpdateUnixTimestamp) SetUpdateNow() {
	t.SetUpdate(time.Now().Unix())
}
