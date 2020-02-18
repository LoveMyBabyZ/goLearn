package goxorm

import (
	"time"
)

type ErpJxcGoodsList struct {
	Addtime                    time.Time `json:"addtime" xorm:"default 'CURRENT_TIMESTAMP' comment('系统添加时间') index(org_id) DATETIME"`
	AmountPurchaseAfterChange  string    `json:"amount_purchase_after_change" xorm:"default 0.00000000 comment('期末库存采购金额') DECIMAL(18,8)"`
	AmountPurchaseBeforeChange string    `json:"amount_purchase_before_change" xorm:"default 0.00000000 comment('期初库存采购金额') DECIMAL(18,8)"`
	AmountSellAfterChange      string    `json:"amount_sell_after_change" xorm:"default 0.00000000 comment('期末库存销售金额') DECIMAL(18,8)"`
	AmountSellBeforeChange     string    `json:"amount_sell_before_change" xorm:"default 0.00000000 comment('期初库存销售金额') DECIMAL(18,8)"`
	BatchNumber                int       `json:"batch_number" xorm:"comment('批次号') index(report_org) INT(11)"`
	BillNumber                 string    `json:"bill_number" xorm:"default '‘’' comment('单据号') index(org_id) VARCHAR(100)"`
	CategoryNamePath           string    `json:"category_name_path" xorm:"default '‘’' comment('品项目录') VARCHAR(100)"`
	CommonName                 string    `json:"common_name" xorm:"comment('通用名') VARCHAR(100)"`
	CreateUserId               int       `json:"create_user_id" xorm:"comment('制单人id') INT(11)"`
	CreateUserName             string    `json:"create_user_name" xorm:"comment('制单人姓名') VARCHAR(50)"`
	GrossProfit                string    `json:"gross_profit" xorm:"default 0.00000000 comment('销售毛利') DECIMAL(18,8)"`
	Id                         int       `json:"id" xorm:"not null pk autoincr comment('id') INT(11)"`
	InBillType                 int       `json:"in_bill_type" xorm:"comment('入库单类型') index(org_id) INT(11)"`
	InstockOrgId               int       `json:"instock_org_id" xorm:"comment('调拨调入机构id') INT(11)"`
	InstockOrgName             string    `json:"instock_org_name" xorm:"comment('调拨调入机构名称') VARCHAR(100)"`
	ItemId                     int       `json:"item_id" xorm:"default 0 comment('品项Id') index(org_id) INT(11)"`
	ItemName                   string    `json:"item_name" xorm:"default '‘’' comment('品项名称') VARCHAR(100)"`
	ItemPurchaseAmount         string    `json:"item_purchase_amount" xorm:"default 0.00000000 comment('采购金额') DECIMAL(18,8)"`
	ItemPurchasePrice          string    `json:"item_purchase_price" xorm:"default 0.00000000 comment('采购单价') DECIMAL(18,8)"`
	ItemSellAmount             string    `json:"item_sell_amount" xorm:"default 0.00000000 comment('销售金额') DECIMAL(18,8)"`
	ItemSellPrice              string    `json:"item_sell_price" xorm:"default 0.00000000 comment('销售单价') DECIMAL(18,8)"`
	ItemSurplusAmount          string    `json:"item_surplus_amount" xorm:"default 0.00000000 comment('品项结余金额') DECIMAL(18,8)"`
	ItemSurplusNum             string    `json:"item_surplus_num" xorm:"default 0.00 comment('品项结余数量') DECIMAL(18,2)"`
	MemberCardId               int       `json:"member_card_id" xorm:"default 0 comment('会员卡Id') INT(11)"`
	MemberCardNumber           string    `json:"member_card_number" xorm:"default '‘’' comment('会员卡编号') VARCHAR(100)"`
	MemberDiscount             string    `json:"member_discount" xorm:"default 0.00 comment('会员折扣') DECIMAL(18,2)"`
	ObjType                    int       `json:"obj_type" xorm:"default 0 comment('产品类型(小暖)') index(org_id) INT(11)"`
	OccurDate                  time.Time `json:"occur_date" xorm:"default 'CURRENT_TIMESTAMP' comment('发生日期 ') index(org_id) DATETIME"`
	OccurPrice                 string    `json:"occur_price" xorm:"comment('发生单价') DECIMAL(18,8)"`
	OccurQuantity              string    `json:"occur_quantity" xorm:"default 0.00 comment('发生数量') DECIMAL(18,2)"`
	OperateUserId              int       `json:"operate_user_id" xorm:"comment('操作人id') INT(11)"`
	OperateUserName            string    `json:"operate_user_name" xorm:"comment('操作人姓名') VARCHAR(50)"`
	OrgId                      int       `json:"org_id" xorm:"default 0 comment('机构Id') index(org_id) index(report_org) INT(11)"`
	OutBillType                int       `json:"out_bill_type" xorm:"comment('出库单类型') index(org_id) INT(11)"`
	PackSpecific               string    `json:"pack_specific" xorm:"comment('包装规格') VARCHAR(100)"`
	ProductType                int       `json:"product_type" xorm:"default 0 comment('产品类型') index(org_id) INT(11)"`
	ProductTypeName            string    `json:"product_type_name" xorm:"comment('项目类别名称') VARCHAR(50)"`
	QuantityAfterChange        string    `json:"quantity_after_change" xorm:"default 0.00 comment('期末库存数量') DECIMAL(18,2)"`
	QuantityBeforeChange       string    `json:"quantity_before_change" xorm:"default 0.00 comment('期初库存数量') DECIMAL(18,2)"`
	RealOutAmount              string    `json:"real_out_amount" xorm:"default 0.00000000 comment('实出金额') DECIMAL(18,8)"`
	RealOutPrice               string    `json:"real_out_price" xorm:"default 0.00000000 comment('实出单价') DECIMAL(18,8)"`
	Remark                     string    `json:"remark" xorm:"comment('备注') VARCHAR(200)"`
	SellDetailId               int       `json:"sell_detail_id" xorm:"default 0 comment('销售明细id') INT(11)"`
	SellHisDiscount            string    `json:"sell_his_discount" xorm:"default 0.00 comment('his折扣') DECIMAL(18,2)"`
	SellIncome                 string    `json:"sell_income" xorm:"default 0.00000000 comment('销售收入') DECIMAL(18,8)"`
	SellOrderId                int       `json:"sell_order_id" xorm:"default 0 comment('销售单Id') INT(11)"`
	SellOrderNumber            string    `json:"sell_order_number" xorm:"default '‘’' comment('销售单编号') VARCHAR(100)"`
	SellPosDiscount            string    `json:"sell_pos_discount" xorm:"default 0.00 comment('pos折扣') DECIMAL(18,2)"`
	StoreUnitId                int       `json:"store_unit_id" xorm:"comment('入库单位') INT(11)"`
	StoreUnitName              string    `json:"store_unit_name" xorm:"comment('入库单位名称') VARCHAR(50)"`
	Summary                    string    `json:"summary" xorm:"default '‘’' comment('摘要') VARCHAR(100)"`
	SupplierId                 int       `json:"supplier_id" xorm:"default 0 comment('供应商Id') index(org_id) INT(11)"`
	SupplierName               string    `json:"supplier_name" xorm:"default '‘’' comment('供应商名称') index(org_id) VARCHAR(100)"`
}
