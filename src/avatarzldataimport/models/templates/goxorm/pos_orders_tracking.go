package goxorm

import (
	"time"
)

type PosOrdersTracking struct {
	CustomerId      int       `json:"customer_id" xorm:"comment('会员Id') index(org_id) INT(20)"`
	CustomerName    string    `json:"customer_name" xorm:"comment('会员名称') VARCHAR(50)"`
	Id              int       `json:"id" xorm:"not null pk autoincr comment('主键') INT(11)"`
	OperatorDate    time.Time `json:"operator_date" xorm:"comment('操作时间') index(org_id) DATETIME"`
	OperatorId      int       `json:"operator_id" xorm:"comment('操作人Id') index(org_id) INT(11)"`
	OperatorName    string    `json:"operator_name" xorm:"comment('操作人姓名') VARCHAR(50)"`
	OrderId         int       `json:"order_id" xorm:"comment('支付Id') INT(11)"`
	OrderNumber     string    `json:"order_number" xorm:"comment('支付流水号') VARCHAR(50)"`
	OrderStatus     int       `json:"order_status" xorm:"comment('支付状态') INT(11)"`
	OrgId           int       `json:"org_id" xorm:"comment('机构Id') index(org_id) INT(11)"`
	PetId           int       `json:"pet_id" xorm:"comment('宠物Id') INT(11)"`
	PetName         string    `json:"pet_name" xorm:"comment('宠物名称') VARCHAR(50)"`
	Remark          string    `json:"remark" xorm:"comment('备注') VARCHAR(200)"`
	UpdateTimestamp time.Time `json:"update_timestamp" xorm:"not null default 'CURRENT_TIMESTAMP(6)' comment('更新时间戳') TIMESTAMP(6)"`
}
