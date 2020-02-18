package goxorm

import (
	"time"
)

type ErpOutPlanDetail struct {
	Id                int       `json:"id" xorm:"not null pk autoincr comment('Id') INT(11)"`
	ItemId            int       `json:"item_id" xorm:"not null default 0 comment('品项Id') INT(11)"`
	ItemName          string    `json:"item_name" xorm:"default '' comment('品项名称') VARCHAR(100)"`
	OutPlanId         int       `json:"out_plan_id" xorm:"not null default 0 comment('出库计划单Id') INT(11)"`
	OutQuantity       string    `json:"out_quantity" xorm:"not null default 0.00 comment('出库数量') DECIMAL(18,2)"`
	OutstoredQuantity string    `json:"outstored_quantity" xorm:"not null default 0.00 comment('已出库数（出库单位）') DECIMAL(18,2)"`
	Remark            string    `json:"remark" xorm:"not null default '' comment('备注') VARCHAR(500)"`
	StoreUnitId       int       `json:"store_unit_id" xorm:"not null default 0 comment('出库单位') INT(11)"`
	StoreUnitName     string    `json:"store_unit_name" xorm:"not null default '' comment('出库单位名称') VARCHAR(50)"`
	UpdateTimestamp   time.Time `json:"update_timestamp" xorm:"not null default 'CURRENT_TIMESTAMP(6)' comment('更新时间戳') TIMESTAMP(6)"`
}
