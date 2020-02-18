package goxorm

import (
	"time"
)

type MdProductOperations struct {
	AddTime      time.Time `json:"add_time" xorm:"not null default 'CURRENT_TIMESTAMP' comment('添加时间') DATETIME"`
	HospitalName string    `json:"hospital_name" xorm:"VARCHAR(255)"`
	Id           int       `json:"id" xorm:"not null pk autoincr INT(11)"`
	IncludeType  int       `json:"include_type" xorm:"not null default 0 comment('纳入类型 0 未纳入 1 总部纳入  2区域纳入') TINYINT(3)"`
	LastModify   time.Time `json:"last_modify" xorm:"not null default 'CURRENT_TIMESTAMP' comment('修改时间') DATETIME"`
	OrgId        int       `json:"org_id" xorm:"not null comment('医院org_id') index(idx_org_ptype) INT(11)"`
	ProductType  int       `json:"product_type" xorm:"not null index(idx_org_ptype) INT(11)"`
	Status       int       `json:"status" xorm:"default 1 comment('状态0->删除，1->正常') TINYINT(2)"`
	WashType     int       `json:"wash_type" xorm:"not null default 0 comment('清洗类型 0 未清洗 1 已清洗') TINYINT(4)"`
}
