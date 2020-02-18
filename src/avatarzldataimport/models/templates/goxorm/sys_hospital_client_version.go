package goxorm

import (
	"time"
)

type SysHospitalClientVersion struct {
	BrandId      string    `json:"brand_id" xorm:"not null default '' comment('品牌id') VARCHAR(32)"`
	CreateTime   time.Time `json:"create_time" xorm:"not null default 'CURRENT_TIMESTAMP' comment('创建时间') DATETIME"`
	Creator      string    `json:"creator" xorm:"VARCHAR(45)"`
	CreatorName  string    `json:"creator_name" xorm:"VARCHAR(32)"`
	HospitalName string    `json:"hospital_name" xorm:"not null default '' comment('医院名称') VARCHAR(64)"`
	Id           int       `json:"id" xorm:"not null pk autoincr INT(11)"`
	InUse        int       `json:"in_use" xorm:"not null comment('是否正在使用 0 否 1 是') INT(1)"`
	OrgId        string    `json:"org_id" xorm:"not null default '' comment('医院机构id') unique(org_vid) VARCHAR(32)"`
	UpdateTime   time.Time `json:"update_time" xorm:"not null default 'CURRENT_TIMESTAMP' comment('最后修改时间') DATETIME"`
	Version      string    `json:"version" xorm:"not null default '' comment('版本号') VARCHAR(16)"`
	VersionId    int       `json:"version_id" xorm:"not null default 0 comment('版本id') unique(org_vid) INT(11)"`
}
