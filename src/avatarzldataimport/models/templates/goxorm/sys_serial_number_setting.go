package goxorm

import (
	"time"
)

type SysSerialNumberSetting struct {
	CreateName string    `json:"create_name" xorm:"comment('创建人') VARCHAR(50)"`
	CreateTime time.Time `json:"create_time" xorm:"comment('创建时间') TIMESTAMP"`
	CurDate    string    `json:"cur_date" xorm:"comment('当前日期') VARCHAR(10)"`
	CurValue   int       `json:"cur_value" xorm:"default 0 comment('序列号当前值') INT(10)"`
	EveryDay   int       `json:"every_day" xorm:"default 1 comment('每天生成，默认1是，0否') TINYINT(1)"`
	Id         int       `json:"id" xorm:"not null pk autoincr comment('主键') INT(11)"`
	InitValue  int       `json:"init_value" xorm:"default 1 comment('序列号初始值') INT(10)"`
	NoLength   int       `json:"no_length" xorm:"default 9 comment('序列号长度') INT(10)"`
	OrgId      int       `json:"org_id" xorm:"comment('机构id') index INT(11)"`
	OrgName    string    `json:"org_name" xorm:"comment('机构名称') VARCHAR(100)"`
	Rule       string    `json:"rule" xorm:"comment('流水号生成规则') VARCHAR(255)"`
	Setp       int       `json:"setp" xorm:"default 1 comment('步长') TINYINT(1)"`
	Type       int       `json:"type" xorm:"default 0 comment('流水号类型，0卡号，1、充值，2、消费，3、退款，4、退卡、5、支付流水号、6、退款流水号') INT(11)"`
	UpdateName string    `json:"update_name" xorm:"comment('修改人') VARCHAR(50)"`
	UpdateTime time.Time `json:"update_time" xorm:"comment('修改时间') TIMESTAMP"`
}
