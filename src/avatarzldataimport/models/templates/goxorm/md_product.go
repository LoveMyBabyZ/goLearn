package goxorm

import (
	"time"
)

type MdProduct struct {
	AddTime                 time.Time `json:"add_time" xorm:"default 'CURRENT_TIMESTAMP' comment('添加时间') DATETIME"`
	ApplicablePet           string    `json:"applicable_pet" xorm:"not null default '' comment('适用宠物') VARCHAR(255)"`
	ApplicablePetName       string    `json:"applicable_pet_name" xorm:"not null default '' comment('适用宠物名称') VARCHAR(500)"`
	ApprovalSn              string    `json:"approval_sn" xorm:"not null default '' comment('批准文号--总部字段') VARCHAR(200)"`
	BarCode                 string    `json:"bar_code" xorm:"not null default '' comment('条形码') VARCHAR(20)"`
	BrandId                 string    `json:"brand_id" xorm:"not null default '' comment('品牌id') VARCHAR(100)"`
	BrandName               string    `json:"brand_name" xorm:"not null default '' comment('品牌名称') VARCHAR(200)"`
	CanNegativeStock        int       `json:"can_negative_stock" xorm:"not null default 0 comment('是否允许零库存销售') TINYINT(2)"`
	CanOrder                int       `json:"can_order" xorm:"not null default 1 comment('可订') TINYINT(2)"`
	CanSell                 int       `json:"can_sell" xorm:"not null default 1 comment('可销') TINYINT(2)"`
	Cate                    int       `json:"cate" xorm:"not null default 0 comment('所属分类--医疗服务类产品字段') INT(11)"`
	CategoryId              int       `json:"category_id" xorm:"not null default 0 comment('目录编码') index(idx_org_ptype_cate) INT(10)"`
	CategoryName            string    `json:"category_name" xorm:"not null default '' comment('目录名称') VARCHAR(100)"`
	CommonName              string    `json:"common_name" xorm:"not null default '' comment('通用名') VARCHAR(100)"`
	CostPrice               string    `json:"cost_price" xorm:"not null default 0.00 comment('成本价') DECIMAL(30,2)"`
	Disabled                int       `json:"disabled" xorm:"not null default 0 comment('停用') TINYINT(2)"`
	DosageByWeightIncrement int       `json:"dosage_by_weight_increment" xorm:"not null default 0 comment('是否按体重递增设置单次用量') TINYINT(2)"`
	DosageByWeightRange     int       `json:"dosage_by_weight_range" xorm:"not null default 0 comment('是否按体重范围设置单次用量') TINYINT(2)"`
	DosageIncreLevel        string    `json:"dosage_incre_level" xorm:"not null default 0.00 comment('标准差(体重每增加weight_incre_level，投药量增加的百分比)') DECIMAL(19,2)"`
	DosingUnitId            int       `json:"dosing_unit_id" xorm:"not null default 0 comment('投药单位ID') INT(11)"`
	DosingUnitName          string    `json:"dosing_unit_name" xorm:"not null default '' comment('投药单位名称') VARCHAR(50)"`
	DosingWayId             int       `json:"dosing_way_id" xorm:"not null default 0 comment('投药方式id') INT(11)"`
	DosingWayName           string    `json:"dosing_way_name" xorm:"not null default '' comment('投药方式名称') VARCHAR(50)"`
	EnglishName             string    `json:"english_name" xorm:"not null default '' comment('英文名') VARCHAR(200)"`
	From                    int       `json:"from" xorm:"not null default 1 comment('数据来源（1：医院自建；2：数据导入）') TINYINT(1)"`
	HandleLevel             int       `json:"handle_level" xorm:"not null default 1 comment('处置等级') TINYINT(2)"`
	Id                      int       `json:"id" xorm:"not null pk autoincr INT(11)"`
	Img                     string    `json:"img" xorm:"not null default '' comment('产品图片--需要确定小暖端最大长度') VARCHAR(300)"`
	Indication              string    `json:"indication" xorm:"not null default '' comment('适应症') VARCHAR(2000)"`
	Ingredient              string    `json:"ingredient" xorm:"not null default '' comment('药品成分') VARCHAR(200)"`
	IsDeleted               int       `json:"is_deleted" xorm:"not null default 0 comment('是否删除（0：未删除；1：已删除）') TINYINT(1)"`
	IsDiscount              int       `json:"is_discount" xorm:"not null default 1 comment('是否打折：0-否 1-是') TINYINT(2)"`
	IsHistory               int       `json:"is_history" xorm:"not null default 0 comment('是否历史数据（0：非历史数据；1：历史数据）') TINYINT(1)"`
	IsPriceControl          int       `json:"is_price_control" xorm:"not null default 0 comment('是否价格管控') TINYINT(1)"`
	IsRecommend             int       `json:"is_recommend" xorm:"not null default 0 comment('是否推荐用药') TINYINT(1)"`
	IsStandard              int       `json:"is_standard" xorm:"not null default 0 comment('是否标准产品') TINYINT(2)"`
	IsSync                  int       `json:"is_sync" xorm:"not null default 0 comment('是否总部下发/是否下发管控') TINYINT(1)"`
	IsTemp                  int       `json:"is_temp" xorm:"not null default 0 comment('是否临时采购') TINYINT(1)"`
	IsValidityControl       int       `json:"is_validity_control" xorm:"not null default 0 comment('是否效期管理--总部字段') TINYINT(2)"`
	ItemCode                string    `json:"item_code" xorm:"not null default '' comment('总部编码') VARCHAR(50)"`
	LastModify              time.Time `json:"last_modify" xorm:"default 'CURRENT_TIMESTAMP' comment('修改时间') DATETIME"`
	LocalCurrencyId         int       `json:"local_currency_id" xorm:"not null default 0 comment('本位币id--总部字段') INT(11)"`
	LocalCurrencyName       string    `json:"local_currency_name" xorm:"not null default '' comment('本位币名称--总部字段') VARCHAR(50)"`
	MemberPrice             string    `json:"member_price" xorm:"not null default 0.00 comment('会员价') DECIMAL(30,2)"`
	NonSensitive            int       `json:"non_sensitive" xorm:"not null default 1 comment('非敏感药') TINYINT(1)"`
	NotBanned               int       `json:"not_banned" xorm:"not null default 1 comment('是否非禁') TINYINT(1)"`
	NotControl              int       `json:"not_control" xorm:"not null default 1 comment('非管控') TINYINT(1)"`
	NotReusable             int       `json:"not_reusable" xorm:"not null default 1 comment('非复用') TINYINT(1)"`
	NuanId                  int       `json:"nuan_id" xorm:"not null comment('小暖ID') unique(unique_orgId_nuanId) INT(11)"`
	ObjType                 int       `json:"obj_type" xorm:"not null default 0 comment('对象类型') index(index_orgId_objType) INT(11)"`
	OrgId                   int       `json:"org_id" xorm:"not null comment('医院ID') index(idx_org_ptype_cate) index(index_orgId_objType) unique(unique_orgId_nuanId) INT(11)"`
	PackSpecific            string    `json:"pack_specific" xorm:"not null default '' comment('包装规格') VARCHAR(200)"`
	PinyinAbbr              string    `json:"pinyin_abbr" xorm:"not null default '' comment('拼音简写') VARCHAR(100)"`
	ProducerCountry         string    `json:"producer_country" xorm:"not null default '' comment('生产国和地区名称--总部字段') VARCHAR(50)"`
	ProducerCountryId       int       `json:"producer_country_id" xorm:"not null default 0 comment('生产国和地区id--总部字段') INT(11)"`
	ProducerId              string    `json:"producer_id" xorm:"not null default '' comment('生产商ID') VARCHAR(100)"`
	ProducerInnerCode       string    `json:"producer_inner_code" xorm:"not null default '' comment('生产商内部编码--总部字段') VARCHAR(200)"`
	ProducerName            string    `json:"producer_name" xorm:"not null default '' comment('生产商名称') VARCHAR(200)"`
	ProductName             string    `json:"product_name" xorm:"not null default '' comment('产品名称') VARCHAR(200)"`
	ProductType             int       `json:"product_type" xorm:"not null comment('产品类型') index(idx_org_ptype_cate) INT(11)"`
	RefPrice                string    `json:"ref_price" xorm:"not null default 0.00 comment('入库含税参考价') DECIMAL(30,2)"`
	RelType                 int       `json:"rel_type" xorm:"not null default 0 comment('关联类型') TINYINT(1)"`
	Remark                  string    `json:"remark" xorm:"not null default '' comment('备注／注意事项') VARCHAR(2000)"`
	SellPrice               string    `json:"sell_price" xorm:"not null default 0.00 comment('销售单价') DECIMAL(30,2)"`
	SellPriceHigh           string    `json:"sell_price_high" xorm:"not null default '' comment('销售价格区间高价') VARCHAR(30)"`
	SellPriceLow            string    `json:"sell_price_low" xorm:"not null default '' comment('销售价格区间低价') VARCHAR(30)"`
	ServiceScope            int       `json:"service_scope" xorm:"not null default 0 comment('服务范围--医疗服务类产品字段') INT(11)"`
	ServiceStatus           int       `json:"service_status" xorm:"not null default 1 comment('服务状态 启用 1 停用 0') TINYINT(1)"`
	SpecificConversion      string    `json:"specific_conversion" xorm:"not null default 0.00000 comment('规格换算') DECIMAL(19,5)"`
	StandardDosage          string    `json:"standard_dosage" xorm:"not null default 0.00 comment('标准体重投药量') DECIMAL(19,2)"`
	StandardWeight          string    `json:"standard_weight" xorm:"not null default 0.00 comment('标准体重') DECIMAL(19,2)"`
	StoreCondition          string    `json:"store_condition" xorm:"not null default '' comment('保存条件') VARCHAR(500)"`
	StoreUnitId             int       `json:"store_unit_id" xorm:"not null default 0 comment('出入库单位ID') INT(11)"`
	StoreUnitName           string    `json:"store_unit_name" xorm:"not null default '' comment('出入库单位名称') VARCHAR(50)"`
	Tax                     string    `json:"tax" xorm:"not null default 0.00 comment('税额') DECIMAL(30,2)"`
	TaxRate                 string    `json:"tax_rate" xorm:"not null default 0.00 comment('税率') DECIMAL(30,2)"`
	Traits                  string    `json:"traits" xorm:"not null default '' comment('性状') VARCHAR(255)"`
	TraitsName              string    `json:"traits_name" xorm:"not null default '' comment('性状名称') VARCHAR(500)"`
	UntaxedPrice            string    `json:"untaxed_price" xorm:"not null default 0.00 comment('未税单价') DECIMAL(30,2)"`
	UseWayId                int       `json:"use_way_id" xorm:"not null default 0 comment('使用方式id') INT(11)"`
	UseWayName              string    `json:"use_way_name" xorm:"not null default '' comment('使用方式名称') VARCHAR(50)"`
	ValidityDay             int       `json:"validity_day" xorm:"not null default 0 comment('效期:天') INT(11)"`
	WeightIncreLevel        string    `json:"weight_incre_level" xorm:"not null default 0.00 comment('体重增量(体重每增加weight_incre_level，投药量增加dosage_incre_level)') DECIMAL(19,2)"`
}
