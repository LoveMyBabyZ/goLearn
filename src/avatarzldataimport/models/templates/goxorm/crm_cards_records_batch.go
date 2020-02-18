package goxorm

import (
	"time"
)

type CrmCardsRecordsBatch struct {
	BatchNumber     string    `json:"batch_number" xorm:"comment('会员卡批次号') VARCHAR(50)"`
	CardId          int       `json:"card_id" xorm:"comment('会员卡Id') INT(11)"`
	CorpusPayMoney  string    `json:"corpus_pay_money" xorm:"comment('本金支付金额') DECIMAL(18,2)"`
	CreateTime      time.Time `json:"create_time" xorm:"comment('记录创建时间') DATETIME"`
	CustomerId      int       `json:"customer_id" xorm:"comment('会员Id') index(org_id) INT(11)"`
	Id              int       `json:"id" xorm:"not null pk autoincr comment('Id') INT(11)"`
	OrgId           int       `json:"org_id" xorm:"comment('机构Id') index(org_id) INT(11)"`
	PayNumber       string    `json:"pay_number" xorm:"comment('销售单支付流水号') VARCHAR(50)"`
	PayTimes        int       `json:"pay_times" xorm:"comment('本金次数支付次数') INT(11)"`
	PresentPayMoney string    `json:"present_pay_money" xorm:"comment('赠送账户支付金额') DECIMAL(18,2)"`
	PresentPayTimes int       `json:"present_pay_times" xorm:"comment('赠送次数支付次数') INT(11)"`
	RecordId        int       `json:"record_id" xorm:"comment('流水账记录Id') INT(11)"`
}
