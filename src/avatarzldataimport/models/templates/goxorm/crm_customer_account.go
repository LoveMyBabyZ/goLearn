package goxorm

import (
	"time"
)

type CrmCustomerAccount struct {
	ChainId         int       `json:"chain_id" xorm:"comment('连锁Id') index(org_id) INT(11)"`
	CustomerId      int       `json:"customer_id" xorm:"comment('会员Id') index(org_id) INT(11)"`
	CustomerLevelId int       `json:"customer_level_id" xorm:"comment('会员等级ID') INT(11)"`
	CustomerName    string    `json:"customer_name" xorm:"comment('会员名称') VARCHAR(50)"`
	Id              int       `json:"id" xorm:"not null pk autoincr comment('Id') INT(11)"`
	Points          string    `json:"points" xorm:"comment('积分') DECIMAL(18)"`
	UpdateTimestamp time.Time `json:"update_timestamp" xorm:"not null default 'CURRENT_TIMESTAMP(6)' comment('时间戳') TIMESTAMP(6)"`
}
