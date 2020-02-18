package goxorm

import (
	"time"
)

type SysPrint struct {
	CreateTime  time.Time `json:"create_time" xorm:"not null default 'CURRENT_TIMESTAMP' DATETIME"`
	Id          int       `json:"id" xorm:"not null pk autoincr INT(11)"`
	Name        string    `json:"name" xorm:"not null default '' comment('打印项目名') VARCHAR(128)"`
	Type        int       `json:"type" xorm:"not null default 1 comment('类型，暂时不用') INT(11)"`
	Updater     string    `json:"updater" xorm:"not null default '' VARCHAR(45)"`
	UpdaterName string    `json:"updater_name" xorm:"not null default '' VARCHAR(32)"`
}
