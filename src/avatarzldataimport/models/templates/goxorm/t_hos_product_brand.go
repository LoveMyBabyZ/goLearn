package goxorm

import (
	"time"
)

type THosProductBrand struct {
	AddTime     time.Time `json:"add_time" xorm:"default 'CURRENT_TIMESTAMP' comment('添加时间') DATETIME"`
	BrandCode   string    `json:"brand_code" xorm:"not null default '' comment('品牌编码') VARCHAR(30)"`
	BrandName   string    `json:"brand_name" xorm:"not null comment('品牌名称') VARCHAR(30)"`
	CompanyId   int       `json:"company_id" xorm:"not null default 4 comment('公司ID') INT(11)"`
	Id          int       `json:"id" xorm:"not null pk autoincr comment('主键') INT(11)"`
	LastModify  time.Time `json:"last_modify" xorm:"default 'CURRENT_TIMESTAMP' comment('修改时间') DATETIME"`
	ProductType int       `json:"product_type" xorm:"not null default 0 comment('品类ID') INT(11)"`
}
