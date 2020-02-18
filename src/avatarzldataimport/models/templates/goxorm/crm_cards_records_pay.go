package goxorm

type CrmCardsRecordsPay struct {
	CardId     int    `json:"card_id" xorm:"comment('会员卡Id') INT(11)"`
	CustomerId int    `json:"customer_id" xorm:"comment('会员Id') index(org_id) INT(11)"`
	Id         int    `json:"id" xorm:"not null pk autoincr comment('Id') INT(11)"`
	OrgId      int    `json:"org_id" xorm:"comment('机构Id') index(org_id) INT(11)"`
	PayMoney   string `json:"pay_money" xorm:"comment('支付金额') DECIMAL(18,2)"`
	PayType    int    `json:"pay_type" xorm:"comment('支付方式Id') INT(11)"`
	RecordId   int    `json:"record_id" xorm:"comment('操作记录Id') index INT(11)"`
}
