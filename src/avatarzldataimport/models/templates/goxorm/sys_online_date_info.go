package goxorm

import (
	"time"
)

type SysOnlineDateInfo struct {
	Id         int       `json:"id" xorm:"not null pk autoincr comment('主键') INT(11)"`
	OnlineDate time.Time `json:"online_date" xorm:"comment('上线日期') DATETIME"`
	OrgCode    string    `json:"org_code" xorm:"comment('机构编码') VARCHAR(50)"`
	OrgId      int       `json:"org_id" xorm:"comment('机构id') INT(11)"`
	ReportDate string    `json:"report_date" xorm:"comment('报表开始统计日期') VARCHAR(50)"`
}
