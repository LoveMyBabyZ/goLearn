package goxorm

import (
	"time"
)

type MdPackage struct {
	AddTime        time.Time `json:"add_time" xorm:"not null default 'CURRENT_TIMESTAMP' comment('创建时间') TIMESTAMP"`
	ChainPackageId int       `json:"chain_package_id" xorm:"comment('连锁套餐PID') INT(11)"`
	ChangeStatus   int       `json:"change_status" xorm:"not null default 1 comment('单院可修改定价（1=>可以，2=>不可以）') TINYINT(2)"`
	EndTime        time.Time `json:"end_time" xorm:"comment('连锁套餐生效结束时间') TIMESTAMP"`
	Id             int       `json:"id" xorm:"not null pk autoincr comment('主键') INT(11)"`
	IssuedId       int       `json:"issued_id" xorm:"comment('下发通知ID') INT(11)"`
	LastModify     time.Time `json:"last_modify" xorm:"not null default 'CURRENT_TIMESTAMP' comment('最后更新时间') TIMESTAMP"`
	NuanId         int       `json:"nuan_id" xorm:"not null default 0 comment('小暖ID') unique(idx_org_tpl) INT(11)"`
	OrgId          int       `json:"org_id" xorm:"not null default 0 comment('机构ID') unique(idx_org_tpl) INT(11)"`
	OriginalPrice  string    `json:"original_price" xorm:"not null default 0.0000 comment('套餐原价') DECIMAL(30,4)"`
	PackageCode    string    `json:"package_code" xorm:"not null comment('套餐编码') VARCHAR(255)"`
	PackageDesc    string    `json:"package_desc" xorm:"comment('"适用说明"') TEXT"`
	PackageName    string    `json:"package_name" xorm:"not null comment('套餐名称') VARCHAR(255)"`
	PackagePrice   string    `json:"package_price" xorm:"not null comment('套餐定价') DECIMAL(30,4)"`
	PackageType    int       `json:"package_type" xorm:"not null default 1 comment('套餐类型（1=>单院套餐，2=>连锁套餐）') TINYINT(2)"`
	StartTime      time.Time `json:"start_time" xorm:"comment('连锁套餐开始生效时间') TIMESTAMP"`
	Status         int       `json:"status" xorm:"not null default 1 comment('状态（1=>正常，2=>删除）') TINYINT(4)"`
}
