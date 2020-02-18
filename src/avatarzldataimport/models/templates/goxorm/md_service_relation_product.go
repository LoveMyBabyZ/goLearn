package goxorm

import (
	"time"
)

type MdServiceRelationProduct struct {
	AddTime             time.Time `json:"add_time" xorm:"default 'CURRENT_TIMESTAMP' comment('添加时间') DATETIME"`
	Amount              int       `json:"amount" xorm:"not null default 0 comment('数量') INT(11)"`
	Days                int       `json:"days" xorm:"not null default 0 comment('天数') INT(11)"`
	DosageOnce          int       `json:"dosage_once" xorm:"not null default 0 comment('单次用量') INT(11)"`
	DosingWayId         int       `json:"dosing_way_id" xorm:"not null default 0 comment('用法') INT(11)"`
	From                int       `json:"from" xorm:"not null default 1 comment('数据来源（1：医院自建；2：数据导入）') TINYINT(1)"`
	GroupId             int       `json:"group_id" xorm:"not null default 0 comment('分组ID') INT(11)"`
	Id                  int       `json:"id" xorm:"not null pk autoincr comment('主键') INT(11)"`
	IsAdd               int       `json:"is_add" xorm:"not null default 0 comment('是否费用累加项(0->非累加;1->累加)') TINYINT(2)"`
	IsDeleted           int       `json:"is_deleted" xorm:"not null default 0 comment('是否删除（0：未删除；1：已删除）') TINYINT(1)"`
	IsHistory           int       `json:"is_history" xorm:"not null default 0 comment('是否历史数据（0：非历史数据；1：历史数据）') TINYINT(1)"`
	IsMust              int       `json:"is_must" xorm:"not null default 0 comment('是否必选项(0->非必选;1->必选)') TINYINT(2)"`
	LastModify          time.Time `json:"last_modify" xorm:"default 'CURRENT_TIMESTAMP' comment('修改时间') DATETIME"`
	ObjType             int       `json:"obj_type" xorm:"default 0 comment('产品对象类型') INT(11)"`
	OrgId               int       `json:"org_id" xorm:"not null default 0 comment('机构ID') index(idx_org_pid) INT(11)"`
	ProductId           int       `json:"product_id" xorm:"not null default 0 comment('产品nuan_ID') index(idx_org_pid) INT(11)"`
	ProductType         int       `json:"product_type" xorm:"not null default 0 comment('产品类型') INT(11)"`
	RelationProductId   int       `json:"relation_product_id" xorm:"not null default 0 comment('关联产品nuan_id') INT(11)"`
	RelationProductName string    `json:"relation_product_name" xorm:"not null default '' comment('关联产品名称') VARCHAR(200)"`
	TimesOneDay         int       `json:"times_one_day" xorm:"not null default 0 comment('次/天') INT(11)"`
}
