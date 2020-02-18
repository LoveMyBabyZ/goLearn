package goxorm

import (
	"time"
)

type ErpInBillDetail struct {
	CategoryNamePath   string    `json:"category_name_path" xorm:"default '' comment('完整的产品目录名称') VARCHAR(200)"`
	CommonName         string    `json:"common_name" xorm:"comment('通用名') VARCHAR(100)"`
	Deviation          string    `json:"deviation" xorm:"default 0.00000000 comment('偏差') DECIMAL(18,8)"`
	ExpiryDate         time.Time `json:"expiry_date" xorm:"default 'CURRENT_TIMESTAMP' comment('截止日期(有效期限)') DATETIME"`
	Id                 int       `json:"id" xorm:"not null pk autoincr comment('id') INT(11)"`
	InAmount           string    `json:"in_amount" xorm:"default 0.00000000 comment('总金额') DECIMAL(18,8)"`
	InBillId           int       `json:"in_bill_id" xorm:"default 0 comment('入库单id') index INT(11)"`
	InPlanDetailId     int       `json:"in_plan_detail_id" xorm:"default 0 comment('入库计划单明细Id') INT(11)"`
	InPrice            string    `json:"in_price" xorm:"default 0.00000000 comment('入库单价') DECIMAL(18,8)"`
	InQuantity         string    `json:"in_quantity" xorm:"default 0.00 comment('入库总数量') DECIMAL(18,2)"`
	ItemId             int       `json:"item_id" xorm:"default 0 comment('品项Id') INT(11)"`
	ItemName           string    `json:"item_name" xorm:"default '' comment('品项名称') VARCHAR(100)"`
	ObjType            int       `json:"obj_type" xorm:"default 0 comment('项目类别(小暖)') INT(11)"`
	PackSpecific       string    `json:"pack_specific" xorm:"default '' comment('包装规格') VARCHAR(100)"`
	PresentPrice       string    `json:"present_price" xorm:"default 0.00000000 comment('赠品单价') DECIMAL(18,8)"`
	PresentQuantity    string    `json:"present_quantity" xorm:"default 0.00 comment('赠品入库数量') DECIMAL(18,2)"`
	ProductBatchNumber string    `json:"product_batch_number" xorm:"default '' comment('产品批号') VARCHAR(100)"`
	ProductType        int       `json:"product_type" xorm:"default 0 comment('项目类别') INT(11)"`
	ProductTypeName    string    `json:"product_type_name" xorm:"default '' comment('项目类别名称') VARCHAR(50)"`
	ProductionDate     time.Time `json:"production_date" xorm:"default 'CURRENT_TIMESTAMP' comment('生产日期') DATETIME"`
	Remark             string    `json:"remark" xorm:"default '' comment('备注') VARCHAR(500)"`
	ReturnAmount       string    `json:"return_amount" xorm:"default 0.00000000 comment('退货总金额') DECIMAL(18,8)"`
	ReturnQuantity     string    `json:"return_quantity" xorm:"default 0.00 comment('退货数量') DECIMAL(18,2)"`
	StoreUnitId        int       `json:"store_unit_id" xorm:"default 0 comment('入库单位') INT(11)"`
	StoreUnitName      string    `json:"store_unit_name" xorm:"default '' comment('入库单位名称') VARCHAR(50)"`
	UpdateTimestamp    time.Time `json:"update_timestamp" xorm:"not null default 'CURRENT_TIMESTAMP(6)' comment('更新时间戳') TIMESTAMP(6)"`
}
