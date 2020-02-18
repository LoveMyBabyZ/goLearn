package goxorm

import (
	"time"
)

type ErpOutBill struct {
	AuditTime      time.Time `json:"audit_time" xorm:"default 'CURRENT_TIMESTAMP' comment('审核时间') index(org_id) DATETIME"`
	AuditorId      int       `json:"auditor_id" xorm:"default 0 comment('审核人Id') INT(11)"`
	AuditorName    string    `json:"auditor_name" xorm:"default '' comment('审核人姓名') VARCHAR(50)"`
	CreateTime     time.Time `json:"create_time" xorm:"default 'CURRENT_TIMESTAMP' comment('创建时间(制单时间)') index(org_id) DATETIME"`
	CreateUserId   int       `json:"create_user_id" xorm:"default 0 comment('创建人id(制单人)') INT(11)"`
	CreateUserName string    `json:"create_user_name" xorm:"default '' comment('创建人姓名(制单人)') VARCHAR(50)"`
	Id             int       `json:"id" xorm:"not null pk autoincr comment('Id') INT(11)"`
	InstockOrgId   int       `json:"instock_org_id" xorm:"default 0 comment('入库机构ID') INT(11)"`
	InstockOrgName string    `json:"instock_org_name" xorm:"comment('入库机构名称') VARCHAR(255)"`
	ItemCount      int       `json:"item_count" xorm:"default 0 comment('总项数') INT(11)"`
	OrgId          int       `json:"org_id" xorm:"default 0 comment('机构id') index(org_id) INT(11)"`
	OutBillNumber  string    `json:"out_bill_number" xorm:"default '' comment('出库单号') unique VARCHAR(50)"`
	OutBillStatus  int       `json:"out_bill_status" xorm:"default 0 comment('出库单状态 
            0：保存 1：待审核  2：已审核  ') INT(11)"`
	OutBillType int `json:"out_bill_type" xorm:"default 0 comment('出库单据类型
            1. 过期出库  、2. 报废处理、3. 领用出库、4. 退货出库 5.销售出库单 6：盘亏出库单') index(org_id) INT(11)"`
	OutPlanId          int       `json:"out_plan_id" xorm:"default 0 comment('出库计划单Id') INT(11)"`
	RealTotalAmount    string    `json:"real_total_amount" xorm:"default 0.00000000 comment('实出总金额') DECIMAL(18,8)"`
	RelationBillId     int       `json:"relation_bill_id" xorm:"default 0 comment('相关单据ID') INT(11)"`
	RelationBillNumber string    `json:"relation_bill_number" xorm:"comment('相关单据号') VARCHAR(255)"`
	TotalAmount        string    `json:"total_amount" xorm:"default 0.00000000 comment('总金额') DECIMAL(18,8)"`
	UpdateTimestamp    time.Time `json:"update_timestamp" xorm:"not null default 'CURRENT_TIMESTAMP(6)' comment('更新时间戳') TIMESTAMP(6)"`
}
