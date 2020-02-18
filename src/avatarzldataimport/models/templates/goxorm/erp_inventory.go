package goxorm

import (
	"time"
)

type ErpInventory struct {
	AddTime                     time.Time `json:"add_time" xorm:"default 'CURRENT_TIMESTAMP' comment('系统添加时间') DATETIME"`
	CategoryNamePath            string    `json:"category_name_path" xorm:"comment('完整的产品目录名称') VARCHAR(200)"`
	CommonName                  string    `json:"common_name" xorm:"comment('通用名') VARCHAR(100)"`
	Deleted                     int       `json:"deleted" xorm:"default 0 comment('0: 未删除1：已删除') index(org_id) INT(11)"`
	Deviation                   string    `json:"deviation" xorm:"default 0.00000000 comment('偏差') DECIMAL(18,8)"`
	ExpiryDate                  time.Time `json:"expiry_date" xorm:"default 'CURRENT_TIMESTAMP' comment('截止日期(有效期限)') DATETIME"`
	Id                          int       `json:"id" xorm:"not null pk autoincr comment('id') INT(11)"`
	InBillDetailId              int       `json:"in_bill_detail_id" xorm:"default 0 comment('入库单明细id') index index(org_id) INT(11)"`
	InBillNumber                string    `json:"in_bill_number" xorm:"default '' comment('入库单号') index(org_id) VARCHAR(50)"`
	ItemId                      int       `json:"item_id" xorm:"default 0 comment('品项Id') index(org_id) INT(11)"`
	ItemName                    string    `json:"item_name" xorm:"default '' comment('品项名称') VARCHAR(100)"`
	ItemOriginalPresentQuantity string    `json:"item_original_present_quantity" xorm:"default 0.00 comment('原始赠品数量') DECIMAL(18,2)"`
	ItemOriginalQuantity        string    `json:"item_original_quantity" xorm:"default 0.00 comment('原始实际入库数量') DECIMAL(18,2)"`
	ItemPresentPrice            string    `json:"item_present_price" xorm:"default 0.00000000 comment('赠品单价') DECIMAL(18,8)"`
	ItemPresentQuantity         string    `json:"item_present_quantity" xorm:"default 0.00 comment('赠品剩余数量') DECIMAL(18,2)"`
	ItemPrice                   string    `json:"item_price" xorm:"default 0.00000000 comment('库存单价') DECIMAL(18,8)"`
	ItemQuantity                string    `json:"item_quantity" xorm:"default 0.00 comment('库存剩余数量') DECIMAL(18,2)"`
	ItemRealPresentQuantity     string    `json:"item_real_present_quantity" xorm:"default 0.00 comment('实际赠送库存剩余数据') DECIMAL(18,2)"`
	ItemRealQuantity            string    `json:"item_real_quantity" xorm:"default 0.00 comment('实际库存剩余数量') DECIMAL(18,2)"`
	ObjType                     int       `json:"obj_type" xorm:"comment('项目类别(小暖)') INT(11)"`
	OrgId                       int       `json:"org_id" xorm:"default 0 comment('机构Id') index(org_id) INT(11)"`
	OriginalDeviation           string    `json:"original_deviation" xorm:"default 0.00000000 comment('原始偏差') DECIMAL(18,8)"`
	PackSpecific                string    `json:"pack_specific" xorm:"comment('项目类别名称') VARCHAR(100)"`
	ProductBatchNumber          string    `json:"product_batch_number" xorm:"default '' comment('产品批号') VARCHAR(100)"`
	ProductType                 int       `json:"product_type" xorm:"comment('项目类别') INT(11)"`
	ProductTypeName             string    `json:"product_type_name" xorm:"comment('项目类别名称') VARCHAR(50)"`
	ProductionDate              time.Time `json:"production_date" xorm:"default 'CURRENT_TIMESTAMP' comment('生产日期') DATETIME"`
	StoreUnitId                 int       `json:"store_unit_id" xorm:"comment('入库单位') INT(11)"`
	StoreUnitName               string    `json:"store_unit_name" xorm:"comment('入库单位名称') VARCHAR(50)"`
	SupplierId                  int       `json:"supplier_id" xorm:"default 0 comment('供应商id') index(org_id) INT(11)"`
	SupplierName                string    `json:"supplier_name" xorm:"default '' comment('供应商名称') VARCHAR(100)"`
	UpdateTimestamp             time.Time `json:"update_timestamp" xorm:"not null default 'CURRENT_TIMESTAMP(6)' comment('更新时间戳') TIMESTAMP(6)"`
}
