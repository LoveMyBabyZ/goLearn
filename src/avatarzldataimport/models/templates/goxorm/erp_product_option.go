package goxorm

import (
	"time"
)

type ErpProductOption struct {
	Id                   int       `json:"id" xorm:"not null pk autoincr comment('id') INT(11)"`
	ItemId               int       `json:"item_id" xorm:"default 0 comment('品项Id') INT(11)"`
	ItemLasPurchasePrice string    `json:"item_las_purchase_price" xorm:"default 0.00000000 comment('最后一次采购入库价格') DECIMAL(18,8)"`
	OrgId                int       `json:"org_id" xorm:"default 0 comment('机构Id') INT(11)"`
	UpdateTimestamp      time.Time `json:"update_timestamp" xorm:"not null default 'CURRENT_TIMESTAMP(6)' comment('更新时间戳') TIMESTAMP(6)"`
}
