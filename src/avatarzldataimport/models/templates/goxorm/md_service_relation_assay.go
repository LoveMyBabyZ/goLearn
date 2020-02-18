package goxorm

import (
	"time"
)

type MdServiceRelationAssay struct {
	AddTime         time.Time `json:"add_time" xorm:"default 'CURRENT_TIMESTAMP' comment('添加时间') DATETIME"`
	From            int       `json:"from" xorm:"not null default 1 comment('数据来源（1：医院自建；2：数据导入）') TINYINT(1)"`
	GroupId         int       `json:"group_id" xorm:"not null default 0 comment('分组ID') INT(11)"`
	Id              int       `json:"id" xorm:"not null pk autoincr comment('主键') INT(11)"`
	IsDeleted       int       `json:"is_deleted" xorm:"not null default 0 comment('是否删除（0：未删除；1：已删除）') TINYINT(1)"`
	IsHistory       int       `json:"is_history" xorm:"not null default 0 comment('是否历史数据（0：非历史数据；1：历史数据）') TINYINT(1)"`
	LastModify      time.Time `json:"last_modify" xorm:"default 'CURRENT_TIMESTAMP' comment('修改时间') DATETIME"`
	ObjType         int       `json:"obj_type" xorm:"default 0 comment('产品对象类型') INT(11)"`
	OrgId           int       `json:"org_id" xorm:"not null default 0 comment('机构ID') index(idx_org_pid) INT(11)"`
	ProductId       int       `json:"product_id" xorm:"not null default 0 comment('产品nuan_iD') index(idx_org_pid) INT(11)"`
	ProductType     int       `json:"product_type" xorm:"not null default 0 comment('产品类型') INT(11)"`
	RelationAssayId int       `json:"relation_assay_id" xorm:"not null default 0 comment('关联化验设置nuan_ID') INT(11)"`
}
