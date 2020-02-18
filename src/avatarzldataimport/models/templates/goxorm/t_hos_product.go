package goxorm

import (
	"time"
)

type THosProduct struct {
	AddTime                time.Time `json:"add_time" xorm:"not null default 'CURRENT_TIMESTAMP' comment('添加时间') TIMESTAMP"`
	ApplicablePet          string    `json:"applicable_pet" xorm:"not null default '' comment('适用宠物') VARCHAR(255)"`
	ApprovalSn             string    `json:"approval_sn" xorm:"not null default '0' comment('批准文号') VARCHAR(200)"`
	BarCode                string    `json:"bar_code" xorm:"not null default '' comment('条形码') VARCHAR(100)"`
	BrandCode              string    `json:"brand_code" xorm:"not null default '' comment('品牌编码') VARCHAR(20)"`
	BrandName              string    `json:"brand_name" xorm:"not null default '' comment('品牌名称') VARCHAR(50)"`
	CanOrder               int       `json:"can_order" xorm:"not null default 1 comment('是否可预定，1是，0否') TINYINT(3)"`
	CanSell                int       `json:"can_sell" xorm:"not null default 1 comment('是否可sell，1是，0否') TINYINT(3)"`
	Category               string    `json:"category" xorm:"not null default '' comment('目录编码') VARCHAR(50)"`
	ChangeCause            string    `json:"change_cause" xorm:"comment('修改原因通知') VARCHAR(200)"`
	ChargeUnit             string    `json:"charge_unit" xorm:"not null default '0' comment('收费单位ID') VARCHAR(11)"`
	ChargeUnitConversion   string    `json:"charge_unit_conversion" xorm:"not null default '' comment('收费单位换算成出入库单位') VARCHAR(255)"`
	Color                  string    `json:"color" xorm:"not null default '' comment('色彩') VARCHAR(255)"`
	CommonName             string    `json:"common_name" xorm:"not null default '' comment('通用名') VARCHAR(100)"`
	CompanyId              int       `json:"company_id" xorm:"not null default 0 comment('公司ID') unique(uniq_idx_pcode) INT(11)"`
	CountryCode            string    `json:"country_code" xorm:"not null default '' comment('生产国编码') VARCHAR(30)"`
	Currency               int       `json:"currency" xorm:"not null default 0 comment('本位币ID') INT(11)"`
	DeliverySite           string    `json:"delivery_site" xorm:"not null default '' comment('发货地') VARCHAR(255)"`
	DesignLink             string    `json:"design_link" xorm:"not null default '' comment('设计链接') VARCHAR(255)"`
	Disabled               int       `json:"disabled" xorm:"not null default 0 comment('是否停用，1是，0否') TINYINT(3)"`
	DosageType             string    `json:"dosage_type" xorm:"not null default '0' comment('剂型ID') VARCHAR(11)"`
	DosingUnit             string    `json:"dosing_unit" xorm:"not null default '0' comment('投药单位ID') VARCHAR(11)"`
	EnglishName            string    `json:"english_name" xorm:"not null default '' comment('英文名') VARCHAR(200)"`
	ForbidNegativeStock    int       `json:"forbid_negative_stock" xorm:"not null default 0 comment('是否禁止负库存销售：0-否 1-是') TINYINT(2)"`
	From                   int       `json:"from" xorm:"not null default 1 comment('数据来源（1：总部自建；2：数据导入；3：医院申请新增）') TINYINT(1)"`
	Id                     int       `json:"id" xorm:"not null pk autoincr INT(10)"`
	Img                    string    `json:"img" xorm:"not null default '' comment('产品图片地址') VARCHAR(255)"`
	Indication             string    `json:"indication" xorm:"not null default '' comment('适应症') VARCHAR(2000)"`
	Ingredient             string    `json:"ingredient" xorm:"not null default '' comment('药品成分') VARCHAR(1000)"`
	InstructionsValidity   int       `json:"instructions_validity" xorm:"default 0 comment('有效期') INT(11)"`
	IsDiscount             int       `json:"is_discount" xorm:"not null default 1 comment('是否打折：0-否 1-是') TINYINT(2)"`
	IsHq                   int       `json:"is_hq" xorm:"not null default 0 comment('是否总部人员新建 1-总部人员新建   0-区域人员新建') TINYINT(2)"`
	IsRecommend            int       `json:"is_recommend" xorm:"not null default 0 comment('是否推荐用药：0-否 1-是') TINYINT(2)"`
	IsReuse                int       `json:"is_reuse" xorm:"not null default 1 comment('是非复用') TINYINT(2)"`
	IsStandard             int       `json:"is_standard" xorm:"not null default 0 comment('是否标准药品：0-非标 1-标准药品') TINYINT(2)"`
	IsTemp                 int       `json:"is_temp" xorm:"not null default 0 comment('是否临时采购：0-否 1-是') TINYINT(2)"`
	IsValidityControl      int       `json:"is_validity_control" xorm:"not null default 0 comment('是否效期管理：0-否 1-是') TINYINT(2)"`
	LastModify             time.Time `json:"last_modify" xorm:"not null default 'CURRENT_TIMESTAMP' comment('更新时间') TIMESTAMP"`
	Material               string    `json:"material" xorm:"not null default '' comment('材质用料') VARCHAR(255)"`
	MinCount               string    `json:"min_count" xorm:"not null default '' comment('最小收费单位') VARCHAR(100)"`
	MinPackageUnit         string    `json:"min_package_unit" xorm:"not null default '' comment('最小包装单位') VARCHAR(20)"`
	NameRemark             string    `json:"name_remark" xorm:"not null default '' comment('名称标记') VARCHAR(10)"`
	NonSensitive           int       `json:"non_sensitive" xorm:"not null default 1 comment('非敏感') TINYINT(1)"`
	NotBanned              int       `json:"not_banned" xorm:"not null default 0 comment('是否非禁药品：0-否 1-是') TINYINT(2)"`
	NotControl             int       `json:"not_control" xorm:"not null default 0 comment('非管控：0-否 1-是') TINYINT(2)"`
	ObjType                int       `json:"obj_type" xorm:"not null default 0 comment('对象类型') INT(11)"`
	PackSpecific           string    `json:"pack_specific" xorm:"not null default '' comment('包装规格') VARCHAR(100)"`
	PinyinAbbr             string    `json:"pinyin_abbr" xorm:"not null default '' comment('拼音简写') VARCHAR(200)"`
	ProductCode            string    `json:"product_code" xorm:"not null comment('产品编码') unique(uniq_idx_pcode) VARCHAR(20)"`
	ProductName            string    `json:"product_name" xorm:"not null comment('产品名称') VARCHAR(255)"`
	ProductSite            string    `json:"product_site" xorm:"not null default '' comment('生产地') VARCHAR(255)"`
	ProductType            string    `json:"product_type" xorm:"not null default '1046' comment('产品类型') VARCHAR(10)"`
	PurchasePackUnit       string    `json:"purchase_pack_unit" xorm:"not null default '' comment('采购包装规格') VARCHAR(255)"`
	PurchaseUnitConversion string    `json:"purchase_unit_conversion" xorm:"not null default '' comment('采购单位换算成采购包装规格') VARCHAR(255)"`
	PurchasingUnit         string    `json:"purchasing_unit" xorm:"not null default '0' comment('采购单位ID') VARCHAR(11)"`
	RefPrice               string    `json:"ref_price" xorm:"not null default 0.00 comment('入库含税参考价') DECIMAL(30,2)"`
	Remark                 string    `json:"remark" xorm:"not null default '' comment('备注／注意事项') VARCHAR(2000)"`
	SellPrice              string    `json:"sell_price" xorm:"not null default 0.00 comment('销售单价') DECIMAL(30,2)"`
	SellPriceHigh          string    `json:"sell_price_high" xorm:"not null default '' comment('销售价格区间高价') VARCHAR(30)"`
	SellPriceLow           string    `json:"sell_price_low" xorm:"not null default '' comment('销售价格区间低价') VARCHAR(30)"`
	SpecificConversion     string    `json:"specific_conversion" xorm:"not null default '' comment('出入库单位换算成投药单位') VARCHAR(255)"`
	SpecificationDetail    string    `json:"specification_detail" xorm:"not null default '' comment('规格详细说明') VARCHAR(255)"`
	SpecificationSize      string    `json:"specification_size" xorm:"not null default '' comment('规格尺寸') VARCHAR(255)"`
	StoreCondition         string    `json:"store_condition" xorm:"not null default '' comment('保存条件') VARCHAR(500)"`
	StoreUnit              string    `json:"store_unit" xorm:"not null default '' comment('出入库单位ID') VARCHAR(255)"`
	StructId               string    `json:"struct_id" xorm:"not null default '' comment('区域ID') VARCHAR(500)"`
	SupplierCode           string    `json:"supplier_code" xorm:"not null default '0' comment('供应商编码') VARCHAR(30)"`
	SupplierInnerCode      string    `json:"supplier_inner_code" xorm:"not null default '' comment('生产商内部编码') VARCHAR(200)"`
	Tax                    string    `json:"tax" xorm:"not null default 0.00 comment('税额') DECIMAL(30,2)"`
	TaxRate                string    `json:"tax_rate" xorm:"not null default 0.00 comment('税率') DECIMAL(30,2)"`
	Traits                 string    `json:"traits" xorm:"not null default '' comment('性状') VARCHAR(255)"`
	UntaxedPrice           string    `json:"untaxed_price" xorm:"not null default 0.00 comment('未税单价') DECIMAL(30,2)"`
	WarehouseCode          string    `json:"warehouse_code" xorm:"default '' comment('大仓编码') VARCHAR(50)"`
}
