package goxorm

import (
	"time"
)

type SysIssuePrint struct {
	BrandId    int       `json:"brand_id" xorm:"comment('品牌id') INT(11)"`
	Id         int       `json:"id" xorm:"not null pk autoincr INT(11)"`
	Orders     int       `json:"orders" xorm:"not null comment('序值') INT(11)"`
	PrintId    int       `json:"print_id" xorm:"not null comment('小票id') INT(11)"`
	Status     int       `json:"status" xorm:"not null comment('状态') INT(1)"`
	UpdateTime time.Time `json:"update_time" xorm:"not null default 'CURRENT_TIMESTAMP' comment('更新时间') DATETIME"`
	Updater    string    `json:"updater" xorm:"not null default '' comment('更新人') VARCHAR(45)"`
}
