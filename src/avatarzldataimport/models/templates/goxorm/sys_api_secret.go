package goxorm

import (
	"time"
)

type SysApiSecret struct {
	AddTime    time.Time `json:"add_time" xorm:"not null default 'CURRENT_TIMESTAMP' comment('创建时间') DATETIME"`
	ApiId      string    `json:"api_id" xorm:"not null default '' comment('api id') VARCHAR(32)"`
	ApiName    string    `json:"api_name" xorm:"not null default '' comment('api名称') VARCHAR(128)"`
	ApiSecret  string    `json:"api_secret" xorm:"not null default '' comment('api的secret') VARCHAR(32)"`
	Id         int       `json:"id" xorm:"not null pk autoincr comment('主键id') INT(11)"`
	LastModify time.Time `json:"last_modify" xorm:"not null default 'CURRENT_TIMESTAMP' comment('最后修改时间') DATETIME"`
	Status     int       `json:"status" xorm:"not null default 10 comment('状态 10 生效 20 实效') INT(3)"`
}
