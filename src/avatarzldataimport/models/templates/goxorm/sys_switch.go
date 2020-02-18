package goxorm

import (
	"time"
)

type SysSwitch struct {
	CreateTime  time.Time `json:"create_time" xorm:"not null default 'CURRENT_TIMESTAMP' comment('开关创建时间') DATETIME"`
	Description string    `json:"description" xorm:"not null comment('描述') TEXT"`
	Id          int       `json:"id" xorm:"not null pk autoincr INT(11)"`
	Name        string    `json:"name" xorm:"not null default '' comment('开关名称') VARCHAR(64)"`
	UpdateTime  time.Time `json:"update_time" xorm:"not null default 'CURRENT_TIMESTAMP' comment('开关更新时间') DATETIME"`
}
