package goxorm

import (
	"time"
)

type MdColumnConfig struct {
	AddTime       time.Time `json:"add_time" xorm:"default 'CURRENT_TIMESTAMP' comment('添加时间') DATETIME"`
	DataType      int       `json:"data_type" xorm:"not null default 0 comment('数据类型') TINYINT(2)"`
	FieldType     int       `json:"field_type" xorm:"not null default 1 comment('字段类型（1：主产品字段；2：关联产品字段；3：关联化验字段）') TINYINT(1)"`
	Id            int       `json:"id" xorm:"not null pk autoincr comment('主键') INT(11)"`
	IsHeadControl int       `json:"is_head_control" xorm:"not null default 0 comment('是否总部管控字段') TINYINT(1)"`
	IsRequired    int       `json:"is_required" xorm:"not null default 0 comment('是否必填 1 必填 0 非必填') TINYINT(1)"`
	LastModify    time.Time `json:"last_modify" xorm:"default 'CURRENT_TIMESTAMP' comment('修改时间') DATETIME"`
	ObjCol        string    `json:"obj_col" xorm:"not null default '' comment('字段') VARCHAR(200)"`
	ObjColName    string    `json:"obj_col_name" xorm:"not null default '' comment('字段名') VARCHAR(200)"`
	ProductType   int       `json:"product_type" xorm:"not null default 0 comment('产品类型') index INT(11)"`
}
