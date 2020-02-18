package goxorm

import (
	"time"
)

type SysCustom struct {
	Id          int       `json:"id" xorm:"not null pk autoincr INT(11)"`
	OrgId       int       `json:"org_id" xorm:"not null comment('机构id') unique(org_type) INT(11)"`
	Type        int       `json:"type" xorm:"not null comment('1代表机构logo配置项;2代表处方签logo;3代表化验报告logo;4代表病例logo;5代表医疗文书配置项') unique(org_type) TINYINT(4)"`
	UpdateTime  time.Time `json:"update_time" xorm:"not null default 'CURRENT_TIMESTAMP' DATETIME"`
	Updater     string    `json:"updater" xorm:"not null default '' VARCHAR(45)"`
	UpdaterName string    `json:"updater_name" xorm:"not null default '' VARCHAR(32)"`
	Used        int       `json:"used" xorm:"not null comment('1代表使用主品牌; 2代表使用子品牌; 3代表使用机构') TINYINT(4)"`
}
