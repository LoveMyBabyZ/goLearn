package goxorm

import (
	"time"
)

type THosDictInvoiceType struct {
	AddTime     time.Time `json:"add_time" xorm:"default 'CURRENT_TIMESTAMP' comment('添加时间') DATETIME"`
	CompanyId   int       `json:"company_id" xorm:"not null comment('公司ID') INT(11)"`
	Id          int       `json:"id" xorm:"not null pk autoincr comment('主键') INT(11)"`
	InvoiceType string    `json:"invoice_type" xorm:"not null comment('发票性质名称') VARCHAR(30)"`
	LastModify  time.Time `json:"last_modify" xorm:"default 'CURRENT_TIMESTAMP' comment('修改时间') DATETIME"`
}
