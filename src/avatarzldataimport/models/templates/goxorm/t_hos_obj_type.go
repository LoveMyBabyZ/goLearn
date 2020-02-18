package goxorm

import (
	"time"
)

type THosObjType struct {
	AddTime         time.Time `json:"add_time" xorm:"default 'CURRENT_TIMESTAMP' comment('添加时间') DATETIME"`
	Id              int       `json:"id" xorm:"not null pk autoincr comment('主键') INT(11)"`
	LastModify      time.Time `json:"last_modify" xorm:"default 'CURRENT_TIMESTAMP' comment('修改时间') DATETIME"`
	ObjCode         int       `json:"obj_code" xorm:"not null default 0 comment('对象类型编码') INT(11)"`
	ObjName         string    `json:"obj_name" xorm:"not null default '' comment('对象类型名称') VARCHAR(15)"`
	ProductTypeCode int       `json:"product_type_code" xorm:"not null default 0 comment('产品类型编码') INT(11)"`
	ProductTypeName string    `json:"product_type_name" xorm:"not null default '' comment('产品类型名称') VARCHAR(15)"`
}
