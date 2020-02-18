package goxorm

import (
	"time"
)

type ErpInventoryLock struct {
	CreateUserId       int       `json:"create_user_id" xorm:"comment('创建人id(制单人)') INT(11)"`
	CreateUserName     string    `json:"create_user_name" xorm:"comment('创建人姓名(制单人)') VARCHAR(100)"`
	ExpiryDate         time.Time `json:"expiry_date" xorm:"comment('截止日期(有效期限)') DATETIME"`
	Id                 int       `json:"id" xorm:"not null pk autoincr comment('id') INT(11)"`
	ItemId             int       `json:"item_id" xorm:"default 0 comment('品项id') INT(11)"`
	ItemName           string    `json:"item_name" xorm:"default '' comment('品项名称') VARCHAR(100)"`
	ItemPresentPrice   string    `json:"item_present_price" xorm:"default 0.00000000 comment('赠品单价') DECIMAL(18,8)"`
	ItemPrice          string    `json:"item_price" xorm:"default 0.00000000 comment('库存单价') DECIMAL(18,8)"`
	LockType           int       `json:"lock_type" xorm:"default 0 comment('锁库类型：1销售、2调拨') INT(11)"`
	LockTypeDetail     int       `json:"lock_type_detail" xorm:"comment('锁库类型明细id：10、商品开单，20、门诊开单，30、住院开单，101、调拨出库') INT(11)"`
	ObjType            int       `json:"obj_type" xorm:"comment('项目类别(小暖)') INT(11)"`
	OrderDetailId      int       `json:"order_detail_id" xorm:"default 0 comment('销售id') INT(11)"`
	OrderId            int       `json:"order_id" xorm:"default 0 comment('销售单id') INT(11)"`
	OrderNumber        string    `json:"order_number" xorm:"default '' comment('销售单号') VARCHAR(100)"`
	OrgId              int       `json:"org_id" xorm:"default 0 comment('机构id') INT(11)"`
	PackSpecific       string    `json:"pack_specific" xorm:"comment('项目类别名称') VARCHAR(100)"`
	PresentQty         string    `json:"present_qty" xorm:"default 0.00 comment('赠品锁库数量') DECIMAL(18,2)"`
	ProductBatchNumber string    `json:"product_batch_number" xorm:"default '' comment('产品批号') VARCHAR(100)"`
	ProductType        int       `json:"product_type" xorm:"comment('项目类别') INT(11)"`
	ProductTypeName    string    `json:"product_type_name" xorm:"comment('项目类别名称') VARCHAR(50)"`
	ProductionDate     time.Time `json:"production_date" xorm:"default 'CURRENT_TIMESTAMP' comment('生产日期') DATETIME"`
	Qty                string    `json:"qty" xorm:"default 0.00 comment('锁库数量') DECIMAL(18,2)"`
	StockData          time.Time `json:"stock_data" xorm:"default 'CURRENT_TIMESTAMP' comment('锁库时间') DATETIME"`
	StockId            int       `json:"stock_id" xorm:"default 0 comment('库存id') INT(11)"`
	StoreUnitId        int       `json:"store_unit_id" xorm:"comment('入库单位') INT(11)"`
	StoreUnitName      string    `json:"store_unit_name" xorm:"comment('入库单位名称') VARCHAR(50)"`
	SupplierId         int       `json:"supplier_id" xorm:"default 0 comment('供应商id') INT(11)"`
	SupplierName       string    `json:"supplier_name" xorm:"default '' comment('供应商名称') VARCHAR(100)"`
}
