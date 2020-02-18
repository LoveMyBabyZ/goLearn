package goxorm

import (
	"time"
)

type MdProductCategory struct {
	AddTime       time.Time `json:"add_time" xorm:"default 'CURRENT_TIMESTAMP' comment('添加时间') DATETIME"`
	CategoryLevel int       `json:"category_level" xorm:"not null default 1 comment('目录层级') TINYINT(2)"`
	CategoryName  string    `json:"category_name" xorm:"not null comment('目录名称') VARCHAR(255)"`
	From          int       `json:"from" xorm:"not null default 1 comment('数据来源（1：医院自建；2：数据导入）') TINYINT(1)"`
	Id            int       `json:"id" xorm:"not null pk autoincr INT(11)"`
	IdPath        string    `json:"id_path" xorm:"not null default '' comment('目录id路径') VARCHAR(200)"`
	IsDeleted     int       `json:"is_deleted" xorm:"not null default 0 comment('已经删除') TINYINT(1)"`
	IsHistory     int       `json:"is_history" xorm:"not null default 0 comment('是否历史数据（0：非历史数据；1：历史数据）') TINYINT(1)"`
	ItemCode      string    `json:"item_code" xorm:"not null default '' comment('总部编码') VARCHAR(100)"`
	LastModify    time.Time `json:"last_modify" xorm:"default 'CURRENT_TIMESTAMP' comment('修改时间') DATETIME"`
	NamePath      string    `json:"name_path" xorm:"not null default '' comment('目录名称路径') VARCHAR(200)"`
	NuanId        int       `json:"nuan_id" xorm:"not null comment('小暖ID') index(idx_orgid_nuanid) INT(11)"`
	OrgId         int       `json:"org_id" xorm:"not null comment('医院ID') index(idx_orgid_nuanid) index(idx_orgid_ptype) INT(11)"`
	ParentId      int       `json:"parent_id" xorm:"not null default 0 comment('上级目录ID') INT(11)"`
	ProductType   int       `json:"product_type" xorm:"not null comment('产品类型') index(idx_orgid_ptype) INT(11)"`
}
