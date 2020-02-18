package goxorm

import (
	"time"
)

type SysIssueHospitalSwitch struct {
	BrandId      string    `json:"brand_id" xorm:"not null default '1' comment('品牌id') unique(company_org_switch) VARCHAR(32)"`
	CreateTime   time.Time `json:"create_time" xorm:"not null default 'CURRENT_TIMESTAMP' comment('下发时间') DATETIME"`
	Creator      string    `json:"creator" xorm:"not null default '' comment('下发人') VARCHAR(45)"`
	HospitalName string    `json:"hospital_name" xorm:"not null default '' VARCHAR(32)"`
	Id           int       `json:"id" xorm:"not null pk autoincr INT(11)"`
	IsDelete     int       `json:"is_delete" xorm:"not null comment('软删除') INT(1)"`
	OrgId        string    `json:"org_id" xorm:"not null default '' comment('机构id，如果为0代表') unique(company_org_switch) VARCHAR(32)"`
	PushTime     time.Time `json:"push_time" xorm:"not null default 'CURRENT_TIMESTAMP' comment('下发操作时间') DATETIME"`
	Pusher       string    `json:"pusher" xorm:"not null default '' comment('操作人') VARCHAR(45)"`
	PusherName   string    `json:"pusher_name" xorm:"not null default '' comment('操作人姓名') VARCHAR(32)"`
	Status       int       `json:"status" xorm:"not null default 0 comment('是否启用0表示未启用 1表示启用') INT(1)"`
	SwitchId     int       `json:"switch_id" xorm:"not null comment('开关id') unique(company_org_switch) INT(11)"`
	UpdateTime   time.Time `json:"update_time" xorm:"not null default 'CURRENT_TIMESTAMP' comment('更新时间') DATETIME"`
	Updater      string    `json:"updater" xorm:"not null default '' VARCHAR(45)"`
}
