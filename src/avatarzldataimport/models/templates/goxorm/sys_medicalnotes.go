package goxorm

import (
	"time"
)

type SysMedicalnotes struct {
	Author      int       `json:"author" xorm:"not null comment('作者ID') INT(11)"`
	AuthorName  string    `json:"author_name" xorm:"not null comment('作者') VARCHAR(100)"`
	BrandId     int       `json:"brand_id" xorm:"comment('品牌ID') INT(11)"`
	BrandName   string    `json:"brand_name" xorm:"comment('品牌') VARCHAR(100)"`
	Code        string    `json:"code" xorm:"not null comment('编号') VARCHAR(50)"`
	CreateDate  time.Time `json:"create_date" xorm:"not null comment('创建时间') DATETIME"`
	Id          int       `json:"id" xorm:"not null pk autoincr INT(11)"`
	Name        string    `json:"name" xorm:"not null comment('名称') VARCHAR(100)"`
	OrgId       int       `json:"org_id" xorm:"not null comment('机构ID') index INT(11)"`
	Status      int       `json:"status" xorm:"default 1 comment('状态') INT(11)"`
	UpdateDate  time.Time `json:"update_date" xorm:"not null comment('修改时间') DATETIME"`
	Updater     int       `json:"updater" xorm:"not null comment('修改者ID') INT(11)"`
	UpdaterName string    `json:"updater_name" xorm:"not null comment('修改者') VARCHAR(50)"`
	Url         string    `json:"Url" xorm:"not null comment('文件路径') VARCHAR(256)"`
}
