package goxorm

import (
	"time"
)

type CrmDepositRecords struct {
	AuditorId          int       `json:"auditor_id" xorm:"default 0 comment('审核人ID') INT(11)"`
	AuditorName        string    `json:"auditor_name" xorm:"comment('审核人') VARCHAR(50)"`
	AuditorTime        time.Time `json:"auditor_time" xorm:"comment('审核时间') DATETIME"`
	Comment            string    `json:"comment" xorm:"comment('备注') VARCHAR(500)"`
	CorpusAfterChange  string    `json:"corpus_after_change" xorm:"comment('本金余额') DECIMAL(18,2)"`
	CorpusBeforeChange string    `json:"corpus_before_change" xorm:"comment('押金发生前金额') DECIMAL(18,2)"`
	CorpusChange       string    `json:"corpus_change" xorm:"comment('本金发生金额') DECIMAL(18,2)"`
	CustomerId         int       `json:"customer_id" xorm:"comment('会员Id') index(org_id) INT(11)"`
	DepositId          int       `json:"deposit_id" xorm:"comment('押金Id') index INT(11)"`
	Id                 int       `json:"id" xorm:"not null pk autoincr comment('id') INT(11)"`
	OccurOrgId         int       `json:"occur_org_id" xorm:"comment('发生地点（医院机构Id）') index(occur_org_id) INT(11)"`
	OccurTime          time.Time `json:"occur_time" xorm:"comment('发生时间') index(occur_org_id) DATETIME"`
	OperateType        int       `json:"operate_type" xorm:"comment('交易类型 ( 1：充值  2：消费扣款  3：退款给客户 )') INT(11)"`
	OperatorId         int       `json:"operator_id" xorm:"comment('操作人用户Id') INT(11)"`
	OperatorName       string    `json:"operator_name" xorm:"comment('操作人姓名') VARCHAR(50)"`
	OrgId              int       `json:"org_id" xorm:"comment('机构Id') index(org_id) INT(11)"`
	PayNumber          string    `json:"pay_number" xorm:"comment('支付流水号') VARCHAR(50)"`
	RmaComment         string    `json:"rma_comment" xorm:"comment('退款备注') VARCHAR(300)"`
}
