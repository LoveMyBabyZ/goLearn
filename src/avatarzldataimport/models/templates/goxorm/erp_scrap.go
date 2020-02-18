package goxorm

import (
	"time"
)

type ErpScrap struct {
	Id              int       `json:"id" xorm:"not null pk autoincr comment('id') INT(11)"`
	ItemId          int       `json:"item_id" xorm:"default 0 comment('品项Id') index(org_id) INT(11)"`
	ItemName        string    `json:"item_name" xorm:"default '' comment('品项名称') index(org_id) VARCHAR(100)"`
	LastUpdateTime  time.Time `json:"last_update_time" xorm:"default 'CURRENT_TIMESTAMP' comment('最后更新时间') index(org_id) DATETIME"`
	OrgId           int       `json:"org_id" xorm:"default 0 comment('机构id') index(org_id) INT(11)"`
	OutBillDetailId int       `json:"out_bill_detail_id" xorm:"default 0 comment('出库单明细Id') INT(11)"`
	OutDate         time.Time `json:"out_date" xorm:"default 'CURRENT_TIMESTAMP' comment('出库日期') index(org_id) DATETIME"`
	OutQuantity     string    `json:"out_quantity" xorm:"default 0.00 comment('出库数量') DECIMAL(18,2)"`
	ScrapQuantity   string    `json:"scrap_quantity" xorm:"default 0.00 comment('已报废数量') DECIMAL(18,2)"`
	UpdateTimestamp time.Time `json:"update_timestamp" xorm:"not null default 'CURRENT_TIMESTAMP(6)' comment('更新时间戳') TIMESTAMP(6)"`
	UpdateUserId    int       `json:"update_user_id" xorm:"default 0 comment('更新人') INT(11)"`
	UpdateUserName  string    `json:"update_user_name" xorm:"default '' comment('更新人姓名') VARCHAR(50)"`
}
