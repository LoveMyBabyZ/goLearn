package goxorm

import (
	"time"
)

type SysPushedOrg struct {
	CreateTime time.Time `json:"create_time" xorm:"not null default 'CURRENT_TIMESTAMP' comment('下发时间') DATETIME"`
	Id         int       `json:"id" xorm:"not null pk autoincr INT(11)"`
	OrgId      string    `json:"org_id" xorm:"not null default '' comment('机构id') VARCHAR(32)"`
	Type       int       `json:"type" xorm:"not null default 0 comment('下发的类型1表示字典，2表示开关') INT(1)"`
}
