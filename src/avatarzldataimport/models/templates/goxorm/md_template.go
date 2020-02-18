package goxorm

import (
	"time"
)

type MdTemplate struct {
	AddTime        time.Time `json:"add_time" xorm:"not null default 'CURRENT_TIMESTAMP' comment('添加时间') TIMESTAMP"`
	DiseasesId     int       `json:"diseases_id" xorm:"comment('病种类型') INT(11)"`
	DiseasesName   string    `json:"diseases_name" xorm:"comment('病种类型名称') VARCHAR(30)"`
	Id             int       `json:"id" xorm:"not null pk autoincr comment('主键') INT(11)"`
	LastModify     time.Time `json:"last_modify" xorm:"not null default 'CURRENT_TIMESTAMP' comment('模板定价') TIMESTAMP"`
	NuanId         int       `json:"nuan_id" xorm:"not null default 0 comment('小暖ID') unique(idx_org_tpl) INT(11)"`
	OrgId          int       `json:"org_id" xorm:"not null default 0 comment('机构ID') unique(idx_org_tpl) INT(11)"`
	Price          string    `json:"price" xorm:"not null default 0.0000 comment('模板价格') DECIMAL(30,4)"`
	Source         int       `json:"source" xorm:"not null default 1 comment('来源（1=>自建，2=>连锁）') TINYINT(3)"`
	TemplateCode   string    `json:"template_code" xorm:"not null comment('编号') VARCHAR(255)"`
	TemplateDesc   string    `json:"template_desc" xorm:"comment('适用说明') TEXT"`
	TemplateName   string    `json:"template_name" xorm:"not null comment('模板名称') VARCHAR(255)"`
	TemplateStatus int       `json:"template_status" xorm:"not null default 1 comment('状态是否删除启用（1=>启用，2=>删除）') TINYINT(2)"`
	TemplateType   int       `json:"template_type" xorm:"not null comment('模板类型（922001=>处方，922002=>实验室检查，922003=>影像学检查）') INT(11)"`
}
