package goxorm

import (
	"time"
)

type PosRmaRefund struct {
	CreateDate       time.Time `json:"create_date" xorm:"comment('创建时间') DATETIME"`
	CreateUserid     int       `json:"create_userid" xorm:"comment('创建人Id') INT(11)"`
	CreateUsername   string    `json:"create_username" xorm:"comment('创建人姓名') VARCHAR(50)"`
	Id               int       `json:"id" xorm:"not null pk autoincr comment('主键') INT(11)"`
	LastUpdateDate   time.Time `json:"last_update_date" xorm:"comment('更新时间') DATETIME"`
	MemberCardId     int       `json:"member_card_id" xorm:"comment('卡Id') INT(11)"`
	MemberCardNumber string    `json:"member_card_number" xorm:"comment('卡号') VARCHAR(50)"`
	OperatorId       int       `json:"operator_id" xorm:"comment('操作人Id') INT(11)"`
	OperatorName     string    `json:"operator_name" xorm:"comment('操作人姓名') VARCHAR(50)"`
	OrderId          int       `json:"order_id" xorm:"not null default 0 comment('订单Id') INT(11)"`
	OrderNumber      string    `json:"order_number" xorm:"comment('支付流水号') index(org_id) VARCHAR(50)"`
	OrgId            int       `json:"org_id" xorm:"not null default 0 comment('机构Id') index(org_id) INT(11)"`
	ProductId        int       `json:"product_id" xorm:"not null default 0 comment('产品Id') INT(11)"`
	RefundAmount     string    `json:"refund_amount" xorm:"comment('退款金额') DECIMAL(18,2)"`
	RefundFrequency  int       `json:"refund_frequency" xorm:"comment('退货次数') INT(11)"`
	RefundType       int       `json:"refund_type" xorm:"comment('退款类型') INT(11)"`
	Remark           string    `json:"remark" xorm:"comment('备注') VARCHAR(200)"`
	RmaId            int       `json:"rma_id" xorm:"not null default 0 comment('退货Id') index(org_id) INT(11)"`
	RmaNumber        string    `json:"rma_number" xorm:"comment('支付流水号') index(org_id) VARCHAR(50)"`
	UpdateTimestamp  time.Time `json:"update_timestamp" xorm:"not null default 'CURRENT_TIMESTAMP(6)' comment('更新时间戳') TIMESTAMP(6)"`
}
