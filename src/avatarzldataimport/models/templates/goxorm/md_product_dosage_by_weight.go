package goxorm

import (
	"time"
)

type MdProductDosageByWeight struct {
	AddTime      time.Time `json:"add_time" xorm:"default 'CURRENT_TIMESTAMP' comment('添加时间') DATETIME"`
	Id           int       `json:"id" xorm:"not null pk autoincr INT(11)"`
	LastModify   time.Time `json:"last_modify" xorm:"default 'CURRENT_TIMESTAMP' comment('修改时间') DATETIME"`
	MaxWeight    string    `json:"max_weight" xorm:"not null default 0.00000 comment('最大体重') DECIMAL(19,5)"`
	MinWeight    string    `json:"min_weight" xorm:"not null default 0.00000 comment('最小体重') DECIMAL(19,5)"`
	OrgId        int       `json:"org_id" xorm:"not null comment('医院ID') index(idx_org_pid) INT(11)"`
	ProductId    int       `json:"product_id" xorm:"not null default 0 comment('产品ID') index(idx_org_pid) INT(11)"`
	ProductType  int       `json:"product_type" xorm:"not null comment('产品类型') INT(11)"`
	SingleDosage string    `json:"single_dosage" xorm:"not null default 0.00000 comment('单次用量') DECIMAL(19,5)"`
}
