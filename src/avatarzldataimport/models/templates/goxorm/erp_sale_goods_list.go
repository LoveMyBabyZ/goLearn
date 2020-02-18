package goxorm

import (
	"time"
)

type ErpSaleGoodsList struct {
	AddTime                 time.Time `json:"add_time" xorm:"default 'CURRENT_TIMESTAMP' comment('添加时间') DATETIME"`
	BillAmount              string    `json:"bill_amount" xorm:"default 0.00000000 comment('帐单金额') DECIMAL(18,8)"`
	BillDetailId            int       `json:"bill_detail_id" xorm:"default 0 comment('单据明细Id') INT(11)"`
	BillId                  int       `json:"bill_id" xorm:"default 0 comment('单据Id') INT(11)"`
	BillNumber              string    `json:"bill_number" xorm:"default '' comment('单据流水号') VARCHAR(100)"`
	BillType                int       `json:"bill_type" xorm:"default 0 comment('单据类型（1：订单 2：RMA）') INT(11)"`
	BillUserId              int       `json:"bill_user_id" xorm:"comment('开单人Id') INT(11)"`
	BillUserName            string    `json:"bill_user_name" xorm:"comment('开单人名称') VARCHAR(50)"`
	CardCorpus              string    `json:"card_corpus" xorm:"default 0.00000000 comment('储值卡本金') DECIMAL(18,8)"`
	CardPayTime             string    `json:"card_pay_time" xorm:"default 0.00000000 comment('次卡') DECIMAL(18,8)"`
	CardPresent             string    `json:"card_present" xorm:"default 0.00000000 comment('储值卡赠送') DECIMAL(18,8)"`
	Cash                    string    `json:"cash" xorm:"default 0.00000000 comment('现金') DECIMAL(18,8)"`
	CellPhone               string    `json:"cell_phone" xorm:"comment('手机号码') VARCHAR(50)"`
	Coupon                  string    `json:"coupon" xorm:"default 0.00000000 comment('代金券') DECIMAL(18,8)"`
	Deposit                 string    `json:"deposit" xorm:"default 0.00000000 comment('押金') DECIMAL(18,8)"`
	EnsureCards             string    `json:"ensure_cards" xorm:"default 0.00000000 comment('保障卡') DECIMAL(18,8)"`
	HisPrice                string    `json:"his_price" xorm:"comment('HIS价格') DECIMAL(18,8)"`
	HourRange               int       `json:"hour_range" xorm:"default 0 comment('时间段') INT(11)"`
	Id                      int       `json:"id" xorm:"not null pk autoincr INT(11)"`
	Insurance               string    `json:"insurance" xorm:"default 0.00000000 comment('保险') DECIMAL(18,8)"`
	ItemCategoryId          int       `json:"item_category_id" xorm:"default 0 comment('品项目录Id') INT(11)"`
	ItemId                  int       `json:"item_id" xorm:"default 0 comment('品项id') INT(11)"`
	ItemName                string    `json:"item_name" xorm:"default '' comment('品项名称') VARCHAR(100)"`
	MemberCardDiscount      string    `json:"member_card_discount" xorm:"default 0.00 comment('会员折扣') DECIMAL(18,2)"`
	MemberCardId            int       `json:"member_card_id" xorm:"default 0 comment('会员Id') INT(11)"`
	MemberCardName          string    `json:"member_card_name" xorm:"default '' comment('会员名称') VARCHAR(50)"`
	MemberCardNumber        string    `json:"member_card_number" xorm:"default '' comment('会员编号') VARCHAR(100)"`
	MemberCardPrice         string    `json:"member_card_price" xorm:"comment('会员价') DECIMAL(18,8)"`
	ObjType                 int       `json:"obj_type" xorm:"default 0 comment('产品类型(小暖)') INT(11)"`
	OccurDate               time.Time `json:"occur_date" xorm:"default 'CURRENT_TIMESTAMP' comment('发生日期 ') index(org_id) DATETIME"`
	OccurQuantity           string    `json:"occur_quantity" xorm:"default 0.00 comment('发生数量') DECIMAL(18,2)"`
	OrgId                   int       `json:"org_id" xorm:"default 0 comment('机构id') index(org_id) INT(11)"`
	OtherWay                string    `json:"other_way" xorm:"default 0.00000000 comment('其他方式') DECIMAL(18,8)"`
	PayTimeDiscount         string    `json:"pay_time_discount" xorm:"comment('次卡折扣金额') DECIMAL(18,8)"`
	PetId                   int       `json:"pet_id" xorm:"default 0 comment('宠物Id') INT(11)"`
	PetName                 string    `json:"pet_name" xorm:"default '' comment('宠物名称') VARCHAR(50)"`
	Pos                     string    `json:"pos" xorm:"default 0.00000000 comment('POS') DECIMAL(18,8)"`
	ProductType             int       `json:"product_type" xorm:"default 0 comment('产品类型') INT(11)"`
	ProductTypeCategoryId   int       `json:"product_type_category_id" xorm:"default 0 comment('产品所属大类(1:销售库存类 2:医疗产品 3:非医疗服务)') TINYINT(4)"`
	PromotionAmount         string    `json:"promotion_amount" xorm:"comment('促销金额') DECIMAL(18,2)"`
	PromotionNumber         string    `json:"promotion_number" xorm:"comment('促销单号') VARCHAR(100)"`
	PromotionOriginalAmount string    `json:"promotion_original_amount" xorm:"comment('促销原始金额') DECIMAL(18,2)"`
	PromotionSource         string    `json:"promotion_source" xorm:"comment('促销来源') VARCHAR(50)"`
	SellAmount              string    `json:"sell_amount" xorm:"default 0.00000000 comment('销售金额') DECIMAL(18,8)"`
	SellDiscount            string    `json:"sell_discount" xorm:"default 0.00000000 comment('销售总折扣') DECIMAL(18,8)"`
	SellIncome              string    `json:"sell_income" xorm:"default 0.00000000 comment('销售收入（折后总金额）') DECIMAL(18,8)"`
	SellPrice               string    `json:"sell_price" xorm:"default 0.00000000 comment('销售单价') DECIMAL(18,8)"`
	SellUserId              int       `json:"sell_user_id" xorm:"comment('收银人Id') INT(11)"`
	SellUserName            string    `json:"sell_user_name" xorm:"comment('收银人名称') VARCHAR(50)"`
	TeamBuy                 string    `json:"team_buy" xorm:"default 0.00000000 comment('团购') DECIMAL(18,8)"`
	ThirdPartyCard          string    `json:"third_party_card" xorm:"default 0.00000000 comment('第三方卡') DECIMAL(18,8)"`
}
