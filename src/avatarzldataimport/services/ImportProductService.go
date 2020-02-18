package services

import (
	"avatarzldataimport/models/excel"
	"avatarzldataimport/models/hmd"
	"avatarzldataimport/modules/util/pinyin"
	"fmt"
	"strconv"
	"strings"
)

type Params struct {
	Type int
	Path string
}

var Trans Checker
var Res []map[string]interface{}
var ExcelHeader []map[string]string

func (params *Params) ImportProduct() (msg string) {
	//var msg string
	if params.Type == 1010 {
		msg = params.importGoods()
	}
	Res = nil
	if msg != "" {
		ProductObj = nil
		return
	} else {
		//fmt.Println(ProductObj)
		hmd.InsertProduct(ProductObj)
	}
	ProductObj = nil
	return

}

func (params *Params) importGoods() (msg string) {
	ExcelHeader = excel.GetGoodsHead()
	ExcelData, err := GetExcelData(params.Path)
	if err != nil {
		fmt.Println(err)
	}
	initGoodsCheck()
	resultChan := make(chan map[string]interface{}, len(ExcelData))
	flagChan := make(chan bool, 1)
	go func() {
		for i, excelData := range ExcelData {
			if len(excelData) == 0 {
				continue
			}
			if excelData[0] == "" {
				continue
			}
			if i == 0 { //校验表头
				if len(excelData) != len(ExcelHeader) || !checkHeader(excelData) {
					msg = msg + "请校验表头是否准确"
					return
				}
				continue
			}
			getResData(excelData, resultChan)

		}
		close(resultChan)
	}()
	go func() {
		for {
			res, ok := <-resultChan
			if !ok {
				flagChan <- true
				close(flagChan)
				break
			}
			if res["msg"] != nil {
				msg = msg + res["msg"].(string)
			}
			GetInsertGoodsData(res)

		}
	}()
	for {
		if <-flagChan {
			return
		}
	}

}

func getResData(excelData []string, resultChan chan map[string]interface{}) {
	//excelData := <-channel
	// 先进行key->value拼接
	res := implodeKeyValue(excelData)
	res = getExcelCategoryPath(res)
	res = translateField(res)
	res["product_code"] = getProductCode(res["product_type"].(int))
	resultChan <- res
}


func getProductCode(productType int)(code string){
	if productType == 1010 {
		code  = "S000"
	}
	code = code + RandCreateString(6)
	if Trans.ProductCodeArr[code] == 1{
		getProductCode(productType)
	}

	return
}

func initGoodsCheck() {
	Trans.InitApplication()
	Trans.InitBarCode()
	Trans.InitCategory()
	Trans.InitBrandCode()
	Trans.InitCountry()
	Trans.InitCurrency()
	Trans.InitType()
	//Trans.InitDosageUnit()
	Trans.InitProductName()
	Trans.InitProductCode()
	Trans.InitStoreCondition()
	Trans.InitStoreUnit()
	Trans.InitSupplier()
	Trans.InitTraitsArr()
}

