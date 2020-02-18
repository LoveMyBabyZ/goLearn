package goxorm

type CrmCustomerLevel struct {
	ChainId        int    `json:"chain_id" xorm:"comment('连锁Id') INT(11)"`
	Discount       string `json:"discount" xorm:"not null comment('整体折扣') DECIMAL(5,4)"`
	DiscountStatus int    `json:"discount_status" xorm:"not null comment('折扣状态 1 统一折扣 2 分品类折扣') INT(11)"`
	DiscountSwitch int    `json:"discount_switch" xorm:"not null comment('折扣开关 1 开 2 关') INT(11)"`
	Id             int    `json:"id" xorm:"not null pk autoincr INT(11)"`
	Level          int    `json:"level" xorm:"comment('会员等级') INT(11)"`
	LevelCode      string `json:"level_code" xorm:"comment('会员等级编码') VARCHAR(20)"`
	Name           string `json:"name" xorm:"comment('会员等级名称') VARCHAR(20)"`
	Proportion     string `json:"proportion" xorm:"comment('积分倍率') DECIMAL(18,2)"`
}
