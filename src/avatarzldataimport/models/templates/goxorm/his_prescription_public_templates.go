package goxorm

import (
	"time"
)

type HisPrescriptionPublicTemplates struct {
	CreateDate   time.Time `json:"create_date" xorm:"comment('创建时间') DATETIME"`
	Current      int       `json:"current" xorm:"comment('是否当前使用') INT(11)"`
	HashCode     string    `json:"hash_code" xorm:"comment('模版文件hash值') VARCHAR(100)"`
	Id           int       `json:"id" xorm:"not null pk autoincr comment('id') INT(11)"`
	TemplateName string    `json:"template_name" xorm:"comment('模版名称') VARCHAR(100)"`
	UpdateDate   time.Time `json:"update_date" xorm:"comment('更新时间') DATETIME"`
	Url          string    `json:"url" xorm:"comment('模版路径(相对)') VARCHAR(200)"`
}
