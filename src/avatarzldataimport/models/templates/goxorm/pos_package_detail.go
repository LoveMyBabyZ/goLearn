package goxorm

import (
	"time"
)

type PosPackageDetail struct {
	CategoryDetailCode  string    `json:"category_detail_code" xorm:"not null default '' comment('保障卡类别明细编码') index(category_detail_code) VARCHAR(50)"`
	CategoryDetailName  string    `json:"category_detail_name" xorm:"not null default '' comment('保障卡类别明细名称') index(category_detail_code) VARCHAR(50)"`
	Count               int       `json:"count" xorm:"comment('数量') INT(11)"`
	CreateUserid        int       `json:"create_userid" xorm:"comment('创建人Id') INT(11)"`
	CreateUsername      string    `json:"create_username" xorm:"comment('创建人姓名') VARCHAR(50)"`
	CreatedDate         time.Time `json:"created_date" xorm:"comment('创建时间') DATETIME"`
	DiscountAmount      string    `json:"discount_amount" xorm:"default 0.00 comment('折扣金额') DECIMAL(18,2)"`
	DiscountType        int       `json:"discount_type" xorm:"not null default 0 comment('优惠方式: 1会员打折 2套餐 3打包 4保障卡预售 5保障卡折扣') index(discount_type) INT(11)"`
	Id                  int       `json:"id" xorm:"not null pk autoincr comment('主键') INT(11)"`
	LastUpdateDate      time.Time `json:"last_update_date" xorm:"comment('更新时间') DATETIME"`
	LastUpdateUserid    int       `json:"last_update_userid" xorm:"comment('更新人Id') INT(11)"`
	LastUpdateUsername  string    `json:"last_update_username" xorm:"comment('更新人姓名') VARCHAR(50)"`
	LockInventorySource int       `json:"lock_inventory_source" xorm:"default 0 comment('锁库来源') INT(11)"`
	ObjType             int       `json:"obj_type" xorm:"not null default 0 comment('对象类型(对应DrugType)') INT(11)"`
	OrderDetailId       int       `json:"order_detail_id" xorm:"comment('支付明细Id') INT(11)"`
	OrderId             int       `json:"order_id" xorm:"not null default 0 comment('订单Id') index(org_id) INT(11)"`
	OrderNumber         string    `json:"order_number" xorm:"comment('支付流水号') index(org_id) VARCHAR(50)"`
	OrgId               int       `json:"org_id" xorm:"not null default 0 comment('机构Id') index(category_detail_code) index(discount_type) index(org_id) INT(11)"`
	OriginalPrice       string    `json:"original_price" xorm:"default 0.00 comment('原始价格') DECIMAL(18,2)"`
	OutStoreUnitName    string    `json:"out_store_unit_name" xorm:"comment('出库单位') VARCHAR(50)"`
	PackageId           int       `json:"package_id" xorm:"default 0 comment('包Id') INT(11)"`
	PackageType         int       `json:"package_type" xorm:"default 0 comment('包类型（1:打包 2:套餐）') TINYINT(4)"`
	PpayId              string    `json:"ppay_id" xorm:"comment('明细流水号') index VARCHAR(50)"`
	ProductId           int       `json:"product_id" xorm:"comment('产品Id') INT(11)"`
	ProductName         string    `json:"product_name" xorm:"comment('产品名称') VARCHAR(100)"`
	ProductType         int       `json:"product_type" xorm:"comment('产品类型') INT(11)"`
	SellUnitName        string    `json:"sell_unit_name" xorm:"comment('销售单位') VARCHAR(50)"`
	TotalAmount         string    `json:"total_amount" xorm:"default 0.00 comment('总金额') DECIMAL(18,2)"`
	UnitPrice           string    `json:"unit_price" xorm:"default 0.00 comment('折后单价') DECIMAL(18,2)"`
	UpdateTimestamp     time.Time `json:"update_timestamp" xorm:"not null default 'CURRENT_TIMESTAMP(6)' comment('更新时间戳') TIMESTAMP(6)"`
}
