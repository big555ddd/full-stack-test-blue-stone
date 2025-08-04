package model

import "github.com/uptrace/bun"

type ActivityLog struct {
	bun.BaseModel `bun:"table:activity_logs"`

	ID         string      `json:"id" bun:",pk,type:varchar(36),default:uuid()" form:"id"`
	Section    string      `json:"section" form:"section"`
	EventType  string      `json:"event_type" bun:",notnull,type:varchar(50)" form:"event_type"`
	StatusCode int         `json:"status_code" bun:",notnull" form:"status_code"`
	Parameters interface{} `json:"parameters" bun:"type:json" form:"parameters"`
	Responses  interface{} `json:"responses" bun:"type:json" form:"responses"`
	Query      interface{} `json:"query" bun:"type:json" form:"query"`
	IpAddress  string      `json:"ip_address" bun:",notnull,type:varchar(50)" form:"ip_address"`
	UserAgent  string      `json:"user_agent" form:"user_agent"`
	CreatedBy  string      `json:"created_by" form:"created_by"`
	CreatedAt  int64       `json:"created_at" bun:"default:unix_timestamp()" form:"created_at"`
}
