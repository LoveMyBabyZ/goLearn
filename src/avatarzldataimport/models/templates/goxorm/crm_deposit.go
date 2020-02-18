package goxorm

import (
	"time"
)

type CrmDeposit struct {
	Balance            string    `json:"balance" xorm:"comment('本金余额') DECIMAL(18,2)"`
	Comment            string    `json:"comment" xorm:"comment('备注') VARCHAR(500)"`
	CreateTime         time.Time `json:"create_time" xorm:"comment('创建时间') DATETIME"`
	CustomerId         int       `json:"customer_id" xorm:"comment('会员Id') index(org_id) INT(11)"`
	DepositType        int       `json:"deposit_type" xorm:"comment('押金类型') index(org_id) INT(11)"`
	FreezeAmount       string    `json:"freeze_amount" xorm:"default 0.00 comment('冻结金额') DECIMAL(18,2)"`
	Id                 int       `json:"id" xorm:"not null pk autoincr comment('Id') INT(11)"`
	IsRefunding        int       `json:"is_refunding" xorm:"default 0 comment('是否退款中(0-未退款 1-退款中)') INT(11)"`
	LastUpdateTime     time.Time `json:"last_update_time" xorm:"comment('最后更新时间') DATETIME"`
	LastUpdateUserId   int       `json:"last_update_user_id" xorm:"comment('最后更新人Id') INT(11)"`
	LastUpdateUserName string    `json:"last_update_user_name" xorm:"comment('最后更新人姓名') VARCHAR(50)"`
	OldId              int       `json:"old_id" xorm:"comment('旧编码') INT(11)"`
	OrgId              int       `json:"org_id" xorm:"comment('机构Id') index(org_id) INT(11)"`
	UpdateTimestamp    time.Time `json:"update_timestamp" xorm:"not null default 'CURRENT_TIMESTAMP(6)' comment('时间戳') TIMESTAMP(6)"`
}
