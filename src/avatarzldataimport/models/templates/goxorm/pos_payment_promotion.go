package goxorm

import (
	"time"
)

type PosPaymentPromotion struct {
	CreateDate              time.Time `json:"create_date" xorm:"comment('创建时间') DATETIME"`
	CreateUserid            int       `json:"create_userid" xorm:"comment('创建人Id') INT(11)"`
	CreateUsername          string    `json:"create_username" xorm:"comment('创建人姓名') VARCHAR(50)"`
	CustomerId              int       `json:"customer_id" xorm:"default 0 comment('会员Id') INT(11)"`
	CustomerName            string    `json:"customer_name" xorm:"comment('会员名称') VARCHAR(50)"`
	Id                      int       `json:"id" xorm:"not null pk autoincr comment('主键') INT(11)"`
	LastUpdateDate          time.Time `json:"last_update_date" xorm:"comment('更新时间') DATETIME"`
	OperatorId              int       `json:"operator_id" xorm:"comment('操作人Id') INT(11)"`
	OperatorName            string    `json:"operator_name" xorm:"comment('操作人姓名') VARCHAR(50)"`
	OrderId                 int       `json:"order_id" xorm:"not null default 0 comment('订单表Id') INT(11)"`
	OrderNumber             string    `json:"order_number" xorm:"comment('订单流水号') VARCHAR(50)"`
	OrgId                   int       `json:"org_id" xorm:"not null default 0 comment('机构Id') INT(11)"`
	PaymentMethodId         int       `json:"payment_method_id" xorm:"default 0 comment('支付表Id') INT(11)"`
	PaymentType             int       `json:"payment_type" xorm:"default 0 comment('支付类型') TINYINT(4)"`
	ProductTypeCategoryId   int       `json:"product_type_category_id" xorm:"default 0 comment('产品所属大类(1:销售库存类 2:医疗产品 3:非医疗服务)') TINYINT(4)"`
	PromotionAmount         string    `json:"promotion_amount" xorm:"comment('促销金额') DECIMAL(18,2)"`
	PromotionNumber         string    `json:"promotion_number" xorm:"comment('促销单号') VARCHAR(100)"`
	PromotionOriginalAmount string    `json:"promotion_original_amount" xorm:"comment('促销原始金额') DECIMAL(18,2)"`
	PromotionSource         string    `json:"promotion_source" xorm:"comment('促销来源') VARCHAR(50)"`
	PromotionSourceId       int       `json:"promotion_source_id" xorm:"comment('促销来源Id') INT(11)"`
	Remark                  string    `json:"remark" xorm:"comment('备注') VARCHAR(200)"`
	UpdateTimestamp         time.Time `json:"update_timestamp" xorm:"not null default 'CURRENT_TIMESTAMP(6)' comment('更新时间戳') TIMESTAMP(6)"`
}
