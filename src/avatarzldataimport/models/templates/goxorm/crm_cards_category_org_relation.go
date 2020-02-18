package goxorm

type CrmCardsCategoryOrgRelation struct {
	CardCategoryId int `json:"card_category_id" xorm:"comment('卡类别Id') index(org_id) INT(11)"`
	Enabled        int `json:"enabled" xorm:"comment('0： 停用  1：启用') INT(11)"`
	Id             int `json:"id" xorm:"not null pk autoincr comment('Id') INT(11)"`
	OrgId          int `json:"org_id" xorm:"comment('机构Id') index(org_id) INT(11)"`
}
