package goxorm

type ErpOutBillExtReturn struct {
	Id              int    `json:"id" xorm:"not null pk autoincr comment('id') INT(11)"`
	InBillId        int    `json:"in_bill_id" xorm:"default 0 comment('对应的采购入库单Id') INT(11)"`
	InBillNumber    string `json:"in_bill_number" xorm:"default '' comment('对应的采购入库单号') VARCHAR(50)"`
	OutBillId       int    `json:"out_bill_id" xorm:"default 0 comment('出库单id') index INT(11)"`
	SourceId        int    `json:"source_id" xorm:"default 0 comment('1: 自建采购单  2：集采采购单') INT(11)"`
	SupplierAddress string `json:"supplier_address" xorm:"default '' comment('供应商地址') VARCHAR(100)"`
	SupplierContact string `json:"supplier_contact" xorm:"default '' comment('供应商联系人') VARCHAR(50)"`
	SupplierId      int    `json:"supplier_id" xorm:"default 0 comment('退货供应商Id') INT(11)"`
	SupplierName    string `json:"supplier_name" xorm:"default '' comment('供应商名称') VARCHAR(100)"`
	SupplierPhone   string `json:"supplier_phone" xorm:"default '' comment('供应商电话') VARCHAR(50)"`
}
