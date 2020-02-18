package goxorm

import (
	"time"
)

type CrmCards struct {
	CardName           string    `json:"card_name" xorm:"comment('卡名称') VARCHAR(50)"`
	CardNumber         string    `json:"card_number" xorm:"comment('卡号') unique VARCHAR(100)"`
	CardType           int       `json:"card_type" xorm:"comment('卡类型(词典表中取数)') index INT(11)"`
	ChainId            int       `json:"chain_id" xorm:"default 0 comment('连锁Id') INT(11)"`
	Chained            int       `json:"chained" xorm:"comment('是否连锁 1：是  0：否') INT(255)"`
	Comment            string    `json:"comment" xorm:"comment('备注') VARCHAR(500)"`
	CreateTime         time.Time `json:"create_time" xorm:"comment('发卡时间（创建时间）') DATETIME"`
	CustomerId         int       `json:"customer_id" xorm:"comment('会员Id') index index(org_id) INT(11)"`
	ExpiryDate         time.Time `json:"expiry_date" xorm:"comment('有效期（过期时间）') DATE"`
	Id                 int       `json:"id" xorm:"not null pk autoincr comment('Id') INT(11)"`
	LastUpdateTime     time.Time `json:"last_update_time" xorm:"comment('最后更新时间') DATETIME"`
	LastUpdateUserId   int       `json:"last_update_user_id" xorm:"comment('最后更新人Id') INT(11)"`
	LastUpdateUserName string    `json:"last_update_user_name" xorm:"comment('最后更新人名字') VARCHAR(50)"`
	OrgId              int       `json:"org_id" xorm:"comment('机构Id') index(org_id) INT(11)"`
	OrgName            string    `json:"org_name" xorm:"comment('机构名称') VARCHAR(50)"`
	Password           string    `json:"password" xorm:"comment('卡密码') VARCHAR(50)"`
	SellerId           int       `json:"seller_id" xorm:"comment('销售人Id') INT(11)"`
	SellerName         string    `json:"seller_name" xorm:"comment('销售人姓名') VARCHAR(50)"`
	Status             int       `json:"status" xorm:"comment('卡片状态 （0：停用 1：启用 2：删除  3：到期  4：已退卡  5：退卡中）') INT(11)"`
	UpdateTimestamp    time.Time `json:"update_timestamp" xorm:"not null default 'CURRENT_TIMESTAMP(6)' comment('时间戳') TIMESTAMP(6)"`
	UseRange           int       `json:"use_range" xorm:"comment('使用范围 1：专项  2：通用') INT(11)"`
	UseType            int       `json:"use_type" xorm:"comment('使用方式 0:计次 1：计额') index INT(11)"`
}
