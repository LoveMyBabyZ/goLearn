package goxorm

import (
	"time"
)

type MdProductType struct {
	AddTime    time.Time `json:"add_time" xorm:"default 'CURRENT_TIMESTAMP' comment('添加时间') DATETIME"`
	Category   int       `json:"category" xorm:"not null default 1 comment('所属大类：1-库存类  2-服务类') TINYINT(1)"`
	Id         int       `json:"id" xorm:"not null pk autoincr comment('主键') INT(11)"`
	LastModify time.Time `json:"last_modify" xorm:"default 'CURRENT_TIMESTAMP' comment('修改时间') DATETIME"`
	TypeCode   int       `json:"type_code" xorm:"not null default 0 comment('产品类型编码') INT(11)"`
	TypeName   string    `json:"type_name" xorm:"not null default '' comment('产品类型名称') VARCHAR(255)"`
}
