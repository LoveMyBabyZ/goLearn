package goxorm

import (
	"time"
)

type SysBillCertPic struct {
	BillId         int       `json:"bill_id" xorm:"not null default 0 comment('单据Id') INT(11)"`
	BillNumber     string    `json:"bill_number" xorm:"comment('单据流水号') VARCHAR(50)"`
	BillType       int       `json:"bill_type" xorm:"comment('单据类型 (1:押金 2:储值卡 3:次卡 4:退货)') INT(11)"`
	CreateDate     time.Time `json:"create_date" xorm:"comment('创建时间') DATETIME"`
	CreateUserid   int       `json:"create_userid" xorm:"comment('创建人Id') INT(11)"`
	CreateUsername string    `json:"create_username" xorm:"comment('创建人姓名') VARCHAR(50)"`
	Id             int       `json:"id" xorm:"not null pk autoincr comment('主键') INT(11)"`
	OrgId          int       `json:"org_id" xorm:"not null default 0 comment('机构Id') INT(11)"`
	PicUrl         string    `json:"pic_url" xorm:"comment('图片URL') VARCHAR(500)"`
	Remark         string    `json:"remark" xorm:"comment('备注') VARCHAR(200)"`
}
