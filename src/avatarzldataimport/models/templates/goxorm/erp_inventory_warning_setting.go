package goxorm

import (
	"time"
)

type ErpInventoryWarningSetting struct {
	CategoryNamePath string    `json:"category_name_path" xorm:"comment('完整的产品目录名称') VARCHAR(200)"`
	CreateTime       time.Time `json:"create_time" xorm:"comment('创建时间') DATETIME"`
	CreateUserId     int       `json:"create_user_id" xorm:"comment('创建人id') INT(11)"`
	CreateUserName   string    `json:"create_user_name" xorm:"comment('创建人姓名') VARCHAR(50)"`
	Id               int       `json:"id" xorm:"not null pk autoincr comment('Id') INT(11)"`
	ItemId           int       `json:"item_id" xorm:"comment('产品id') INT(11)"`
	ItemName         string    `json:"item_name" xorm:"comment('产品名称') VARCHAR(100)"`
	ObjType          int       `json:"obj_type" xorm:"comment('项目类别(小暖)') INT(11)"`
	OrgId            int       `json:"org_id" xorm:"comment('机构id') INT(11)"`
	PackSpecific     string    `json:"pack_specific" xorm:"comment('包装规格') VARCHAR(100)"`
	ProductType      int       `json:"product_type" xorm:"comment('项目类别') INT(11)"`
	ProductTypeName  string    `json:"product_type_name" xorm:"comment('项目类别名称') VARCHAR(50)"`
	SafetyType       int       `json:"safety_type" xorm:"default 1 comment('库存安全基准：1、库存可用数量 2、库存总数量') INT(11)"`
	SafetyValue      int       `json:"safety_value" xorm:"comment('库存安全值') INT(11)"`
	Status           int       `json:"status" xorm:"default 1 comment('状态：1、启用 2、关闭') INT(11)"`
	StoreUnitId      int       `json:"store_unit_id" xorm:"comment('入库单位') INT(11)"`
	StoreUnitName    string    `json:"store_unit_name" xorm:"comment('入库单位名称') VARCHAR(50)"`
	UpdateTimestamp  time.Time `json:"update_timestamp" xorm:"not null default 'CURRENT_TIMESTAMP(6)' comment('更新时间戳') TIMESTAMP(6)"`
	UpdateUserId     int       `json:"update_user_id" xorm:"comment('修改人id') INT(11)"`
	UpdateUserName   string    `json:"update_user_name" xorm:"comment('修改人姓名') VARCHAR(50)"`
}
