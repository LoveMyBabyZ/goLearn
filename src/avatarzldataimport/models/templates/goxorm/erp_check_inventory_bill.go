package goxorm

import (
	"time"
)

type ErpCheckInventoryBill struct {
	AuditTime       time.Time `json:"audit_time" xorm:"default 'CURRENT_TIMESTAMP' comment('审核时间') index(org_id) DATETIME"`
	AuditorId       int       `json:"auditor_id" xorm:"default 0 comment('审核人Id') INT(11)"`
	AuditorName     string    `json:"auditor_name" xorm:"default '' comment('审核人姓名') VARCHAR(50)"`
	BookTotalAmount string    `json:"book_total_amount" xorm:"default 0.00000000 comment('账面总金额') DECIMAL(18,8)"`
	CheckBillNumber string    `json:"check_bill_number" xorm:"default '' comment('盘点单号') unique VARCHAR(50)"`
	CheckBillStatus int       `json:"check_bill_status" xorm:"default 0 comment('盘点单状态 
 0：保存 1：盘点中 2：待盘差 3：盘差中  4：待审核  5：已审核 ') INT(11)"`
	CheckDate              time.Time `json:"check_date" xorm:"default 'CURRENT_TIMESTAMP' comment('盘点日期') DATETIME"`
	CheckDifferenceDate    time.Time `json:"check_difference_date" xorm:"comment('盘差日期') DATETIME"`
	CheckResultTotalAmount string    `json:"check_result_total_amount" xorm:"default 0.00000000 comment('实际总金额') DECIMAL(18,8)"`
	CreateTime             time.Time `json:"create_time" xorm:"default 'CURRENT_TIMESTAMP' comment('创建时间(制单时间)') index(org_id) DATETIME"`
	CreateUserId           int       `json:"create_user_id" xorm:"default 0 comment('创建人id(制单人)') INT(11)"`
	CreateUserName         string    `json:"create_user_name" xorm:"default '' comment('创建人姓名(制单人)') VARCHAR(50)"`
	Id                     int       `json:"id" xorm:"not null pk autoincr comment('Id') INT(11)"`
	ItemCount              int       `json:"item_count" xorm:"default 0 comment('总项数') INT(11)"`
	OrgId                  int       `json:"org_id" xorm:"default 0 comment('机构id') index(org_id) INT(11)"`
	UpdateTimestamp        time.Time `json:"update_timestamp" xorm:"not null default 'CURRENT_TIMESTAMP(6)' comment('更新时间戳') TIMESTAMP(6)"`
}
