package goxorm

type CrmCardsCategoryUseRange struct {
	CardCategoryId    int `json:"card_category_id" xorm:"INT(11)"`
	CardCategoryValue int `json:"card_category_value" xorm:"comment('卡类别关联值') INT(11)"`
	Id                int `json:"id" xorm:"not null pk autoincr comment('Id') INT(11)"`
	ItemCategoryId    int `json:"item_category_id" xorm:"comment('可使用目录(专项卡时用)') INT(11)"`
	ItemId            int `json:"item_id" xorm:"comment('次卡时绑定的条目Id（次卡用）') INT(11)"`
	OrgId             int `json:"org_id" xorm:"comment('机构Id') index INT(11)"`
}
