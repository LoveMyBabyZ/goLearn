package goxorm

import (
	"time"
)

type MdMenuConfig struct {
	AddTime    time.Time `json:"add_time" xorm:"default 'CURRENT_TIMESTAMP' comment('添加时间') TIMESTAMP"`
	AuthCode   string    `json:"auth_code" xorm:"not null default '' comment('权限点编码') VARCHAR(255)"`
	BrandId    int       `json:"brand_id" xorm:"not null default 0 comment('品牌ID') INT(11)"`
	Code       string    `json:"code" xorm:"not null default '' comment('menu编码') index VARCHAR(100)"`
	Id         int       `json:"id" xorm:"not null pk autoincr comment('自增ID') INT(11)"`
	Name       string    `json:"name" xorm:"not null default '' comment('menu名称') VARCHAR(100)"`
	ParentCode string    `json:"parent_code" xorm:"not null default '' comment('父级编码') VARCHAR(100)"`
	Status     int       `json:"status" xorm:"not null default 1 comment('状态 0 禁用 1 启用') TINYINT(1)"`
	Type       string    `json:"type" xorm:"not null default '' comment('menu 类型') VARCHAR(32)"`
	Url        string    `json:"url" xorm:"not null default '' comment('跳转地址') VARCHAR(255)"`
}
