package goxorm

import (
	"time"
)

type SysLogPush struct {
	CreateTime time.Time `json:"create_time" xorm:"not null default 'CURRENT_TIMESTAMP' comment('创建时间') DATETIME"`
	Creator    string    `json:"creator" xorm:"not null default '' comment('执行人') VARCHAR(45)"`
	Id         int       `json:"id" xorm:"not null pk autoincr INT(11)"`
	Info       string    `json:"info" xorm:"not null comment('日志信息') TEXT"`
	Type       int       `json:"type" xorm:"not null comment('类型，1是字典下发日志2是开关下发') INT(1)"`
}
