package goxorm

import (
	"time"
)

type ErpInBill struct {
	AuditTime      time.Time `json:"audit_time" xorm:"default 'CURRENT_TIMESTAMP' comment('审核时间') index(org_id) DATETIME"`
	AuditorId      int       `json:"auditor_id" xorm:"default 0 comment('审核人Id') INT(11)"`
	AuditorName    string    `json:"auditor_name" xorm:"default '' comment('审核人姓名') VARCHAR(50)"`
	CreateTime     time.Time `json:"create_time" xorm:"default 'CURRENT_TIMESTAMP' comment('创建时间(制单时间)') index(org_id) DATETIME"`
	CreateUserId   int       `json:"create_user_id" xorm:"default 0 comment('创建人id(制单人)') INT(11)"`
	CreateUserName string    `json:"create_user_name" xorm:"default '' comment('创建人姓名(制单人)') VARCHAR(50)"`
	DiscountAmount string    `json:"discount_amount" xorm:"default 0.00000000 comment('优惠金额') DECIMAL(18,8)"`
	Id             int       `json:"id" xorm:"not null pk autoincr comment('Id') INT(11)"`
	InBillNumber   string    `json:"in_bill_number" xorm:"default '' comment('入库单号') unique VARCHAR(50)"`
	InBillStatus   int       `json:"in_bill_status" xorm:"default 0 comment('入库单状态 
            0：保存 1：待审核  2：已审核  3：已审核-有退货') INT(11)"`
	InBillType int `json:"in_bill_type" xorm:"default 0 comment('单据类型
            1. 采购入库单  2.结余入库  3.销售盘盈入库  4.盘点盘盈入库 5.销售退货入库') index(org_id) INT(11)"`
	InPlanId        int    `json:"in_plan_id" xorm:"default 0 comment('入库计划单Id') INT(11)"`
	InRemark        string `json:"in_remark" xorm:"comment('备注') VARCHAR(200)"`
	ItemCount       int    `json:"item_count" xorm:"default 0 comment('总项数') INT(11)"`
	NcOrderNumber   string `json:"nc_order_number" xorm:"default '' comment('nc单号') VARCHAR(40)"`
	OrgId           int    `json:"org_id" xorm:"default 0 comment('机构id') index(org_id) INT(11)"`
	OutstockOrgId   int    `json:"outstock_org_id" xorm:"default 0 comment('出库机构ID') INT(11)"`
	OutstockOrgName string `json:"outstock_org_name" xorm:"comment('出库机构名称') VARCHAR(255)"`
	PayStatus       int    `json:"pay_status" xorm:"default 0 comment('付款状态
            0：未付款  1：已付款') INT(11)"`
	ReceivedDate       time.Time `json:"received_date" xorm:"default 'CURRENT_TIMESTAMP' comment('到货日期') DATETIME"`
	RelationBillId     int       `json:"relation_bill_id" xorm:"default 0 comment('相关单据ID') INT(11)"`
	RelationBillNumber string    `json:"relation_bill_number" xorm:"comment('相关单据号') VARCHAR(255)"`
	SourceId           int       `json:"source_id" xorm:"default 0 comment('入库单来源
            1：自建采购单  2：集采平台  ') INT(11)"`
	SupplierAddress string    `json:"supplier_address" xorm:"default '' comment('供应商地址') VARCHAR(100)"`
	SupplierContact string    `json:"supplier_contact" xorm:"default '' comment('供应商联系人') VARCHAR(50)"`
	SupplierId      int       `json:"supplier_id" xorm:"default 0 comment('供应商Id') index(org_id) INT(11)"`
	SupplierName    string    `json:"supplier_name" xorm:"default '' comment('供应商名称') index(org_id) VARCHAR(100)"`
	SupplierPhone   string    `json:"supplier_phone" xorm:"default '' comment('供应商电话') VARCHAR(50)"`
	TotalAmount     string    `json:"total_amount" xorm:"default 0.00000000 comment('总金额') DECIMAL(18,8)"`
	UpdateTimestamp time.Time `json:"update_timestamp" xorm:"not null default 'CURRENT_TIMESTAMP(6)' comment('更新时间戳') TIMESTAMP(6)"`
}
