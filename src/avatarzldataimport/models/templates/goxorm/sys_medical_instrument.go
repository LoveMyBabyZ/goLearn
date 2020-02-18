package goxorm

import (
	"time"
)

type SysMedicalInstrument struct {
	CreateTime  time.Time `json:"create_time" xorm:"not null default 'CURRENT_TIMESTAMP' comment('创建时间') DATETIME"`
	FilePath    string    `json:"file_path" xorm:"not null default '' comment('文件路径') VARCHAR(1024)"`
	Id          int       `json:"id" xorm:"not null pk autoincr INT(11)"`
	Name        string    `json:"name" xorm:"not null default '' comment('文书名') VARCHAR(64)"`
	OrgId       int       `json:"org_id" xorm:"not null comment('机构/品牌id') index INT(11)"`
	Status      int       `json:"status" xorm:"not null default 0 comment('启用状态0代表未启用，1代表启用') INT(1)"`
	UpdateTime  time.Time `json:"update_time" xorm:"not null default 'CURRENT_TIMESTAMP' comment('更新时间') DATETIME"`
	Updater     string    `json:"updater" xorm:"not null default '' comment('更新人Id') VARCHAR(45)"`
	UpdaterName string    `json:"updater_name" xorm:"VARCHAR(32)"`
	Used        int       `json:"used" xorm:"default 2 comment('1代表主品牌; 2代表子品牌; 3代表机构') TINYINT(4)"`
}
