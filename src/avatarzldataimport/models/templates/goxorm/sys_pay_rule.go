package goxorm

import (
	"time"
)

type SysPayRule struct {
	CreateTime     time.Time `json:"create_time" xorm:"comment('创建时间') DATETIME"`
	CreateUserId   int       `json:"create_user_id" xorm:"comment('创建人id') INT(11)"`
	CreateUserName string    `json:"create_user_name" xorm:"comment('创建人姓名') VARCHAR(50)"`
	OrgId          int       `json:"org_id" xorm:"not null comment('机构id') INT(11)"`
	PayType        int       `json:"pay_type" xorm:"not null default 1 comment('付款方式：1先本金后赠送、2等比例') INT(11)"`
}
