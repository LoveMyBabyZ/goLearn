package goxorm

import (
	"time"
)

type SysCompanyDict struct {
	BrandId     int       `json:"brand_id" xorm:"not null comment('品牌id') unique(code_brand) INT(11)"`
	Code        int       `json:"code" xorm:"not null comment('总部端字典code->id') unique(code_brand) INT(11)"`
	CreateTime  time.Time `json:"create_time" xorm:"not null default 'CURRENT_TIMESTAMP' comment('创建时间') DATETIME"`
	Creator     string    `json:"creator" xorm:"not null default '' comment('创建人id') VARCHAR(45)"`
	CreatorName string    `json:"creator_name" xorm:"not null default '' comment('创建人姓名') VARCHAR(32)"`
	EnglishName string    `json:"english_name" xorm:"not null default '' comment('英文名->EnglishName') VARCHAR(500)"`
	Id          int       `json:"id" xorm:"not null pk autoincr comment('主键id') INT(11)"`
	ItemValue   string    `json:"item_value" xorm:"not null default '' VARCHAR(512)"`
	Name        string    `json:"name" xorm:"not null default '' comment('总部端字典名称->Name') VARCHAR(128)"`
	Orders      int       `json:"orders" xorm:"not null default 0 comment('顺序->Orders') INT(11)"`
	ParentCode  int       `json:"parent_code" xorm:"not null default 0 comment('父级code') INT(11)"`
	Path        string    `json:"path" xorm:"not null default '' comment('递归路径') VARCHAR(1024)"`
	Status      int       `json:"status" xorm:"not null default 0 comment('状态 1启用 0不启用') INT(3)"`
	UpdateTime  time.Time `json:"update_time" xorm:"not null default 'CURRENT_TIMESTAMP' comment('修改时间') DATETIME"`
	Updater     string    `json:"updater" xorm:"not null default '' comment('修改人id') VARCHAR(45)"`
	UpdaterName string    `json:"updater_name" xorm:"not null default '' comment('修改人姓名') VARCHAR(32)"`
}
