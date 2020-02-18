package goxorm

import (
	"time"
)

type ErpOutBillDetail struct {
	CategoryNamePath   string    `json:"category_name_path" xorm:"default '' comment('完整的产品目录名称') VARCHAR(200)"`
	CommonName         string    `json:"common_name" xorm:"comment('通用名') VARCHAR(100)"`
	Deviation          string    `json:"deviation" xorm:"default 0.00000000 comment('偏差') DECIMAL(18,8)"`
	ExpiryDate         time.Time `json:"expiry_date" xorm:"default 'CURRENT_TIMESTAMP' comment('截止日期(有效期限)') DATETIME"`
	Id                 int       `json:"id" xorm:"not null pk autoincr comment('id') INT(11)"`
	InBillDetailId     int       `json:"in_bill_detail_id" xorm:"default 0 comment('入库单明细id（供应商退货传）') INT(11)"`
	InventoryId        int       `json:"inventory_id" xorm:"default 0 comment('库存Id(根据此Id锁库存)') INT(11)"`
	ItemId             int       `json:"item_id" xorm:"default 0 comment('品项Id') INT(11)"`
	ItemName           string    `json:"item_name" xorm:"default '' comment('品项名称') VARCHAR(100)"`
	ObjType            int       `json:"obj_type" xorm:"default 0 comment('项目类别(小暖)') INT(11)"`
	OrderDetailId      int       `json:"order_detail_id" xorm:"default 0 comment('销售订单明细id') INT(11)"`
	OrderPackageId     int       `json:"order_package_id" xorm:"default 0 comment('打包明细id') INT(11)"`
	OutAmount          string    `json:"out_amount" xorm:"default 0.00000000 comment('出库总金额（包含赠品）') DECIMAL(18,8)"`
	OutBillId          int       `json:"out_bill_id" xorm:"default 0 comment('出库单id') index INT(11)"`
	OutPlanDetailId    int       `json:"out_plan_detail_id" xorm:"default 0 comment('出库计划单明细Id') INT(11)"`
	OutQuantity        string    `json:"out_quantity" xorm:"default 0.00 comment('出库总数量') DECIMAL(18,2)"`
	PackSpecific       string    `json:"pack_specific" xorm:"default '' comment('包装规格') VARCHAR(100)"`
	ProductBatchNumber string    `json:"product_batch_number" xorm:"default '' comment('产品批号') VARCHAR(100)"`
	ProductType        int       `json:"product_type" xorm:"default 0 comment('项目类别') INT(11)"`
	ProductTypeName    string    `json:"product_type_name" xorm:"default '' comment('项目类别名称') VARCHAR(50)"`
	ProductionDate     time.Time `json:"production_date" xorm:"default 'CURRENT_TIMESTAMP' comment('生产日期') DATETIME"`
	RealOutAmount      string    `json:"real_out_amount" xorm:"default 0.00000000 comment('实出金额') DECIMAL(18,8)"`
	RealOutPrice       string    `json:"real_out_price" xorm:"default 0.00000000 comment('实出单价') DECIMAL(18,8)"`
	Remark             string    `json:"remark" xorm:"default '' comment('备注') VARCHAR(500)"`
	SellReturnQuantity string    `json:"sell_return_quantity" xorm:"default 0.00 comment('已退数量') DECIMAL(18,2)"`
	StoreUnitId        int       `json:"store_unit_id" xorm:"default 0 comment('库存单位') INT(11)"`
	StoreUnitName      string    `json:"store_unit_name" xorm:"default '' comment('库存单位名称') VARCHAR(50)"`
	UpdateTimestamp    time.Time `json:"update_timestamp" xorm:"not null default 'CURRENT_TIMESTAMP(6)' comment('更新时间戳') TIMESTAMP(6)"`
}
