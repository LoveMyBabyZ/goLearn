package goxorm

import (
	"time"
)

type PosOrdersDetail struct {
	ActlyPayed            string    `json:"actly_payed" xorm:"comment('实际支付金额') DECIMAL(18,2)"`
	AssistantId           int       `json:"assistant_id" xorm:"not null default 0 comment('助理Id') INT(11)"`
	AssistantName         string    `json:"assistant_name" xorm:"comment('助理姓名') VARCHAR(50)"`
	BusinessId            int       `json:"business_id" xorm:"default 0 comment('单据Id') INT(11)"`
	BusinessNumber        string    `json:"business_number" xorm:"comment('单据流水号') VARCHAR(50)"`
	BusinessType          int       `json:"business_type" xorm:"default 0 comment('单据类型（1:销售单 2:美容单）') TINYINT(4)"`
	CancelDate            time.Time `json:"cancel_date" xorm:"comment('取消时间') DATETIME"`
	CategoryDetailCode    string    `json:"category_detail_code" xorm:"not null default '' comment('保障卡类别明细编码') index(category_detail_code) VARCHAR(50)"`
	CategoryDetailName    string    `json:"category_detail_name" xorm:"not null default '' comment('保障卡类别明细名称') index(category_detail_code) VARCHAR(50)"`
	CategoryId            int       `json:"category_id" xorm:"comment('目录编码') INT(11)"`
	ChargeDatetime        time.Time `json:"charge_datetime" xorm:"comment('下单时间') DATETIME"`
	Count                 int       `json:"count" xorm:"comment('数量') INT(11)"`
	CreateUserid          int       `json:"create_userid" xorm:"comment('创建人Id') INT(11)"`
	CreateUsername        string    `json:"create_username" xorm:"comment('创建人姓名') VARCHAR(50)"`
	CreatedDate           time.Time `json:"created_date" xorm:"comment('创建时间') DATETIME"`
	CustomerId            int       `json:"customer_id" xorm:"comment('会员Id') index(customer_id) index(orgid) INT(11)"`
	CustomerName          string    `json:"customer_name" xorm:"comment('会员名称') VARCHAR(50)"`
	DataChannel           int       `json:"data_channel" xorm:"comment('销售渠道（4:住院）') INT(11)"`
	DiscountAmount        string    `json:"discount_amount" xorm:"comment('折扣金额') DECIMAL(18,2)"`
	DiscountRate          string    `json:"discount_rate" xorm:"comment('折扣率') DECIMAL(18,2)"`
	DiscountType          int       `json:"discount_type" xorm:"not null default 0 comment('优惠方式: 1会员打折 2套餐 3打包 4保障卡预售 5保障卡折扣') index(discount_type) INT(11)"`
	Edited                int       `json:"edited" xorm:"comment('是否修改过价格（0:未改 1:已改）') TINYINT(4)"`
	Id                    int       `json:"id" xorm:"not null pk autoincr comment('主键') INT(11)"`
	InpatientId           int       `json:"inpatient_id" xorm:"not null default 0 comment('住院Id') INT(11)"`
	InpatientState        int       `json:"inpatient_state" xorm:"comment('住院状态') INT(11)"`
	IsDelete              int       `json:"is_delete" xorm:"comment('是否取消（0:正常 1:已取消）') index(orgid) TINYINT(4)"`
	LastUpdateDate        time.Time `json:"last_update_date" xorm:"comment('更新时间') DATETIME"`
	LastUpdateUserid      int       `json:"last_update_userid" xorm:"comment('更新人Id') INT(11)"`
	LastUpdateUsername    string    `json:"last_update_username" xorm:"comment('更新人姓名') VARCHAR(50)"`
	LockInventorySource   int       `json:"lock_inventory_source" xorm:"default 0 comment('锁库来源') INT(11)"`
	ObjType               int       `json:"obj_type" xorm:"not null default 0 comment('对象类型(对应DrugType)') INT(11)"`
	OrderGuid             string    `json:"order_guid" xorm:"comment('订单Guid') VARCHAR(50)"`
	OrderId               int       `json:"order_id" xorm:"not null default 0 comment('订单Id') index index(orgid) INT(11)"`
	OrderNumber           string    `json:"order_number" xorm:"comment('支付流水号') VARCHAR(50)"`
	OrderTime             time.Time `json:"order_time" xorm:"not null comment('处方下单时间') DATETIME"`
	OrgId                 int       `json:"org_id" xorm:"not null default 0 comment('机构Id') index(category_detail_code) index(customer_id) index(discount_type) index(orgid) index(payed_date) index(pet_id) INT(11)"`
	OriPhysicianId        int       `json:"ori_physician_id" xorm:"not null default 0 comment('挂号指定医生Id') INT(11)"`
	OriPhysicianName      string    `json:"ori_physician_name" xorm:"not null default '' comment('挂号指定医生姓名') VARCHAR(50)"`
	OriginalPrice         string    `json:"original_price" xorm:"comment('原始价格') DECIMAL(18,2)"`
	OutStoreUnitName      string    `json:"out_store_unit_name" xorm:"comment('出库单位') VARCHAR(50)"`
	PackageId             int       `json:"package_id" xorm:"default 0 comment('包Id') INT(11)"`
	PackageType           int       `json:"package_type" xorm:"default 0 comment('包类型（1:打包 2:套餐）') TINYINT(4)"`
	PayedDate             time.Time `json:"payed_date" xorm:"comment('支付时间') index(payed_date) DATETIME"`
	PetId                 int       `json:"pet_id" xorm:"comment('宠物Id') index(orgid) index(pet_id) INT(11)"`
	PetName               string    `json:"pet_name" xorm:"comment('宠物名称') VARCHAR(50)"`
	PpayId                string    `json:"ppay_id" xorm:"comment('明细流水号') index(orgid) index VARCHAR(50)"`
	PrescriptionId        int       `json:"prescription_id" xorm:"not null default 0 comment('处方Id') index INT(11)"`
	ProductId             int       `json:"product_id" xorm:"comment('产品Id') INT(11)"`
	ProductName           string    `json:"product_name" xorm:"comment('产品名称') VARCHAR(100)"`
	ProductType           int       `json:"product_type" xorm:"comment('产品类型') INT(11)"`
	Remark                string    `json:"remark" xorm:"comment('备注') VARCHAR(200)"`
	SellToOutConvertValue string    `json:"sell_to_out_convert_value" xorm:"comment('转换因子') DECIMAL(18,2)"`
	SellUnitName          string    `json:"sell_unit_name" xorm:"comment('销售单位') VARCHAR(50)"`
	SellerId              int       `json:"seller_id" xorm:"not null default 0 comment('销售人员Id') INT(11)"`
	SellerName            string    `json:"seller_name" xorm:"comment('销售人员姓名') VARCHAR(50)"`
	Specifications        string    `json:"specifications" xorm:"comment('包装规格') VARCHAR(50)"`
	TotalAmount           string    `json:"total_amount" xorm:"comment('总金额') DECIMAL(18,2)"`
	UnitPrice             string    `json:"unit_price" xorm:"comment('单价') DECIMAL(18,2)"`
	UpdateTimestamp       time.Time `json:"update_timestamp" xorm:"not null default 'CURRENT_TIMESTAMP(6)' comment('更新时间戳') TIMESTAMP(6)"`
}
