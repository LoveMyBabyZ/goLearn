package goxorm

import (
	"time"
)

type ErpOutPlan struct {
	AuditTime      time.Time `json:"audit_time" xorm:"default 'CURRENT_TIMESTAMP' comment('审核时间') DATETIME"`
	AuditorId      int       `json:"auditor_id" xorm:"default 0 comment('审核人Id') INT(11)"`
	AuditorName    string    `json:"auditor_name" xorm:"default '' comment('审核人姓名') VARCHAR(50)"`
	CreateTime     time.Time `json:"create_time" xorm:"default 'CURRENT_TIMESTAMP' comment('创建时间(制单时间)') DATETIME"`
	CreateUserId   int       `json:"create_user_id" xorm:"default 0 comment('创建人id(制单人)') INT(11)"`
	CreateUserName string    `json:"create_user_name" xorm:"default '' comment('创建人姓名(制单人)') VARCHAR(50)"`
	Id             int       `json:"id" xorm:"not null pk autoincr comment('id') INT(11)"`
	InBillId       int       `json:"in_bill_id" xorm:"default 0 comment('对应的采购入库单Id') INT(11)"`
	InBillNumber   string    `json:"in_bill_number" xorm:"default '' comment('对应的采购入库单号') VARCHAR(50)"`
	ItemCount      int       `json:"item_count" xorm:"default 0 comment('总项数') INT(11)"`
	OrderType      int       `json:"order_type" xorm:"default 0 comment('单据类型') INT(11)"`
	OrgId          int       `json:"org_id" xorm:"default 0 comment('机构Id') INT(11)"`
	OutBillNumber  string    `json:"out_bill_number" xorm:"default '' comment('对应出库单据号') VARCHAR(500)"`
	OutPlanNumber  string    `json:"out_plan_number" xorm:"default '' comment('出库计划单号') VARCHAR(50)"`
	OutPlanStatus  int       `json:"out_plan_status" xorm:"default 0 comment('单据状态  
            0:保存 1：待审核 2:已审核  3：已审核-入库中 4：已审核-已完成  -1：已取消') INT(11)"`
	RequestUserId   int    `json:"request_user_id" xorm:"default 0 comment('领用人Id') INT(11)"`
	RequestUserName string `json:"request_user_name" xorm:"default '' comment('领用人姓名') VARCHAR(50)"`
	SourceId        int    `json:"source_id" xorm:"default 0 comment('订单来源
            1：自建单据  2：集采平台  ') INT(11)"`
	SupplierAddress string    `json:"supplier_address" xorm:"default '' comment('供应商地址') VARCHAR(100)"`
	SupplierContact string    `json:"supplier_contact" xorm:"default '' comment('供应商联系人') VARCHAR(50)"`
	SupplierId      int       `json:"supplier_id" xorm:"default 0 comment('供应商id') INT(11)"`
	SupplierName    string    `json:"supplier_name" xorm:"default '' comment('供应商name') VARCHAR(100)"`
	SupplierPhone   string    `json:"supplier_phone" xorm:"default '' comment('供应商电话') VARCHAR(50)"`
	TotalAmount     string    `json:"total_amount" xorm:"default 0.00000000 comment('总金额') DECIMAL(18,8)"`
	UpdateTimestamp time.Time `json:"update_timestamp" xorm:"not null default 'CURRENT_TIMESTAMP(6)' comment('更新时间戳') TIMESTAMP(6)"`
}
