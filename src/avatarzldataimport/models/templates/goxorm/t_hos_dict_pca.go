package goxorm

import (
	"time"
)

type THosDictPca struct {
	AddTime        time.Time `json:"add_time" xorm:"default 'CURRENT_TIMESTAMP' comment('添加时间') DATETIME"`
	Area           string    `json:"area" xorm:"default '' comment('区域名称') unique(UNIQUE_PRO_CITY_AREA) VARCHAR(50)"`
	AreaCode       string    `json:"area_code" xorm:"not null default '' comment('省市区域code') unique VARCHAR(10)"`
	AreaType       int       `json:"area_type" xorm:"not null default 1 comment('区域级别 1 省 2 市 3区') TINYINT(2)"`
	City           string    `json:"city" xorm:"default '' comment('城市名称') unique(UNIQUE_PRO_CITY_AREA) VARCHAR(50)"`
	Id             int       `json:"id" xorm:"not null pk autoincr comment('自增id') INT(11)"`
	LastModify     time.Time `json:"last_modify" xorm:"default 'CURRENT_TIMESTAMP' comment('更新时间') DATETIME"`
	ParentAreaCode string    `json:"parent_area_code" xorm:"not null default '' comment('父级省市区area code') VARCHAR(10)"`
	Province       string    `json:"province" xorm:"not null default '' comment('省份名称') unique(UNIQUE_PRO_CITY_AREA) VARCHAR(50)"`
}
