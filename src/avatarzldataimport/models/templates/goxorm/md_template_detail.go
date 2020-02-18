package goxorm

import (
	"time"
)

type MdTemplateDetail struct {
	AddTime        time.Time `json:"add_time" xorm:"not null default 'CURRENT_TIMESTAMP' comment('创建时间') TIMESTAMP"`
	DosingUnitId   int       `json:"dosing_unit_id" xorm:"comment('投药单位') INT(11)"`
	DosingUnitName string    `json:"dosing_unit_name" xorm:"comment('投药单位名称') VARCHAR(30)"`
	DosingWayId    int       `json:"dosing_way_id" xorm:"default 0 comment('投药方式') INT(11)"`
	DosingWayName  string    `json:"dosing_way_name" xorm:"comment('投药方式描述') VARCHAR(32)"`
	Group          int       `json:"group" xorm:"comment('分组') INT(11)"`
	Id             int       `json:"id" xorm:"not null pk autoincr comment('主键') INT(11)"`
	ItemCode       string    `json:"item_code" xorm:"default '0' comment('产品总部编码') VARCHAR(50)"`
	LastModify     time.Time `json:"last_modify" xorm:"not null default 'CURRENT_TIMESTAMP' comment('最后更新时间') TIMESTAMP"`
	NumDays        int       `json:"num_days" xorm:"not null default 1 comment('天数') INT(11)"`
	OrgId          int       `json:"org_id" xorm:"not null default 0 comment('机构ID') index(idx_org_proid) index(idx_org_tplid) INT(11)"`
	Price          string    `json:"price" xorm:"not null default 0.0000 comment('单价') DECIMAL(30,4)"`
	ProductId      int       `json:"product_id" xorm:"not null default 0 comment('产品nuan_id') index(idx_org_proid) INT(11)"`
	ProductName    string    `json:"product_name" xorm:"not null comment('产品名称') VARCHAR(255)"`
	SingleDose     string    `json:"single_dose" xorm:"not null default 0.00000 comment('单次用量') DECIMAL(19,5)"`
	Specific       string    `json:"specific" xorm:"comment('产品规格') VARCHAR(255)"`
	Status         int       `json:"status" xorm:"not null default 1 comment('状态（1=>启用，2=>删除）') TINYINT(11)"`
	StoreUnitId    int       `json:"store_unit_id" xorm:"comment('出库单位') INT(11)"`
	StoreUnitName  string    `json:"store_unit_name" xorm:"comment('出入库单位名称') VARCHAR(32)"`
	TemplateId     int       `json:"template_id" xorm:"not null comment('模板关联ID') index(idx_org_tplid) INT(11)"`
	TimeDay        int       `json:"time_day" xorm:"not null comment('次/天') INT(11)"`
	TotalAmount    string    `json:"total_amount" xorm:"not null default 0.0000 comment('总金额') DECIMAL(30,4)"`
	TotalNumber    int       `json:"total_number" xorm:"not null default 1 comment('数量') INT(11)"`
	Type           int       `json:"type" xorm:"not null default 1 comment('组合类型（1=>模板）') TINYINT(2)"`
}
