package goxorm

import (
	"time"
)

type CrmCardsRecords struct {
	AuditorId                int       `json:"auditor_id" xorm:"default 0 comment('审核人ID') INT(11)"`
	AuditorName              string    `json:"auditor_name" xorm:"comment('审核人') VARCHAR(50)"`
	AuditorTime              time.Time `json:"auditor_time" xorm:"comment('审核时间') DATETIME"`
	CardId                   int       `json:"card_id" xorm:"comment('会员卡Id') index INT(11)"`
	Comment                  string    `json:"comment" xorm:"comment('备注') VARCHAR(500)"`
	CorpusAfterChange        string    `json:"corpus_after_change" xorm:"comment('本金余额') DECIMAL(18,2)"`
	CorpusBeforeChange       string    `json:"corpus_before_change" xorm:"comment('本金发生前金额') DECIMAL(18,2)"`
	CorpusChange             string    `json:"corpus_change" xorm:"comment('本金发生金额') DECIMAL(18,2)"`
	CustomerId               int       `json:"customer_id" xorm:"comment('会员Id') index index(org_id) INT(11)"`
	Id                       int       `json:"id" xorm:"not null pk autoincr comment('id') INT(11)"`
	OccurOrgId               int       `json:"occur_org_id" xorm:"comment('发生地点（医院机构Id）') index(idx_occur_org_id_occur_time) INT(11)"`
	OccurOrgName             string    `json:"occur_org_name" xorm:"comment('发生地点（医院机构名称）') VARCHAR(255)"`
	OccurTime                time.Time `json:"occur_time" xorm:"comment('发生时间') index(idx_occur_org_id_occur_time) index DATETIME"`
	OperateType              int       `json:"operate_type" xorm:"comment('交易类型 ( 1：开卡  2：充值  3：消费扣款  4：消费退款（增加金额）  5：退款给客户   6：客户退卡)') INT(11)"`
	OperatorId               int       `json:"operator_id" xorm:"comment('操作人用户Id') INT(11)"`
	OperatorName             string    `json:"operator_name" xorm:"comment('操作人姓名') VARCHAR(50)"`
	OrgId                    int       `json:"org_id" xorm:"comment('机构Id') index(org_id) INT(11)"`
	PayNumber                string    `json:"pay_number" xorm:"comment('支付流水号') VARCHAR(50)"`
	PresentAfterChange       string    `json:"present_after_change" xorm:"comment('赠送部分余额') DECIMAL(18,2)"`
	PresentBeforeChange      string    `json:"present_before_change" xorm:"comment('赠送部分发生前金额') DECIMAL(18,2)"`
	PresentChange            string    `json:"present_change" xorm:"comment('赠送部分发生金额') DECIMAL(18,2)"`
	PresentTimesAfterChange  int       `json:"present_times_after_change" xorm:"comment('赠送次数发生后次数') INT(11)"`
	PresentTimesBeforeChange int       `json:"present_times_before_change" xorm:"comment('赠送次数发生前次数') INT(11)"`
	PresentTimesChange       int       `json:"present_times_change" xorm:"comment('赠送次数发生次数') INT(11)"`
	RecordNumber             string    `json:"record_number" xorm:"comment('流水号') VARCHAR(50)"`
	RefundNumber             string    `json:"refund_number" xorm:"comment('退款流水号') VARCHAR(50)"`
	RmaComment               string    `json:"rma_comment" xorm:"comment('退款备注') VARCHAR(300)"`
	TimesAfterChange         int       `json:"times_after_change" xorm:"comment('本金次数发生后次数') INT(11)"`
	TimesBeforeChange        int       `json:"times_before_change" xorm:"comment('本金次数发生前次数') INT(11)"`
	TimesChange              int       `json:"times_change" xorm:"comment('本金次数发生次数') INT(11)"`
}
