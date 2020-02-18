package goxorm

import (
	"time"
)

type SysIssueInstrument struct {
	CreateTime   time.Time `json:"create_time" xorm:"not null default 'CURRENT_TIMESTAMP' comment('创建时间') DATETIME"`
	Creator      string    `json:"creator" xorm:"not null default '' comment('创建人id') VARCHAR(45)"`
	CreatorName  string    `json:"creator_name" xorm:"not null default '' comment('创建人姓名') VARCHAR(32)"`
	Id           int       `json:"id" xorm:"not null pk autoincr INT(11)"`
	InstrumentId int       `json:"instrument_id" xorm:"not null comment('文书id') INT(11)"`
	OrgId        int       `json:"org_id" xorm:"not null comment('机构id') index INT(11)"`
	Status       int       `json:"status" xorm:"not null comment('1代表启用 0代表不启用') INT(1)"`
}
