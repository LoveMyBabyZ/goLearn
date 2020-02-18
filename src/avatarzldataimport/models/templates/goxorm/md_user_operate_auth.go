package goxorm

import (
	"time"
)

type MdUserOperateAuth struct {
	AuthCode   string    `json:"auth_code" xorm:"not null comment('菜单权限的编码') VARCHAR(255)"`
	AuthName   string    `json:"auth_name" xorm:"not null comment('菜单权限的名字') VARCHAR(255)"`
	ClassName  string    `json:"class_name" xorm:"not null default '' comment('对应的操作的类名,  前面带着目录避免重复，例如：Log/Lists') index VARCHAR(100)"`
	CreateTime time.Time `json:"create_time" xorm:"not null default 'CURRENT_TIMESTAMP' TIMESTAMP"`
	ExtraType  string    `json:"extra_type" xorm:"not null default '' comment('额外类型参数') VARCHAR(255)"`
	Id         int       `json:"id" xorm:"not null pk autoincr INT(11)"`
	Status     int       `json:"status" xorm:"not null default 1 comment('权限是否失效 0代表失效  1代表启用') TINYINT(1)"`
	UpdateTime time.Time `json:"update_time" xorm:"not null default 'CURRENT_TIMESTAMP' TIMESTAMP"`
}
