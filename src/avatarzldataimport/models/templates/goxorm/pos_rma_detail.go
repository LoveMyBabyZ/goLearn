package goxorm

import (
	"time"
)

type PosRmaDetail struct {
	CreateUserid       int       `json:"create_userid" xorm:"comment('创建人Id') INT(11)"`
	CreateUsername     string    `json:"create_username" xorm:"comment('创建人姓名') VARCHAR(50)"`
	CreatedDate        time.Time `json:"created_date" xorm:"comment('创建时间') DATETIME"`
	CustomerId         int       `json:"customer_id" xorm:"comment('会员Id') INT(11)"`
	CustomerName       string    `json:"customer_name" xorm:"comment('会员名称') VARCHAR(50)"`
	Id                 int       `json:"id" xorm:"not null pk autoincr comment('主键') INT(11)"`
	LastUpdateDate     time.Time `json:"last_update_date" xorm:"comment('更新时间') DATETIME"`
	LastUpdateUserid   int       `json:"last_update_userid" xorm:"comment('更新人Id') INT(11)"`
	LastUpdateUsername string    `json:"last_update_username" xorm:"comment('更新人姓名') VARCHAR(50)"`
	ObjType            int       `json:"obj_type" xorm:"not null default 0 comment('对象类型(对应DrugType)') INT(11)"`
	OrderDetailId      int       `json:"order_detail_id" xorm:"comment('支付明细表Id') INT(11)"`
	OrderId            int       `json:"order_id" xorm:"not null default 0 comment('订单Id') INT(11)"`
	OrderNumber        string    `json:"order_number" xorm:"comment('支付流水号') index(org_id) VARCHAR(50)"`
	OrgId              int       `json:"org_id" xorm:"not null default 0 comment('机构Id') index(org_id) index(rma_date) INT(11)"`
	PetId              int       `json:"pet_id" xorm:"comment('宠物Id') INT(11)"`
	PetName            string    `json:"pet_name" xorm:"comment('宠物名称') VARCHAR(50)"`
	PpayId             string    `json:"ppay_id" xorm:"comment('明细流水号') VARCHAR(50)"`
	ProductId          int       `json:"product_id" xorm:"comment('产品Id') INT(11)"`
	ProductName        string    `json:"product_name" xorm:"comment('产品名称') VARCHAR(100)"`
	ProductType        int       `json:"product_type" xorm:"comment('产品类型') INT(11)"`
	Remark             string    `json:"remark" xorm:"comment('备注') VARCHAR(200)"`
	RmaAmount          string    `json:"rma_amount" xorm:"comment('退货金额') DECIMAL(18,2)"`
	RmaCount           int       `json:"rma_count" xorm:"comment('退货数量') INT(11)"`
	RmaDate            time.Time `json:"rma_date" xorm:"comment('退货时间') index(rma_date) DATETIME"`
	RmaDiscountAmount  string    `json:"rma_discount_amount" xorm:"default 0.00 comment('退款折扣金额') DECIMAL(18,2)"`
	RmaId              int       `json:"rma_id" xorm:"not null default 0 comment('退货Id') index(org_id) INT(11)"`
	RmaNumber          string    `json:"rma_number" xorm:"comment('退货流水号') index(org_id) VARCHAR(50)"`
	RmaTotalAmount     string    `json:"rma_total_amount" xorm:"default 0.00 comment('应退金额') DECIMAL(18,2)"`
	UpdateTimestamp    time.Time `json:"update_timestamp" xorm:"not null default 'CURRENT_TIMESTAMP(6)' comment('更新时间戳') TIMESTAMP(6)"`
}
