package goxorm

type CrmCustomerLevelDiscounts struct {
	Discount         string `json:"discount" xorm:"not null comment('打折倍率') DECIMAL(5,4)"`
	DiscountCategory int    `json:"discount_category" xorm:"not null comment('打折品类') INT(11)"`
	DiscountType     int    `json:"discount_type" xorm:"not null comment('打折类型') INT(11)"`
	DiscountTypeCn   string `json:"discount_type_cn" xorm:"not null comment('打折类型中文') VARCHAR(20)"`
	Id               int    `json:"id" xorm:"not null pk autoincr INT(11)"`
	LevelCode        string `json:"level_code" xorm:"comment('会员等级编码') VARCHAR(20)"`
	LevelId          int    `json:"level_id" xorm:"not null comment('会员等级Id') INT(11)"`
}
