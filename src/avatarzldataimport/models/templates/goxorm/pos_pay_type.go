package goxorm

import (
	"time"
)

type PosPayType struct {
	AllowPay           int       `json:"allow_pay" xorm:"comment('1：收银台启用 0：收银台停用') TINYINT(4)"`
	AllowRecharge      int       `json:"allow_recharge" xorm:"comment('1：充值启用  0：充值停用') TINYINT(4)"`
	CreateTime         time.Time `json:"create_time" xorm:"comment('创建时间') DATETIME"`
	CreateUserId       int       `json:"create_user_id" xorm:"comment('创建人Id') INT(11)"`
	CreateUserName     string    `json:"create_user_name" xorm:"comment('创建人姓名') VARCHAR(50)"`
	Enabled            int       `json:"enabled" xorm:"comment('0： 停用  1：启用') TINYINT(4)"`
	Id                 int       `json:"id" xorm:"not null pk autoincr comment('id') INT(11)"`
	IsDefault          int       `json:"is_default" xorm:"comment('1：默认支付方式 0：非默认') TINYINT(4)"`
	LastUpdateTime     time.Time `json:"last_update_time" xorm:"comment('最后更新时间') DATETIME"`
	LastUpdateUserId   int       `json:"last_update_user_id" xorm:"comment('最后更新人Id') INT(11)"`
	LastUpdateUserName string    `json:"last_update_user_name" xorm:"comment('最后更新人名字') VARCHAR(50)"`
	OrgId              int       `json:"org_id" xorm:"comment('机构Id') index INT(11)"`
	PayTypeName        string    `json:"pay_type_name" xorm:"comment('支付方式名称') VARCHAR(50)"`
	PayTypeValue       int       `json:"pay_type_value" xorm:"comment('支付方式值') INT(11)"`
	ShowRemark         int       `json:"show_remark" xorm:"comment('1：收款页面显示备注  0：不显示') TINYINT(4)"`
	SortValue          int       `json:"sort_value" xorm:"comment('排序值 ') INT(11)"`
}