func translateField(data map[string]interface{}) (res map[string]interface{}) {
	if Trans.ProductNameArr[data["product_name"].(string)] == 1 { //证明产品存在
		data["msg"] = data["msg"].(string) + "产品【" + data["product_name"].(string) + "】已存在"
	} else {
		Trans.ProductNameArr[data["product_name"].(string)] = 1
	}
	if Trans.BarCodeArr[data["bar_code"].(string)] == 1 && data["bar_code"].(string) != "无" { //证明产品存在
		data["msg"] = data["msg"].(string) + "条形码【" + data["bar_code"].(string) + "】已存在"
		Trans.ProductNameArr[data["bar_code"].(string)] = 1
	}

	if Trans.BrandArr[data["brand_name"].(string)] == 0 {
		fmt.Println(Trans.BrandArr[data["brand_name"].(string)], data["brand_name"])
		data["msg"] = data["msg"].(string) + "品牌【" + data["brand_name"].(string) + "】有误"
	}
	data["brand_code"] = Trans.BrandArr[data["brand_name"].(string)]

	if Trans.CurrencyArr[data["currency_name"].(string)] == 0 {
		data["msg"] = data["msg"].(string) + "本位币【" + data["currency_name"].(string) + "】有误"
	}
	data["currency"] = Trans.CurrencyArr[data["currency_name"].(string)]

	if data["obj_type_name"] != nil {
		data["obj_type"] = Trans.ObjTypeArr[data["obj_type_name"].(string)]
	} else {
		data["obj_type"] = Trans.ObjTypeArr[data["category_name1"].(string)]
	}
	data["product_type"] = Trans.ProductTypeArr[data["category_name1"].(string)]

	if Trans.SupplierArr[data["supplier_name"].(string)] == "" {
		data["msg"] = data["msg"].(string) + "生产商【" + data["supplier_name"].(string) + "】有误"
	}
	data["supplier_code"] = Trans.SupplierArr[data["supplier_name"].(string)]

	if Trans.CountryArr[data["cninfo"].(string)] == 0 {
		data["msg"] = data["msg"].(string) + "生产国【" + data["cninfo"].(string) + "】有误"
	}
	data["country_code"] = Trans.CountryArr[data["cninfo"].(string)]

	if Trans.StoreUnitArr[data["store_unit_name"].(string)] == 0 {
		data["msg"] = data["msg"].(string) + "出入库单位【" + data["store_unit_name"].(string) + "】有误"
	}
	data["store_unit"] = Trans.StoreUnitArr[data["store_unit_name"].(string)]

	if Trans.CategoryArr[data["category_path"].(string)] == "" {
		data["msg"] = data["msg"].(string) + "目录【" + data["category_path"].(string) + "】有误"
	}
	data["category"] = Trans.CategoryArr[data["category_path"].(string)]
	if data["store_condition_name"] != "" {
		storeConditionArr := strings.Split(data["store_condition_name"].(string), ",")
		for _, v := range storeConditionArr {
			storeCondition := Trans.StoreConditionArr[v]
			if storeCondition != 0 {
				if data["store_condition"] == nil {
					data["store_condition"] = strconv.Itoa(storeCondition)
				} else {
					data["store_condition"] = data["store_condition"].(string) + "," + strconv.Itoa(storeCondition)
				}
			} else {
				data["msg"] = data["msg"].(string) + "保存条件【" + v + "】有误"
			}
		}
	}

	if data["applicable_pet_name"] != "" {
		ApplicableArr := strings.Split(data["applicable_pet_name"].(string), ",")
		for _, v := range ApplicableArr {
			applicablePet := Trans.ApplicationArr[v]
			if applicablePet != 0 {
				if data["applicable_pet"] == nil {
					data["applicable_pet"] = strconv.Itoa(applicablePet)
				} else {
					data["applicable_pet"] = data["applicable_pet"].(string) + "," + strconv.Itoa(applicablePet)
				}
			} else {
				data["msg"] = data["msg"].(string) + "适用宠物【" + v + "】有误"
			}
		}
	} else {
		data["applicable_pet"] = ""
	}

	if data["traits_name"] != "" {
		traitsArr := strings.Split(data["traits_name"].(string), ",")
		for _, v := range traitsArr {
			traits := Trans.TraitsArr[v]
			if traits != 0 {
				if data["traits"] == nil {
					data["traits"] = strconv.Itoa(traits)
				} else {
					data["traits"] = data["traits"].(string) + "," + strconv.Itoa(traits)
				}
			} else {
				data["msg"] = data["msg"].(string) + "性状【" + v + "】有误"
			}
		}
	}
	data["pinyin_abbr"], _ = pinyin.New(data["product_name"].(string)).Split("").Mode(pinyin.GetFirstUpper).Convert()

	res = make(map[string]interface{})
	res = data
	return
}

func getExcelCategoryPath(data map[string]interface{}) (res map[string]interface{}) {

	data["category_path"] = data["category_name1"].(string)
	if data["category_name2"].(string) != "" {
		data["category_path"] = data["category_path"].(string) + ">" + data["category_name2"].(string)
	}
	if data["category_name3"] != "" {
		data["category_path"] = data["category_path"].(string) + ">" + data["category_name3"].(string)
	}
	if data["category_name4"] != "" {
		data["category_path"] = data["category_path"].(string) + ">" + data["category_name4"].(string)
	}
	res = make(map[string]interface{})
	res = data
	return
}

// 对值与key value 进行拼装
func implodeKeyValue(excelData []string) (res map[string]interface{}) {
	headers := ExcelHeader
	res = make(map[string]interface{})
	//for i, v := range excelData {
	//	res[headers[i]["field"]] = v
	//	if headers[i]["field"] == "remark" {
	//		fmt.Println(headers[i]["field"], "-----",v)
	//	}
	//	if headers[i]["is_master"] == "1" { // 如果是必填项校验非空
	//		if v == "" {
	//			res["msg"] = headers[i]["name"] + "不能为空，"
	//		}
	//	}
	//}
	res["msg"] = ""
	leng := len(excelData) - 1
	for i, v := range headers {
		val := ""
		if i > leng {
			val = ""
		} else {
			val = excelData[i]
		}

		res[v["field"]] = val
		if v["is_master"] == "1" && val == "" {
			res["msg"] = headers[i]["name"] + "不能为空，"
		}
	}
	return
}

// 校验表头
func checkHeader(data []string) (res bool) {
	headers := ExcelHeader
	res = true
	for i, v := range data {
		if headers[i]["name"] != v { // 如果是必填项校验非空
			res = false
		}
	}
	return
}
