package goxorm

import (
	"time"
)

type THosCompany struct {
	AddTime      time.Time `json:"add_time" xorm:"default 'CURRENT_TIMESTAMP' comment('添加时间') DATETIME"`
	Id           int       `json:"id" xorm:"not null pk autoincr INT(11)"`
	StructCode   string    `json:"struct_code" xorm:"not null comment('公司编码，权限系统对外暴露的Code') VARCHAR(32)"`
	StructId     int       `json:"struct_id" xorm:"not null comment('公司ID') INT(11)"`
	StructName   string    `json:"struct_name" xorm:"default '' comment('公司名称') VARCHAR(50)"`
	StructStatus int       `json:"struct_status" xorm:"default 0 comment('状态 0正常 1禁用') TINYINT(3)"`
}
