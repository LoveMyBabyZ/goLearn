package goxorm

import (
	"time"
)

type ErpAllocateApplicationsDetails struct {
	ActualQuantity   string    `json:"actual_quantity" xorm:"default 0.00 comment('实际数量') DECIMAL(18,2)"`
	ApplicationId    int       `json:"application_id" xorm:"default 0 comment('主表ID') INT(11)"`
	ApplyQuantity    string    `json:"apply_quantity" xorm:"default 0.00 comment('申请数量') DECIMAL(18,2)"`
	CategoryNamePath string    `json:"category_name_path" xorm:"default '' comment('完整的产品目录名称') VARCHAR(200)"`
	CommonName       string    `json:"common_name" xorm:"default '' comment('通用名') VARCHAR(100)"`
	Id               int       `json:"id" xorm:"not null pk autoincr comment('id') INT(11)"`
	ItemCode         string    `json:"item_code" xorm:"default '' comment('产品编码') VARCHAR(50)"`
	ItemName         string    `json:"item_name" xorm:"default '' comment('品项名称') VARCHAR(100)"`
	ItemUnitName     string    `json:"item_unit_name" xorm:"default '' comment('产品单位名称') VARCHAR(20)"`
	NuanIdApply      int       `json:"nuan_id_apply" xorm:"comment('申请方暧Id') INT(11)"`
	NuanIdSource     int       `json:"nuan_id_source" xorm:"comment('出库方暖Id') INT(11)"`
	ObjType          int       `json:"obj_type" xorm:"default 0 comment('项目类别(小暖)') INT(11)"`
	PackSpecific     string    `json:"pack_specific" xorm:"default '' comment('包装规格') VARCHAR(100)"`
	ProductType      int       `json:"product_type" xorm:"default 0 comment('项目类别') INT(11)"`
	ProductTypeName  string    `json:"product_type_name" xorm:"default '' comment('项目类别名称') VARCHAR(50)"`
	Remark           string    `json:"remark" xorm:"default '' comment('备注') VARCHAR(100)"`
	UpdateTimestamp  time.Time `json:"update_timestamp" xorm:"not null default 'CURRENT_TIMESTAMP(6)' comment('更新时间戳') TIMESTAMP(6)"`
}
