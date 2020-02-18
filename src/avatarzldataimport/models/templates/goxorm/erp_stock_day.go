package goxorm

import (
	"time"
)

type ErpStockDay struct {
	Addtime          time.Time `json:"addtime" xorm:"default 'CURRENT_TIMESTAMP' comment('添加时间') index(org_id) DATETIME"`
	CategoryNamePath string    `json:"category_name_path" xorm:"comment('完整的产品目录名称') VARCHAR(200)"`
	Id               int       `json:"id" xorm:"not null pk autoincr comment('id') INT(11)"`
	ItemId           int       `json:"item_id" xorm:"default 0 comment('产品id') index(org_id) INT(11)"`
	ItemName         string    `json:"item_name" xorm:"comment('品项名称') VARCHAR(100)"`
	OrgId            int       `json:"org_id" xorm:"default 0 comment('机构Id') index(org_id) INT(11)"`
	PackSpecific     string    `json:"pack_specific" xorm:"comment('包装规格') VARCHAR(100)"`
	ProductTypeName  string    `json:"product_type_name" xorm:"comment('项目类别名称') VARCHAR(50)"`
	StockAmount      string    `json:"stock_amount" xorm:"default 0.00000000 comment('库存金额') DECIMAL(18,8)"`
	StockDate        time.Time `json:"stock_date" xorm:"default 'CURRENT_TIMESTAMP' comment('库存日期') index(org_id) DATETIME"`
	StockQuantity    string    `json:"stock_quantity" xorm:"default 0.00 comment('库存数量') DECIMAL(18,2)"`
	StoreUnitName    string    `json:"store_unit_name" xorm:"comment('单位名称') VARCHAR(50)"`
	SupplierId       int       `json:"supplier_id" xorm:"default 0 comment('供应商Id') INT(11)"`
}
