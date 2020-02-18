package goxorm

import (
	"time"
)

type ErpInPlanDetail struct {
	CategoryNamePath          string    `json:"category_name_path" xorm:"comment('完整的产品目录名称') VARCHAR(200)"`
	Id                        int       `json:"id" xorm:"not null pk autoincr comment('Id') INT(11)"`
	InPlanId                  int       `json:"in_plan_id" xorm:"not null default 0 comment('入库计划单Id') INT(11)"`
	InstoredAmount            string    `json:"instored_amount" xorm:"not null default 0.00000000 comment('已入库金额') DECIMAL(18,8)"`
	InstoredQuantity          string    `json:"instored_quantity" xorm:"not null default 0.00 comment('已入库数（入库单位）') DECIMAL(18,2)"`
	IsValidityControl         int       `json:"is_validity_control" xorm:"default 0 comment('是否效期管理') TINYINT(1)"`
	ItemId                    int       `json:"item_id" xorm:"not null default 0 comment('品项Id') INT(11)"`
	ItemName                  string    `json:"item_name" xorm:"default '' comment('品项名称') VARCHAR(100)"`
	ObjType                   int       `json:"obj_type" xorm:"comment('项目类别(小暖)') INT(11)"`
	ObjTypeName               string    `json:"obj_type_name" xorm:"VARCHAR(50)"`
	PackSpecific              string    `json:"pack_specific" xorm:"comment('包装规格') VARCHAR(100)"`
	PresentQuantity           string    `json:"present_quantity" xorm:"not null default 0.00 comment('赠品数量') DECIMAL(18,2)"`
	ProductType               int       `json:"product_type" xorm:"comment('项目类别') INT(11)"`
	ProductTypeName           string    `json:"product_type_name" xorm:"comment('项目类别名称') VARCHAR(50)"`
	PurchasePrice             string    `json:"purchase_price" xorm:"not null default 0.00000000 comment('采购单价') DECIMAL(18,8)"`
	PurchaseQuantity          string    `json:"purchase_quantity" xorm:"not null default 0.00 comment('采购数量') DECIMAL(18,2)"`
	PurchaseToStoreConversion string    `json:"purchase_to_store_conversion" xorm:"not null default 0.00 comment('采购单位到库存单位换算值') DECIMAL(18,2)"`
	PurchaseUnit              int       `json:"purchase_unit" xorm:"not null default 0 comment('采购单位') INT(11)"`
	PurchaseUnitName          string    `json:"purchase_unit_name" xorm:"not null default '' comment('采购单位名称') VARCHAR(50)"`
	RefPrice                  string    `json:"ref_price" xorm:"default 0.00 comment('入库含税参考价') DECIMAL(18,2)"`
	Remark                    string    `json:"remark" xorm:"default '' comment('备注') VARCHAR(500)"`
	StoreUnitId               int       `json:"store_unit_id" xorm:"not null default 0 comment('入库单位') INT(11)"`
	StoreUnitName             string    `json:"store_unit_name" xorm:"not null default '' comment('入库单位名称') VARCHAR(50)"`
	TotalAmount               string    `json:"total_amount" xorm:"comment('总金额') DECIMAL(18,8)"`
	UpdateTimestamp           time.Time `json:"update_timestamp" xorm:"not null default 'CURRENT_TIMESTAMP(6)' comment('更新时间戳') TIMESTAMP(6)"`
}
