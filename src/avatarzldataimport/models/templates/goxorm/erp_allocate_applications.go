package goxorm

import (
	"time"
)

type ErpAllocateApplications struct {
	AuditorId          int       `json:"auditor_id" xorm:"default 0 comment('申请机构审核人Id') INT(11)"`
	AuditorName        string    `json:"auditor_name" xorm:"default '' comment('申请机构审核人名称') VARCHAR(50)"`
	AuditorTime        time.Time `json:"auditor_time" xorm:"default 'CURRENT_TIMESTAMP' comment('申请机构审核时间') DATETIME"`
	CreateOrgId        int       `json:"create_org_id" xorm:"default 0 comment('申请机构ID') INT(11)"`
	CreateOrgName      string    `json:"create_org_name" xorm:"default '' comment('申请机构名称') VARCHAR(100)"`
	CreateRemark       string    `json:"create_remark" xorm:"default '' comment('申请方备注') VARCHAR(500)"`
	CreateTime         time.Time `json:"create_time" xorm:"default 'CURRENT_TIMESTAMP' comment('创建时间') DATETIME"`
	CreateUserId       int       `json:"create_user_id" xorm:"default 0 comment('创建人ID') INT(11)"`
	CreateUserName     string    `json:"create_user_name" xorm:"default '' comment('创建人名称') VARCHAR(50)"`
	Id                 int       `json:"id" xorm:"not null pk autoincr comment('id') INT(11)"`
	ModifyTime         time.Time `json:"modify_time" xorm:"default 'CURRENT_TIMESTAMP' comment('修改人时间') DATETIME"`
	ModifyUserId       int       `json:"modify_user_id" xorm:"default 0 comment('修改人ID') INT(11)"`
	ModifyUserName     string    `json:"modify_user_name" xorm:"default '' comment('修改人名称') VARCHAR(50)"`
	OrderNumber        string    `json:"order_number" xorm:"default '' comment('单据号') VARCHAR(100)"`
	SourceAuditor1Id   int       `json:"source_auditor1_id" xorm:"default 0 comment('来源机构初审人id') INT(11)"`
	SourceAuditor1Name string    `json:"source_auditor1_name" xorm:"default '' comment('来源机构初审人名称') VARCHAR(50)"`
	SourceAuditor1Time time.Time `json:"source_auditor1_time" xorm:"default 'CURRENT_TIMESTAMP' comment('来源机构初审时间') DATETIME"`
	SourceAuditor2Id   int       `json:"source_auditor2_id" xorm:"default 0 comment('来源机构复审人id') INT(11)"`
	SourceAuditor2Name string    `json:"source_auditor2_name" xorm:"default '' comment('来源机构复审人名称') VARCHAR(50)"`
	SourceAuditor2Time time.Time `json:"source_auditor2_time" xorm:"default 'CURRENT_TIMESTAMP' comment('来源机构复审时间') DATETIME"`
	SourceOrgId        int       `json:"source_org_id" xorm:"default 0 comment('来源机构ID') INT(11)"`
	SourceOrgName      string    `json:"source_org_name" xorm:"default '' comment('来源机构名称') VARCHAR(100)"`
	SourceRemark       string    `json:"source_remark" xorm:"default '' comment('来源方备注') VARCHAR(500)"`
	Status             int       `json:"status" xorm:"default 0 comment('状态 0:保存  10：已提交  20：发起方审核   30：接收方初审    40：接收方复审') INT(11)"`
	UpdateTimestamp    time.Time `json:"update_timestamp" xorm:"not null default 'CURRENT_TIMESTAMP(6)' comment('更新时间戳') TIMESTAMP(6)"`
}
