package goxorm

import (
	"time"
)

type THosDictCurrency struct {
	AddTime      time.Time `json:"add_time" xorm:"default 'CURRENT_TIMESTAMP' comment('添加时间') DATETIME"`
	CompanyId    int       `json:"company_id" xorm:"not null comment('公司ID') INT(11)"`
	CurrencyAbbr string    `json:"currency_abbr" xorm:"not null comment('货币缩写') VARCHAR(30)"`
	CurrencyName string    `json:"currency_name" xorm:"not null comment('货币名称') VARCHAR(30)"`
	Id           int       `json:"id" xorm:"not null pk autoincr comment('主键') INT(11)"`
	LastModify   time.Time `json:"last_modify" xorm:"default 'CURRENT_TIMESTAMP' comment('修改时间') DATETIME"`
}
