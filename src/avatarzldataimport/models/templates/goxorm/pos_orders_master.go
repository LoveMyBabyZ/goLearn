package goxorm

import (
	"time"
)

type PosOrdersMaster struct {
	ActulyPayed        string    `json:"actuly_payed" xorm:"comment('实付金额') DECIMAL(18,2)"`
	CreateUserid       int       `json:"create_userid" xorm:"comment('创建人Id') INT(11)"`
	CreateUsername     string    `json:"create_username" xorm:"comment('创建人姓名') VARCHAR(50)"`
	CreatedDate        time.Time `json:"created_date" xorm:"comment('创建时间') DATETIME"`
	CustomerId         int       `json:"customer_id" xorm:"comment('会员Id') index(customer_id) INT(11)"`
	CustomerName       string    `json:"customer_name" xorm:"comment('会员名称') VARCHAR(50)"`
	CustomerNumber     string    `json:"customer_number" xorm:"comment('会员编码') VARCHAR(50)"`
	DiscountAmount     string    `json:"discount_amount" xorm:"comment('优惠金额') DECIMAL(18,2)"`
	Id                 int       `json:"id" xorm:"not null pk autoincr comment('主键') INT(11)"`
	KeepAccount        string    `json:"keep_account" xorm:"comment('记账金额') DECIMAL(18,2)"`
	KeepOrderId        int       `json:"keep_order_id" xorm:"default 0 comment('记账订单Id') INT(11)"`
	LastPrintTime      time.Time `json:"last_print_time" xorm:"comment('最后打印小票时间') DATETIME"`
	LastUpdateDate     time.Time `json:"last_update_date" xorm:"comment('更新时间') DATETIME"`
	LastUpdateUserid   int       `json:"last_update_userid" xorm:"comment('更新人Id') INT(11)"`
	LastUpdateUsername string    `json:"last_update_username" xorm:"comment('更新人姓名') VARCHAR(50)"`
	OrderNumber        string    `json:"order_number" xorm:"comment('支付流水号') VARCHAR(50)"`
	OrderStatus        int       `json:"order_status" xorm:"comment('状态（0:未支付 10:已支付 11:记账支付成功 13:退款中 15:部分退款 20:全部退款）') TINYINT(4)"`
	OrderType          int       `json:"order_type" xorm:"default 0 comment('支付类型（1:正常 2:记账 3:预售）') TINYINT(4)"`
	OrgId              int       `json:"org_id" xorm:"not null default 0 comment('机构Id') index(customer_id) index INT(11)"`
	PayedDate          time.Time `json:"payed_date" xorm:"comment('支付时间') DATETIME"`
	PetId              int       `json:"pet_id" xorm:"comment('宠物Id') INT(11)"`
	PetName            string    `json:"pet_name" xorm:"comment('宠物名称') VARCHAR(50)"`
	PrintSerialNumber  string    `json:"print_serial_number" xorm:"comment('小票打印编号') VARCHAR(50)"`
	TotalAmount        string    `json:"total_amount" xorm:"comment('总金额') DECIMAL(18,2)"`
	UpdateTimestamp    time.Time `json:"update_timestamp" xorm:"not null default 'CURRENT_TIMESTAMP(6)' comment('更新时间戳') TIMESTAMP(6)"`
}
