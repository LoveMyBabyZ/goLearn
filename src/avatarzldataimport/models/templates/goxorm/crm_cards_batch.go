package goxorm

import (
	"time"
)

type CrmCardsBatch struct {
	BatchNumber        string    `json:"batch_number" xorm:"comment('批次号') VARCHAR(50)"`
	CardId             int       `json:"card_id" xorm:"comment('卡Id') index INT(11)"`
	ChargeTimes        int       `json:"charge_times" xorm:"comment('充值次数') INT(11)"`
	Comment            string    `json:"comment" xorm:"comment('备注') VARCHAR(500)"`
	CorpusBalance      string    `json:"corpus_balance" xorm:"comment('本金余额') DECIMAL(18,2)"`
	CostMoney          string    `json:"cost_money" xorm:"comment('实际花费金额') DECIMAL(18,2)"`
	CreateTime         time.Time `json:"create_time" xorm:"comment('充值时间（批次创建时间）') DATETIME"`
	CustomerId         int       `json:"customer_id" xorm:"comment('会员Id') index(org_id) INT(11)"`
	Id                 int       `json:"id" xorm:"not null pk autoincr comment('Id') INT(11)"`
	LastUpdateTime     time.Time `json:"last_update_time" xorm:"comment('最后更新时间') DATETIME"`
	LastUpdateUserId   int       `json:"last_update_user_id" xorm:"comment('最后更新人Id') INT(11)"`
	LastUpdateUserName string    `json:"last_update_user_name" xorm:"comment('最后更新人名字') VARCHAR(50)"`
	LeftTimes          int       `json:"left_times" xorm:"comment('本金次数剩余次数') INT(11)"`
	OrgId              int       `json:"org_id" xorm:"comment('机构Id') index(org_id) INT(11)"`
	OriginalPrice      string    `json:"original_price" xorm:"comment('次卡原单价') DECIMAL(18,2)"`
	PresentBalance     string    `json:"present_balance" xorm:"comment('赠送账户余额（赠送帐户，计额时有效）') DECIMAL(18,2)"`
	PresentLeftTimes   int       `json:"present_left_times" xorm:"comment('赠送次数剩余次数') INT(11)"`
	PresentMoney       string    `json:"present_money" xorm:"comment('实际赠送金额') DECIMAL(18,2)"`
	PresentTimes       int       `json:"present_times" xorm:"comment('赠送次数') INT(11)"`
	PricePerTime       string    `json:"price_per_time" xorm:"comment('次卡实际单价(按次计算时有效，计额时为0)') DECIMAL(18,2)"`
	RunOut             int       `json:"run_out" xorm:"default 0 comment('是否已用完  1: 已经用完  0：未用完') INT(11)"`
	SellerId           int       `json:"seller_id" xorm:"comment('销售人Id') INT(11)"`
	SellerName         string    `json:"seller_name" xorm:"comment('销售人姓名') VARCHAR(50)"`
	UpdateTimestamp    time.Time `json:"update_timestamp" xorm:"not null default 'CURRENT_TIMESTAMP(6)' comment('时间戳') TIMESTAMP(6)"`
}
