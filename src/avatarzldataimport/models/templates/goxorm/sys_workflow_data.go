package goxorm

import (
	"time"
)

type SysWorkflowData struct {
	BusinessData   string    `json:"business_data" xorm:"comment('业务数据') MEDIUMTEXT"`
	BusinessId     int       `json:"business_id" xorm:"not null default 0 comment('业务Id') INT(11)"`
	BusinessNumber string    `json:"business_number" xorm:"comment('业务流水号') VARCHAR(50)"`
	BusinessType   int       `json:"business_type" xorm:"comment('业务类型(1:押金 2:储值卡 3:次卡 4:退货)') INT(11)"`
	CreateDate     time.Time `json:"create_date" xorm:"comment('创建时间') DATETIME"`
	CreateUserid   int       `json:"create_userid" xorm:"comment('创建人Id') INT(11)"`
	CreateUsername string    `json:"create_username" xorm:"comment('创建人姓名') VARCHAR(50)"`
	Id             int       `json:"id" xorm:"not null pk autoincr comment('主键') INT(11)"`
	OrderId        int       `json:"order_id" xorm:"default 0 comment('订单Id') INT(11)"`
	OrgId          int       `json:"org_id" xorm:"not null default 0 comment('机构Id') INT(11)"`
	Remark         string    `json:"remark" xorm:"comment('备注') VARCHAR(200)"`
}
