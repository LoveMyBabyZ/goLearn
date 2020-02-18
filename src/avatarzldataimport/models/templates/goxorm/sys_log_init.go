package goxorm

import (
	"time"
)

type SysLogInit struct {
	CreateTime time.Time `json:"create_time" xorm:"not null default 'CURRENT_TIMESTAMP' comment('创建时间') DATETIME"`
	Creator    string    `json:"creator" xorm:"not null default '' comment('执行人') VARCHAR(45)"`
	Id         int       `json:"id" xorm:"not null pk autoincr INT(11)"`
	Info       string    `json:"info" xorm:"not null comment('日志信息') TEXT"`
	OrgId      int       `json:"org_id" xorm:"not null default 0 comment('公司id') INT(11)"`
	Type       int       `json:"type" xorm:"not null comment('类型') INT(1)"`
}
