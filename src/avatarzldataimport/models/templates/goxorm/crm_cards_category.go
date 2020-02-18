package goxorm

import (
	"time"
)

type CrmCardsCategory struct {
	CategoryCode       string    `json:"category_code" xorm:"comment('类目') VARCHAR(50)"`
	CategoryName       string    `json:"category_name" xorm:"comment('卡名称') VARCHAR(20)"`
	CategoryValue      int       `json:"category_value" xorm:"comment('卡类别值') INT(11)"`
	ChainId            int       `json:"chain_id" xorm:"comment('连锁ID') INT(11)"`
	Chained            int       `json:"chained" xorm:"comment('是否连锁 1：是  0：否') INT(255)"`
	CorpusMoney        string    `json:"corpus_money" xorm:"comment('本金金额') DECIMAL(18,2)"`
	CorpusTime         int       `json:"corpus_time" xorm:"comment('本金次数') INT(11)"`
	CreateTime         time.Time `json:"create_time" xorm:"comment('创建时间') DATETIME"`
	CreateUserId       int       `json:"create_user_id" xorm:"comment('创建人Id') INT(11)"`
	CreateUserName     string    `json:"create_user_name" xorm:"comment('创建人姓名') VARCHAR(50)"`
	Deleted            int       `json:"deleted" xorm:"comment('0:未删除  1：已删除') INT(11)"`
	Enabled            int       `json:"enabled" xorm:"comment('0： 停用  1：启用') INT(11)"`
	Id                 int       `json:"id" xorm:"not null pk autoincr comment('Id') INT(11)"`
	Illustrate         string    `json:"illustrate" xorm:"comment('说明') VARCHAR(500)"`
	ItemCode           string    `json:"item_code" xorm:"comment('产品总部编码') VARCHAR(50)"`
	ItemId             int       `json:"item_id" xorm:"comment('产品ID') INT(11)"`
	ItemName           string    `json:"item_name" xorm:"comment('产品名称') VARCHAR(50)"`
	ItemPrice          string    `json:"item_price" xorm:"DECIMAL(18,2)"`
	LastUpdateTime     time.Time `json:"last_update_time" xorm:"comment('最后更新时间') DATETIME"`
	LastUpdateUserId   int       `json:"last_update_user_id" xorm:"comment('最后更新人Id') INT(11)"`
	LastUpdateUserName string    `json:"last_update_user_name" xorm:"comment('最后更新人名字') VARCHAR(50)"`
	OrgId              int       `json:"org_id" xorm:"comment('机构Id') index INT(11)"`
	PresentMoney       string    `json:"present_money" xorm:"comment('赠送金额') DECIMAL(18,2)"`
	PresentTime        int       `json:"present_time" xorm:"comment('赠送次数') INT(11)"`
	SortValue          int       `json:"sort_value" xorm:"comment('排序值 ') INT(11)"`
	UsageMode          int       `json:"usage_mode" xorm:"comment('使用方式  0：计次卡  1：计额卡（储值卡）') INT(11)"`
	UsageRange         int       `json:"usage_range" xorm:"comment('1:专项 2：通用') INT(11)"`
}
