package excel

func GetGoodsHead() (arr []map[string]string) {
	//var arr [...]map[string]interface{}

	arr = []map[string]string{
		{
			"name":      "一级目录",
			"field":     "category_name1",
			"is_master": "1",
		},
		{
			"name":      "二级目录",
			"field":     "category_name2",
			"is_master": "0",
		},
		{
			"name":      "三级目录",
			"field":     "category_name3",
			"is_master": "0",
		},
		{
			"name":      "四级目录",
			"field":     "category_name4",
			"is_master": "0",
		},
		{
			"name":      "产品名称",
			"field":     "product_name",
			"is_master": "1",
		},
		{
			"name":      "英文名称",
			"field":     "english_name",
			"is_master": "0",
		},
		{
			"name":      "通用名",
			"field":     "common_name",
			"is_master": "1",
		},
		{
			"name":      "生产商名称",
			"field":     "supplier_name",
			"is_master": "1",
		},
		{
			"name":      "品牌名称",
			"field":     "brand_name",
			"is_master": "1",
		},
		{
			"name":      "包装规格",
			"field":     "pack_specific",
			"is_master": "1",
		},
		{
			"name":      "条形码",
			"field":     "bar_code",
			"is_master": "1",
		},
		{
			"name":      "是否非禁商品：0-否 1-是",
			"field":     "not_banned",
			"is_master": "1",
		},
		{
			"name":      "是否效期管理：0-否 1-是",
			"field":     "is_validity_control",
			"is_master": "1",
		},

		{
			"name":      "说明书有效期",
			"field":     "instructions_validity",
			"is_master": "0",
		},
		{
			"name":      "是否禁止负库存销售：0-否 1-是",
			"field":     "forbid_negative_stock",
			"is_master": "1",
		},
		{
			"name":      "出入库单位",
			"field":     "store_unit_name",
			"is_master": "1",
		},
		{
			"name":      "入库含税参考价",
			"field":     "ref_price",
			"is_master": "0",
		},
		{
			"name":      "销售价格",
			"field":     "sell_price",
			"is_master": "1",
		},
		{
			"name":      "销售价格区间（最小值）",
			"field":     "sell_price_low",
			"is_master": "0",
		},
		{
			"name":      "销售价格区间（最大值）",
			"field":     "sell_price_high",
			"is_master": "0",
		},
		{
			"name":      "是否打折：0-否 1-是",
			"field":     "is_discount",
			"is_master": "1",
		},
		{
			"name":      "税率%",
			"field":     "tax_rate",
			"is_master": "0",
		},
		{
			"name":      "未税单价",
			"field":     "untaxed_price",
			"is_master": "0",
		},
		{
			"name":      "是否可订，1是，0否",
			"field":     "can_order",
			"is_master": "1",
		},
		{
			"name":      "是否可销，1是，0否",
			"field":     "can_sell",
			"is_master": "1",
		},
		{
			"name":      "是否停用，1是，0否",
			"field":     "disabled",
			"is_master": "1",
		},
		{
			"name":      "本位币",
			"field":     "currency_name",
			"is_master": "1",
		},
		{
			"name":      "保存条件(多选可用英文逗号隔开',')",
			"field":     "store_condition_name",
			"is_master": "1",
		},
		{
			"name":      "适用宠物(多选可用英文逗号隔开',')",
			"field":     "applicable_pet_name",
			"is_master": "0",
		},
		{
			"name":      "性状(多选可用英文逗号隔开',')",
			"field":     "traits_name",
			"is_master": "0",
		},
		{
			"name":      "批准文号",
			"field":     "approval_sn",
			"is_master": "1",
		},
		{
			"name":      "生产商内部编码",
			"field":     "supplier_inner_code",
			"is_master": "0",
		},
		{
			"name":      "生产国和地区",
			"field":     "cninfo",
			"is_master": "1",
		},
		{
			"name":      "有效成分",
			"field":     "ingredient",
			"is_master": "1",
		},
		{
			"name":      "备注",
			"field":     "remark",
			"is_master": "0",
		},
		{
			"name":      "临时编码",
			"field":     "indication",
			"is_master": "0",
		},
	}
	return arr
}
