package goxorm

import (
	"time"
)

type MdPackageIssued struct {
	AddTime           time.Time `json:"add_time" xorm:"default 'CURRENT_TIMESTAMP' TIMESTAMP"`
	CenterHospitalId  int       `json:"center_hospital_id" xorm:"not null comment('中心医院ID') INT(11)"`
	ChainId           int       `json:"chain_id" xorm:"default 0 comment('连锁ID') INT(11)"`
	ChainPackageId    int       `json:"chain_package_id" xorm:"not null comment('连锁套餐ID') INT(11)"`
	Id                int       `json:"id" xorm:"not null pk autoincr INT(11)"`
	LastModify        time.Time `json:"last_modify" xorm:"default 'CURRENT_TIMESTAMP' TIMESTAMP"`
	NoticeStatus      int       `json:"notice_status" xorm:"not null default 1 comment('下发通知状态（1=>未接收，2=>已接收，3=>已失效，4=>已删除）') TINYINT(3)"`
	PackageId         int       `json:"package_id" xorm:"comment('生成的单院套餐ID') INT(11)"`
	PackageName       string    `json:"package_name" xorm:"not null comment('连锁套餐名称') VARCHAR(255)"`
	ReceiveHospitalId int       `json:"receive_hospital_id" xorm:"not null comment('接收医院ID') INT(11)"`
}
