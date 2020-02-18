package goxorm

import (
	"time"
)

type PosPaymentMethod struct {
	Amount                  string    `json:"amount" xorm:"comment('总金额') DECIMAL(18,2)"`
	CreateDate              time.Time `json:"create_date" xorm:"comment('创建时间') DATETIME"`
	CreateUserid            int       `json:"create_userid" xorm:"comment('创建人Id') INT(11)"`
	CreateUsername          string    `json:"create_username" xorm:"comment('创建人姓名') VARCHAR(50)"`
	CustomerId              int       `json:"customer_id" xorm:"comment('会员Id') index(org_id) INT(11)"`
	CustomerName            string    `json:"customer_name" xorm:"comment('会员名称') VARCHAR(50)"`
	Frequency               int       `json:"frequency" xorm:"comment('次数') INT(11)"`
	FrequencyDiscountAmount string    `json:"frequency_discount_amount" xorm:"default 0.00 comment('次卡折扣金额') DECIMAL(18,2)"`
	Id                      int       `json:"id" xorm:"not null pk autoincr comment('主键') INT(11)"`
	LastUpdateDate          time.Time `json:"last_update_date" xorm:"comment('更新时间') DATETIME"`
	MemberCardId            int       `json:"member_card_id" xorm:"comment('卡Id') INT(11)"`
	MemberCardNumber        string    `json:"member_card_number" xorm:"comment('卡号') VARCHAR(50)"`
	OperatorId              int       `json:"operator_id" xorm:"comment('操作人Id') INT(11)"`
	OperatorName            string    `json:"operator_name" xorm:"comment('操作人姓名') VARCHAR(50)"`
	OrderId                 int       `json:"order_id" xorm:"not null default 0 comment('订单Id') index index(org_id) INT(11)"`
	OrderNumber             string    `json:"order_number" xorm:"comment('支付流水号') index(org_id) VARCHAR(50)"`
	OrgId                   int       `json:"org_id" xorm:"not null default 0 comment('机构Id') index(org_id) INT(11)"`
	PaymentType             int       `json:"payment_type" xorm:"comment('支付类型') INT(11)"`
	ProductId               int       `json:"product_id" xorm:"not null default 0 comment('产品Id') INT(11)"`
	Remark                  string    `json:"remark" xorm:"comment('备注') VARCHAR(200)"`
	UpdateTimestamp         time.Time `json:"update_timestamp" xorm:"not null default 'CURRENT_TIMESTAMP(6)' comment('更新时间戳') TIMESTAMP(6)"`
}
