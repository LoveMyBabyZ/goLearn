package goxorm

import (
	"time"
)

type PosRmaMaster struct {
	AuditorId          int       `json:"auditor_id" xorm:"default 0 comment('审核人ID') INT(11)"`
	AuditorName        string    `json:"auditor_name" xorm:"comment('审核人') VARCHAR(50)"`
	AuditorTime        time.Time `json:"auditor_time" xorm:"comment('审核时间') DATETIME"`
	CreateUserid       int       `json:"create_userid" xorm:"comment('创建人Id') INT(11)"`
	CreateUsername     string    `json:"create_username" xorm:"comment('创建人姓名') VARCHAR(50)"`
	CreatedDate        time.Time `json:"created_date" xorm:"comment('创建时间') DATETIME"`
	CustomerId         int       `json:"customer_id" xorm:"comment('会员Id') index(org_id) INT(11)"`
	CustomerName       string    `json:"customer_name" xorm:"comment('会员名称') VARCHAR(50)"`
	CustomerNumber     string    `json:"customer_number" xorm:"comment('会员编码') VARCHAR(50)"`
	Id                 int       `json:"id" xorm:"not null pk autoincr comment('主键') INT(11)"`
	InstoreStatus      int       `json:"instore_status" xorm:"comment('入库状态') INT(11)"`
	LastPrintTime      time.Time `json:"last_print_time" xorm:"comment('最后打印时间') DATETIME"`
	LastUpdateDate     time.Time `json:"last_update_date" xorm:"comment('更新时间') DATETIME"`
	LastUpdateUserid   int       `json:"last_update_userid" xorm:"comment('更新人Id') INT(11)"`
	LastUpdateUsername string    `json:"last_update_username" xorm:"comment('更新人姓名') VARCHAR(50)"`
	OrderId            int       `json:"order_id" xorm:"not null default 0 comment('订单Id') INT(11)"`
	OrderNumber        string    `json:"order_number" xorm:"comment('支付流水号') index(org_id) VARCHAR(50)"`
	OrgId              int       `json:"org_id" xorm:"not null default 0 comment('机构Id') index(org_id) INT(11)"`
	PetId              int       `json:"pet_id" xorm:"comment('宠物Id') INT(11)"`
	PetName            string    `json:"pet_name" xorm:"comment('宠物名称') VARCHAR(50)"`
	PrintSerialNumber  string    `json:"print_serial_number" xorm:"comment('小票打印编号') VARCHAR(50)"`
	RmaAmount          string    `json:"rma_amount" xorm:"comment('总金额') DECIMAL(18,2)"`
	RmaComment         string    `json:"rma_comment" xorm:"comment('退款备注') VARCHAR(300)"`
	RmaDate            time.Time `json:"rma_date" xorm:"comment('退货时间') DATETIME"`
	RmaNumber          string    `json:"rma_number" xorm:"comment('退货流水号') index(org_id) VARCHAR(50)"`
	RmaReason          string    `json:"rma_reason" xorm:"comment('退款原因') VARCHAR(100)"`
	RmaReasonId        int       `json:"rma_reason_id" xorm:"comment('退款原因Id') INT(11)"`
	RmaStatus          int       `json:"rma_status" xorm:"comment('状态(1:正常退货)') INT(11)"`
	RmaTotalAmount     string    `json:"rma_total_amount" xorm:"default 0.00 comment('应退金额') DECIMAL(18,2)"`
	RmaType            int       `json:"rma_type" xorm:"default 0 comment('RMA类型 1:普通退货 2:记帐退货') INT(11)"`
	UpdateTimestamp    time.Time `json:"update_timestamp" xorm:"not null default 'CURRENT_TIMESTAMP(6)' comment('更新时间戳') TIMESTAMP(6)"`
}
