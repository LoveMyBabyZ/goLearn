package goxorm

import (
	"time"
)

type CrmCustomerAccountDetail struct {
	ChainId              int       `json:"chain_id" xorm:"comment('连锁Id') INT(11)"`
	ConsumptionAmount    string    `json:"consumption_amount" xorm:"comment('消费金额') DECIMAL(18,2)"`
	ConsumptionFrequency int       `json:"consumption_frequency" xorm:"comment('消费次数') INT(11)"`
	CreateTime           time.Time `json:"create_time" xorm:"comment('创建时间') DATETIME"`
	CreateUserId         int       `json:"create_user_id" xorm:"comment('创建人Id') INT(11)"`
	CreateUserName       string    `json:"create_user_name" xorm:"comment('创建人名字') VARCHAR(50)"`
	CustomerId           int       `json:"customer_id" xorm:"comment('会员Id') INT(11)"`
	CustomerLevelId      int       `json:"customer_level_id" xorm:"comment('会员等级') INT(11)"`
	CustomerName         string    `json:"customer_name" xorm:"comment('会员名称') VARCHAR(50)"`
	Id                   int       `json:"id" xorm:"not null pk autoincr INT(11)"`
	LastUpdateTime       time.Time `json:"last_update_time" xorm:"comment('修改时间') DATETIME"`
	LastUpdateUserId     int       `json:"last_update_user_id" xorm:"comment('修改人Id') INT(11)"`
	LastUpdateUserName   string    `json:"last_update_user_name" xorm:"comment('修改人名字') VARCHAR(50)"`
	OrderId              int       `json:"order_id" xorm:"comment('订单Id') INT(11)"`
	OrderNumber          string    `json:"order_number" xorm:"comment('订单流水号') VARCHAR(50)"`
	OrgId                int       `json:"org_id" xorm:"comment('机构Id') INT(11)"`
	PaymentType          int       `json:"payment_type" xorm:"comment('支付方式') INT(11)"`
	Points               string    `json:"points" xorm:"comment('积分') DECIMAL(18)"`
	Proportion           string    `json:"proportion" xorm:"comment('积分倍率') DECIMAL(18,2)"`
	RmaId                int       `json:"rma_id" xorm:"INT(11)"`
	RmaNumber            string    `json:"rma_number" xorm:"VARCHAR(50)"`
	TransactionType      int       `json:"transaction_type" xorm:"comment('交易类型(1:付款.2:退款)') INT(11)"`
	UpdateTimestamp      time.Time `json:"update_timestamp" xorm:"not null default 'CURRENT_TIMESTAMP(6)' comment('更新时间戳') TIMESTAMP(6)"`
}
