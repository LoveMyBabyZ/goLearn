package goxorm

import (
	"time"
)

type HisDiseaseTreatment struct {
	Content      string    `json:"content" xorm:"comment('治疗建议内容') TEXT"`
	CreateTime   time.Time `json:"create_time" xorm:"default 'CURRENT_TIMESTAMP' comment('创建时间') DATETIME"`
	DiseaseId    int       `json:"disease_id" xorm:"not null default 0 comment('关联疾病id') INT(11)"`
	Id           int       `json:"id" xorm:"not null pk autoincr INT(11)"`
	OperatorId   int       `json:"operator_id" xorm:"comment('操作者id') INT(11)"`
	OperatorName string    `json:"operator_name" xorm:"comment('操作者 中文姓名') VARCHAR(50)"`
	Status       int       `json:"status" xorm:"default 0 comment('-1:删除 0:正常') TINYINT(4)"`
	Type         int       `json:"type" xorm:"not null comment('治疗建议类型 1:检查建议, 2:诊断要点, 3:治疗建议') TINYINT(4)"`
	UpdateTime   time.Time `json:"update_time" xorm:"default 'CURRENT_TIMESTAMP' comment('更新时间') DATETIME"`
}
