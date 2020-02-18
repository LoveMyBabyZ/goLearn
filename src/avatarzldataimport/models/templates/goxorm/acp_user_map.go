package goxorm

import (
	"time"
)

type AcpUserMap struct {
	AddTime      time.Time `json:"add_time" xorm:"default 'CURRENT_TIMESTAMP' comment('新增时间') DATETIME"`
	Id           int       `json:"id" xorm:"not null pk autoincr INT(10)"`
	LastModify   time.Time `json:"last_modify" xorm:"default 'CURRENT_TIMESTAMP' comment('修改时间') DATETIME"`
	NuanId       int       `json:"nuan_id" xorm:"not null default 0 comment('小暖id') INT(10)"`
	Phone        string    `json:"phone" xorm:"not null default '' comment('手机号码') VARCHAR(45)"`
	UserId       int       `json:"user_id" xorm:"not null default 0 comment('ACP用户id') INT(10)"`
	UserName     string    `json:"user_name" xorm:"not null default '' comment('用户名，可用于登录') VARCHAR(45)"`
	UserRealName string    `json:"user_real_name" xorm:"not null default '' comment('用户真实姓名') VARCHAR(45)"`
}
