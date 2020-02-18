package goxorm

import (
	"time"
)

type PosPrintConfig struct {
	BusinessStatus int       `json:"business_status" xorm:"comment('单据状态（0:未勾选 1:勾选）') TINYINT(4)"`
	BusinessType   int       `json:"business_type" xorm:"comment('单据类型（1.订单 2:RMA单）') TINYINT(4)"`
	Id             int       `json:"id" xorm:"not null pk autoincr comment('主键') INT(11)"`
	OperatorDate   time.Time `json:"operator_date" xorm:"comment('操作时间') DATETIME"`
	OperatorId     int       `json:"operator_id" xorm:"comment('操作人Id') INT(11)"`
	OperatorName   string    `json:"operator_name" xorm:"comment('操作人姓名') VARCHAR(50)"`
	OrgId          int       `json:"org_id" xorm:"comment('机构Id') INT(11)"`
	Remark         string    `json:"remark" xorm:"comment('备注') VARCHAR(100)"`
}
