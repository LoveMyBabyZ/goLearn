package goxorm

import (
	"time"
)

type SysAuditRecords struct {
	ApplicantDate     time.Time `json:"applicant_date" xorm:"comment('申请时间') DATETIME"`
	ApplicantUserid   int       `json:"applicant_userid" xorm:"comment('申请人Id') INT(11)"`
	ApplicantUsername string    `json:"applicant_username" xorm:"comment('申请人姓名') VARCHAR(50)"`
	AuditDate         time.Time `json:"audit_date" xorm:"comment('审核时间') DATETIME"`
	AuditType         int       `json:"audit_type" xorm:"not null comment('审核类型 (1:短信 2:密钥)') INT(11)"`
	AuditUserid       int       `json:"audit_userid" xorm:"comment('审核人Id') INT(11)"`
	AuditUsername     string    `json:"audit_username" xorm:"comment('审核人姓名') VARCHAR(50)"`
	BusinessNumber    string    `json:"business_number" xorm:"comment('业务流水号') VARCHAR(50)"`
	BusinessType      int       `json:"business_type" xorm:"not null comment('业务类型(1:押金 2:储值卡 3:次卡 4:退货 11:订单商品打折)') INT(11)"`
	CustomerId        int       `json:"customer_id" xorm:"comment('会员Id') INT(11)"`
	Discount          string    `json:"discount" xorm:"default 0.00 comment('折扣') DECIMAL(18,2)"`
	Id                int       `json:"id" xorm:"not null pk autoincr comment('主键') INT(11)"`
	OrgId             int       `json:"org_id" xorm:"not null comment('机构Id') INT(11)"`
	Remark            string    `json:"remark" xorm:"comment('备注') VARCHAR(200)"`
	SmsContent        string    `json:"sms_content" xorm:"comment('短信内容') VARCHAR(500)"`
}
