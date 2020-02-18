package goxorm

import (
	"time"
)

type HisDiseaseAssistant struct {
	Content      string    `json:"content" xorm:"not null default '' comment('描述') VARCHAR(255)"`
	CreateTime   time.Time `json:"create_time" xorm:"default 'CURRENT_TIMESTAMP' comment('创建时间') DATETIME"`
	DiseaseId    int       `json:"disease_id" xorm:"not null default 0 comment('关联疾病id') INT(11)"`
	Id           int       `json:"id" xorm:"not null pk autoincr INT(11)"`
	OperatorId   int       `json:"operator_id" xorm:"comment('操作者id') INT(11)"`
	OperatorName string    `json:"operator_name" xorm:"comment('操作者 中文姓名') VARCHAR(50)"`
	Sort         int       `json:"sort" xorm:"not null default 1 comment('排序') INT(11)"`
	Status       int       `json:"status" xorm:"default 0 comment('-1:删除 0:正常') TINYINT(4)"`
	Type         int       `json:"type" xorm:"not null comment('辅助信息类型 1:常见症状, 2:体格检查') TINYINT(4)"`
	UpdateTime   time.Time `json:"update_time" xorm:"default 'CURRENT_TIMESTAMP' comment('更新时间') DATETIME"`
}
