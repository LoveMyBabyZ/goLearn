package goxorm

import (
	"time"
)

type THosColumnConfig struct {
	AddTime        time.Time `json:"add_time" xorm:"default 'CURRENT_TIMESTAMP' comment('添加时间') DATETIME"`
	CompanyId      int       `json:"company_id" xorm:"not null default 4 comment('公司ID') INT(11)"`
	FieldType      int       `json:"field_type" xorm:"not null default 1 comment('字段类型（1：主产品字段；2：关联产品字段；3：关联化验字段）') TINYINT(1)"`
	Id             int       `json:"id" xorm:"not null pk autoincr comment('主键') INT(11)"`
	IsHeadControl  int       `json:"is_head_control" xorm:"not null default 0 comment('是否总部管控 0 不管控 1 管控') TINYINT(1)"`
	IsRequired     int       `json:"is_required" xorm:"not null default 0 comment('是否必填 1 必填 0 非必填') TINYINT(1)"`
	LastModify     time.Time `json:"last_modify" xorm:"default 'CURRENT_TIMESTAMP' comment('修改时间') DATETIME"`
	OperatorId     string    `json:"operator_id" xorm:"not null comment('操作人ID') VARCHAR(20)"`
	OperatorName   string    `json:"operator_name" xorm:"not null comment('操作人名字') VARCHAR(20)"`
	ProductKey     string    `json:"product_key" xorm:"not null default '' comment('字段key') index(INDEX_PROTYPE_PROKEY) VARCHAR(100)"`
	ProductKeyName string    `json:"product_key_name" xorm:"not null comment('字段名称') VARCHAR(50)"`
	ProductType    int       `json:"product_type" xorm:"not null default 0 comment('品类ID') index(INDEX_PROTYPE_PROKEY) INT(11)"`
}
