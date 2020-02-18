package goxorm

import (
	"time"
)

type SysHospitalDict struct {
	Code        int       `json:"code" xorm:"not null comment('医院端字典code->id') unique(code_org) INT(11)"`
	CreateTime  time.Time `json:"create_time" xorm:"not null default 'CURRENT_TIMESTAMP' comment('创建时间') DATETIME"`
	Creator     string    `json:"creator" xorm:"not null default '' comment('创建人id') VARCHAR(45)"`
	CreatorName string    `json:"creator_name" xorm:"not null default '' comment('创建人姓名') VARCHAR(32)"`
	DictType    int       `json:"dict_type" xorm:"not null comment('字典类型 1 总部端 2 医院端') INT(3)"`
	EnglishName string    `json:"english_name" xorm:"not null default '' comment('英文名->EnglishName') VARCHAR(128)"`
	Id          int       `json:"id" xorm:"not null pk autoincr comment('主键id') INT(11)"`
	ItemValue   string    `json:"item_value" xorm:"not null default '' VARCHAR(512)"`
	Level       int       `json:"level" xorm:"not null default 0 comment('管控级别 0 不管控 1 品牌管控') INT(4)"`
	Name        string    `json:"name" xorm:"not null default '' comment('医院端字典名称->Name') VARCHAR(128)"`
	Orders      int       `json:"orders" xorm:"not null default 0 comment('排序->Orders') INT(11)"`
	OrgId       int       `json:"org_id" xorm:"not null comment('医院id，或者品牌id') unique(code_org) INT(11)"`
	ParentCode  int       `json:"parent_code" xorm:"not null default 0 comment('父级code') INT(11)"`
	Path        string    `json:"path" xorm:"not null default '' comment('字典相关逻辑\n  * 需求梳理，数据结构构建校对\n  * 接口开发（doing）\n  * 接口文档整理\n递归路径') VARCHAR(1024)"`
	Status      int       `json:"status" xorm:"not null default 0 comment('字典状态  1启用 0不启用') INT(3)"`
	UpdateTime  time.Time `json:"update_time" xorm:"not null default 'CURRENT_TIMESTAMP' comment('最后修改时间') DATETIME"`
	Updater     string    `json:"updater" xorm:"not null default '' comment('更新人id') VARCHAR(45)"`
	UpdaterName string    `json:"updater_name" xorm:"not null default '' comment('更新人姓名') VARCHAR(32)"`
}
