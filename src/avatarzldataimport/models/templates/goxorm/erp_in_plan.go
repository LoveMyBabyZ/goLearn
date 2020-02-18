package goxorm

import (
	"time"
)

type ErpInPlan struct {
	AuditTime      time.Time `json:"audit_time" xorm:"default 'CURRENT_TIMESTAMP' comment('审核时间') index(org_id) DATETIME"`
	AuditorId      int       `json:"auditor_id" xorm:"default 0 comment('审核人Id') INT(11)"`
	AuditorName    string    `json:"auditor_name" xorm:"default '' comment('审核人姓名') VARCHAR(50)"`
	CreateTime     time.Time `json:"create_time" xorm:"default 'CURRENT_TIMESTAMP' comment('创建时间(制单时间)') index(org_id) DATETIME"`
	CreateUserId   int       `json:"create_user_id" xorm:"not null default 0 comment('创建人id(制单人)') INT(11)"`
	CreateUserName string    `json:"create_user_name" xorm:"not null default '' comment('创建人姓名(制单人)') VARCHAR(50)"`
	Id             int       `json:"id" xorm:"not null pk autoincr comment('id') INT(11)"`
	InBillNumber   string    `json:"in_bill_number" xorm:"comment('对应入库单据号') VARCHAR(500)"`
	InPlanNumber   string    `json:"in_plan_number" xorm:"not null default '' comment('入库计划单号') unique VARCHAR(50)"`
	InPlanStatus   int       `json:"in_plan_status" xorm:"not null default 0 comment('单据状态  
  0:保存 10：待审核 20:已审核 21：取消采购 30：已审核-入库中 40：已审核-已完成  -1：删除 ') INT(11)"`
	ItemCount          int    `json:"item_count" xorm:"not null default 0 comment('总项数') INT(11)"`
	OrderType          int    `json:"order_type" xorm:"not null default 0 comment('单据类型 1：采购') index(org_id) INT(11)"`
	OrgId              int    `json:"org_id" xorm:"not null default 0 comment('机构Id') index(org_id) INT(11)"`
	RelationBillNumber string `json:"relation_bill_number" xorm:"default '' comment('相关单据号') index(org_id) VARCHAR(50)"`
	SourceId           int    `json:"source_id" xorm:"not null default 0 comment('订单来源
            1：自建采购单  2：集采平台  ') index(org_id) INT(11)"`
	SupplierAddress string    `json:"supplier_address" xorm:"not null default '' comment('供应商地址') VARCHAR(100)"`
	SupplierContact string    `json:"supplier_contact" xorm:"not null default '' comment('供应商联系人') VARCHAR(50)"`
	SupplierId      int       `json:"supplier_id" xorm:"not null default 0 comment('供应商id') index(org_id) INT(11)"`
	SupplierName    string    `json:"supplier_name" xorm:"not null default '' comment('供应商name') index(org_id) VARCHAR(100)"`
	SupplierPhone   string    `json:"supplier_phone" xorm:"not null default '' comment('供应商电话') VARCHAR(50)"`
	TotalAmount     string    `json:"total_amount" xorm:"not null default 0.00000000 comment('总金额') DECIMAL(18,8)"`
	UpdateTimestamp time.Time `json:"update_timestamp" xorm:"not null default 'CURRENT_TIMESTAMP(6)' comment('更新时间戳') TIMESTAMP(6)"`
}
