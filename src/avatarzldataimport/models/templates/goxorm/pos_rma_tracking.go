package goxorm

import (
	"time"
)

type PosRmaTracking struct {
	CustomerId      int       `json:"customer_id" xorm:"comment('会员Id') index(org_id) INT(20)"`
	CustomerName    string    `json:"customer_name" xorm:"comment('会员名称') VARCHAR(50)"`
	Id              int       `json:"id" xorm:"not null pk autoincr comment('主键') INT(11)"`
	OperatorDate    time.Time `json:"operator_date" xorm:"comment('操作时间') DATETIME"`
	OperatorId      int       `json:"operator_id" xorm:"comment('操作人Id') INT(11)"`
	OperatorName    string    `json:"operator_name" xorm:"comment('操作人姓名') VARCHAR(50)"`
	OrgId           int       `json:"org_id" xorm:"comment('机构Id') index(org_id) INT(11)"`
	PetId           int       `json:"pet_id" xorm:"comment('宠物Id') INT(11)"`
	PetName         string    `json:"pet_name" xorm:"comment('宠物名称') VARCHAR(50)"`
	Remark          string    `json:"remark" xorm:"comment('备注') VARCHAR(200)"`
	RmaId           int       `json:"rma_id" xorm:"comment('退货Id') index(org_id) INT(11)"`
	RmaNumber       string    `json:"rma_number" xorm:"comment('退货流水号') index(org_id) VARCHAR(50)"`
	RmaStatus       int       `json:"rma_status" xorm:"comment('退货状态') INT(11)"`
	UpdateTimestamp time.Time `json:"update_timestamp" xorm:"not null default 'CURRENT_TIMESTAMP(6)' comment('更新时间戳') TIMESTAMP(6)"`
}
