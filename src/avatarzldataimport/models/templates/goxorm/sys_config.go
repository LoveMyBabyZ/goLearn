package goxorm

import (
	"time"
)

type SysConfig struct {
	BrandId     int       `json:"brand_id" xorm:"not null comment('品牌id') unique(company_unique_config) INT(11)"`
	Content     string    `json:"content" xorm:"not null comment('配置内容') TEXT"`
	CreateTime  time.Time `json:"create_time" xorm:"not null default 'CURRENT_TIMESTAMP' comment('创建时间') DATETIME"`
	Creator     string    `json:"creator" xorm:"not null default '' comment('创建人') VARCHAR(45)"`
	CreatorName string    `json:"creator_name" xorm:"not null default '' comment('穿件人姓名') VARCHAR(32)"`
	Id          int       `json:"id" xorm:"not null pk autoincr INT(11)"`
	Type        int       `json:"type" xorm:"not null default 0 comment('配置项（l1代表ogo）') unique(company_unique_config) INT(1)"`
	UpdateTime  time.Time `json:"update_time" xorm:"not null default 'CURRENT_TIMESTAMP' comment('更新时间') DATETIME"`
	Used        int       `json:"used" xorm:"default 2 comment('1代表主品牌; 2代表子品牌; 3代表机构') unique(company_unique_config) TINYINT(4)"`
}
