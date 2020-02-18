package goxorm

import (
	"time"
)

type ErpOutBillDetailBatch struct {
	Deviation          string    `json:"deviation" xorm:"default 0.00000000 comment('偏差') DECIMAL(18,8)"`
	ExpiryDate         time.Time `json:"expiry_date" xorm:"default 'CURRENT_TIMESTAMP' comment('截止日期(有效期限)') DATETIME"`
	Id                 int       `json:"id" xorm:"not null pk autoincr comment('id') INT(11)"`
	InventoryId        int       `json:"inventory_id" xorm:"default 0 comment('库存Id(根据此Id锁库存)') INT(11)"`
	ItemId             int       `json:"item_id" xorm:"default 0 comment('品项Id') INT(11)"`
	ItemName           string    `json:"item_name" xorm:"default '' comment('品项名称') VARCHAR(100)"`
	OutAmount          string    `json:"out_amount" xorm:"default 0.00000000 comment('出库总金额（包含赠品）') DECIMAL(18,8)"`
	OutBillDetailId    int       `json:"out_bill_detail_id" xorm:"default 0 comment('出库计划单明细Id') INT(11)"`
	OutBillId          int       `json:"out_bill_id" xorm:"default 0 comment('出库单id') index INT(11)"`
	OutPrice           string    `json:"out_price" xorm:"default 0.00000000 comment('出库单价') DECIMAL(18,8)"`
	OutQuantity        string    `json:"out_quantity" xorm:"default 0.00 comment('出库总数量') DECIMAL(18,2)"`
	PresentPrice       string    `json:"present_price" xorm:"default 0.00000000 comment('赠品单价') DECIMAL(18,8)"`
	PresentQuantity    string    `json:"present_quantity" xorm:"default 0.00 comment('赠品出库数量') DECIMAL(18,2)"`
	ProductBatchNumber string    `json:"product_batch_number" xorm:"default '' comment('产品批号') VARCHAR(100)"`
	ProductionDate     time.Time `json:"production_date" xorm:"default 'CURRENT_TIMESTAMP' comment('生产日期') DATETIME"`
	Remark             string    `json:"remark" xorm:"default '' comment('备注') VARCHAR(500)"`
	UpdateTimestamp    time.Time `json:"update_timestamp" xorm:"not null default 'CURRENT_TIMESTAMP(6)' comment('更新时间戳') TIMESTAMP(6)"`
}
