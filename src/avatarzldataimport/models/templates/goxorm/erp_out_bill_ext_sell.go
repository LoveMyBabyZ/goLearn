package goxorm

type ErpOutBillExtSell struct {
	Id         int `json:"id" xorm:"not null pk autoincr comment('id') INT(11)"`
	OutBillId  int `json:"out_bill_id" xorm:"default 0 comment('出库单id') INT(11)"`
	SellBillId int `json:"sell_bill_id" xorm:"default 0 comment('销售单Id') INT(11)"`
}
