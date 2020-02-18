package goxorm

import (
	"time"
)

type ErpCheckInventoryBillDetail struct {
	BookAmount        string    `json:"book_amount" xorm:"default 0.00000000 comment('账面金额') DECIMAL(18,8)"`
	BookQuantity      string    `json:"book_quantity" xorm:"default 0.00 comment('账面库存数量') DECIMAL(18,2)"`
	CheckBillId       int       `json:"check_bill_id" xorm:"default 0 comment('盘点单id') index INT(11)"`
	CheckQuantity     string    `json:"check_quantity" xorm:"default 0.00 comment('实盘数量') DECIMAL(18,2)"`
	CheckStatus       int       `json:"check_status" xorm:"default 0 comment('0：未盘点  1：已盘点') INT(11)"`
	DeficientAmount   string    `json:"deficient_amount" xorm:"default 0.00000000 comment('盘亏金额') DECIMAL(18,8)"`
	DifferentAmount   string    `json:"different_amount" xorm:"default 0.00000000 comment('差异金额') DECIMAL(18,8)"`
	DifferentQuantity string    `json:"different_quantity" xorm:"default 0.00 comment('差异数量') DECIMAL(18,2)"`
	DifferentReason   string    `json:"different_reason" xorm:"default '' comment('差异原因') VARCHAR(500)"`
	Id                int       `json:"id" xorm:"not null pk autoincr comment('id') INT(11)"`
	ItemId            int       `json:"item_id" xorm:"default 0 comment('品项Id') INT(11)"`
	ItemName          string    `json:"item_name" xorm:"default '' comment('品项名称') VARCHAR(100)"`
	Location1         string    `json:"location1" xorm:"default 0.00 comment('库位1') DECIMAL(18,2)"`
	Location2         string    `json:"location2" xorm:"default 0.00 comment('库位2') DECIMAL(18,2)"`
	Location3         string    `json:"location3" xorm:"default 0.00 comment('库位3') DECIMAL(18,2)"`
	Location4         string    `json:"location4" xorm:"default 0.00 comment('库位4') DECIMAL(18,2)"`
	Location5         string    `json:"location5" xorm:"default 0.00 comment('库位5') DECIMAL(18,2)"`
	RecheckQuantity   string    `json:"recheck_quantity" xorm:"default 0.00 comment('复盘数量') DECIMAL(18,2)"`
	ReferencePrice    string    `json:"reference_price" xorm:"default 0.00000000 comment('参考单价') DECIMAL(18,8)"`
	SurplusAmount     string    `json:"surplus_amount" xorm:"default 0.00000000 comment('盘盈金额') DECIMAL(18,8)"`
	UpdateTimestamp   time.Time `json:"update_timestamp" xorm:"not null default 'CURRENT_TIMESTAMP(6)' comment('更新时间戳') TIMESTAMP(6)"`
}
