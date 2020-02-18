package goxorm

import (
	"time"
)

type SysPrivateDict struct {
	Code        int       `json:"code" xorm:"not null comment('医院端字典code') INT(11)"`
	CreateTime  time.Time `json:"create_time" xorm:"not null default 'CURRENT_TIMESTAMP' comment('创建时间') DATETIME"`
	Creator     string    `json:"creator" xorm:"not null default '0' comment('创建人id') VARCHAR(45)"`
	CreatorName string    `json:"creator_name" xorm:"not null default '' comment('创建人姓名') VARCHAR(32)"`
	Id          int       `json:"id" xorm:"not null pk autoincr comment('主键id') INT(11)"`
	OrgId       int       `json:"org_id" xorm:"not null comment('组织id') INT(11)"`
	Status      int       `json:"status" xorm:"not null default 0 comment('字典状态 1代表启用 0代表不启用') INT(1)"`
	UpdateTime  time.Time `json:"update_time" xorm:"not null default 'CURRENT_TIMESTAMP' comment('最后修改时间') DATETIME"`
	Updater     string    `json:"updater" xorm:"not null default '0' comment('更新人id') VARCHAR(45)"`
	UpdaterName string    `json:"updater_name" xorm:"not null default '' comment('更新人姓名') VARCHAR(32)"`
}
