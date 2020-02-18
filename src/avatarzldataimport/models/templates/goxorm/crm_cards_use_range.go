package goxorm

type CrmCardsUseRange struct {
	CardId         int `json:"card_id" xorm:"comment('卡Id') INT(11)"`
	CardUseType    int `json:"card_use_type" xorm:"comment('卡使用方式Id') INT(11)"`
	CustomerId     int `json:"customer_id" xorm:"comment('会员Id') index(org_id) INT(11)"`
	Id             int `json:"id" xorm:"not null pk autoincr comment('Id') INT(11)"`
	ItemCategoryId int `json:"item_category_id" xorm:"comment('可使用目录') INT(11)"`
	ItemId         int `json:"item_id" xorm:"comment('次卡时绑定的条目Id') INT(11)"`
	OrgId          int `json:"org_id" xorm:"comment('机构Id') index(org_id) INT(11)"`
}
