package services

import (
	"avatarzldataimport/models/hmd"
	"strconv"
	"time"
)

var ProductObj []*hmd.THosProduct

func GetInsertGoodsData(data map[string]interface{}) {
	insertObj := hmd.THosProduct{}
	insertObj.ProductType = data["product_type"].(int)
	insertObj.ProductName = data["product_name"].(string)
	insertObj.ProductCode = data["product_code"].(string)
	insertObj.ObjType = data["obj_type"].(int)
	insertObj.CompanyId = 2
	insertObj.Category = data["category"].(string)
	insertObj.Indication = data["ingredient"].(string)
	insertObj.SupplierCode = data["supplier_code"].(string)
	insertObj.BrandCode = strconv.Itoa(data["brand_code"].(int))
	insertObj.BrandName = data["brand_name"].(string)
	insertObj.PackSpecific = data["pack_specific"].(string)
	insertObj.PackSpecific = data["pack_specific"].(string)
	insertObj.BarCode = data["bar_code"].(string)
	insertObj.CommonName = data["common_name"].(string)
	insertObj.PinyinAbbr = data["pinyin_abbr"].(string)
	insertObj.SupplierInnerCode = data["supplier_inner_code"].(string)
	insertObj.NotBanned, _ = strconv.Atoi(data["not_banned"].(string))
	insertObj.IsValidityControl, _ = strconv.Atoi(data["is_validity_control"].(string))
	insertObj.InstructionsValidity, _ = strconv.Atoi(data["instructions_validity"].(string))
	insertObj.ForbidNegativeStock, _ = strconv.Atoi(data["forbid_negative_stock"].(string))
	insertObj.StoreUnit = strconv.Itoa(data["store_unit"].(int))
	insertObj.RefPrice, _ = strconv.ParseFloat(data["ref_price"].(string), 32)
	insertObj.SellPrice, _ = strconv.ParseFloat(data["sell_price"].(string), 32)
	insertObj.SellPriceLow = data["sell_price_low"].(string)
	insertObj.ApprovalSn = data["approval_sn"].(string)
	insertObj.SellPriceHigh = data["sell_price_high"].(string)
	insertObj.IsDiscount, _ = strconv.Atoi(data["is_discount"].(string))
	insertObj.TaxRate, _ = strconv.ParseFloat(data["tax_rate"].(string), 32)
	insertObj.UntaxedPrice, _ = strconv.ParseFloat(data["untaxed_price"].(string), 32)
	insertObj.CanOrder, _ = strconv.Atoi(data["can_order"].(string))
	insertObj.CanSell, _ = strconv.Atoi(data["can_sell"].(string))
	insertObj.Disabled, _ = strconv.Atoi(data["disabled"].(string))
	insertObj.Currency = data["currency"].(int)
	insertObj.StoreCondition = data["store_condition"].(string)
	insertObj.ApplicablePet = data["applicable_pet"].(string)
	insertObj.CountryCode = strconv.Itoa(data["country_code"].(int))
	insertObj.Remark = data["remark"].(string)
	insertObj.Indication = data["indication"].(string)
	insertObj.AddTime = time.Now().Format("2006-01-02 15:04:05")
	insertObj.LastModify = time.Now().Format("2006-01-02 15:04:05")
	insertObj.IsReuse = 1      // 默认
	insertObj.NonSensitive = 1 // 默认
	insertObj.From = 2         // 默认

	ProductObj = append(ProductObj, &insertObj)
}
