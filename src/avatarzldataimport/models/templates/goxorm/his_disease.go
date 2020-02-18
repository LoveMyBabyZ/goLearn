package goxorm

import (
	"time"
)

type HisDisease struct {
	Code            string    `json:"code" xorm:"not null default '' comment('病种编号') VARCHAR(10)"`
	CreateTime      time.Time `json:"create_time" xorm:"default 'CURRENT_TIMESTAMP' comment('创建时间') DATETIME"`
	DiseaseType     int       `json:"disease_type" xorm:"not null comment('病种类型 1:总部端,2:医院端') TINYINT(4)"`
	Enable          int       `json:"enable" xorm:"not null default 1 comment('状态 1:启用 0:禁用') TINYINT(4)"`
	Id              int       `json:"id" xorm:"not null pk autoincr comment('自增id') INT(11)"`
	Level           int       `json:"level" xorm:"not null default 0 comment('病种/分类层级') TINYINT(4)"`
	Name            string    `json:"name" xorm:"not null default '' comment('病种/分类名称') VARCHAR(50)"`
	NameEn          string    `json:"name_en" xorm:"not null default '' comment('病种英文名称') VARCHAR(100)"`
	NameFirstPinyin string    `json:"name_first_pinyin" xorm:"not null default '' comment('病种拼音首字母') VARCHAR(50)"`
	NameFullPinyin  string    `json:"name_full_pinyin" xorm:"not null default '' comment('病种全拼') VARCHAR(200)"`
	OperatorId      int       `json:"operator_id" xorm:"comment('操作者id') INT(11)"`
	OperatorName    string    `json:"operator_name" xorm:"comment('操作者 中文姓名') VARCHAR(50)"`
	Orgid           int       `json:"orgid" xorm:"not null default 0 comment('机构(医院)id') INT(11)"`
	Pid             int       `json:"pid" xorm:"not null default 0 comment('父病种id') INT(11)"`
	Status          int       `json:"status" xorm:"default 0 comment('-1:删除 0:正常') TINYINT(4)"`
	UpdateTime      time.Time `json:"update_time" xorm:"default 'CURRENT_TIMESTAMP' comment('更新时间') DATETIME"`
}
