package services

import (
	"avatarzldataimport/models/hmd"
	"avatarzldataimport/models/hospital"
	"avatarzldataimport/modules/log"
)

type Checker struct {
	ProductNameArr    map[string]int
	ProductCodeArr    map[string]int
	SupplierArr       map[string]string
	BarCodeArr        map[string]int
	BrandArr          map[string]int
	CountryArr        map[string]int
	CurrencyArr       map[string]int
	DosageUnitArr     map[string]int
	StoreUnitArr      map[string]int
	StoreConditionArr map[string]int
	ApplicationArr    map[string]int
	TraitsArr         map[string]int
	ProductTypeArr    map[string]int
	ObjTypeArr        map[string]int
	CategoryArr       map[string]string
}

// 目录
func (check *Checker) InitCategory() {
	data, err := hmd.GetCategoryPath()
	if err != nil {
		log.ErrorLog.Println("获取产品目录有误")
	}
	check.CategoryArr = data
}

// 获取已经存在的产品名称
func (check *Checker) InitProductName() {
	data, err := hmd.GetProductName()
	if err != nil {
		log.ErrorLog.Println("获取产品名称有误")
	}
	check.ProductNameArr = data
}

// 获取供应商
func (check *Checker) InitSupplier() {
	data, err := hmd.GetSupplierArr()
	if err != nil {
		log.ErrorLog.Println("供应商获取有误")
	}
	check.SupplierArr = data
}

// 获取已经存在的条形码
func (check *Checker) InitBarCode() {
	data, err := hmd.GetBarCode()
	if err != nil {
		log.ErrorLog.Println("条形码获取有误")
	}
	check.BarCodeArr = data
}

// 获取已经存在的条形码
func (check *Checker) InitProductCode() {
	data, err := hmd.GetProductCode()
	if err != nil {
		log.ErrorLog.Println("已有产品编码获取有误")
	}
	check.ProductCodeArr = data
}

// 产品类型
func (check *Checker) InitType() {
	data, err := hmd.GetProductType()
	if err != nil {
		log.ErrorLog.Println("产品类型获取有误")
	}
	check.ProductTypeArr = data
	objData, err := hmd.GetProductObjType()
	if err != nil {
		log.ErrorLog.Println("对象类型获取有误")
	}
	check.ObjTypeArr = objData
}



// 品牌字典
func (check *Checker) InitBrandCode() {
	data, err := hospital.GetDict(15214)
	if err != nil {
		log.ErrorLog.Println("获取品牌字典失败")
	}
	check.BrandArr = data
}

// 国家字典
func (check *Checker) InitCountry() {
	data, err := hospital.GetDict(930001)
	if err != nil {
		log.ErrorLog.Println("获取国家字典失败")
	}
	check.CountryArr = data
}

// 本位币字典
func (check *Checker) InitCurrency() {
	data, err := hospital.GetDict(4001)
	if err != nil {
		log.ErrorLog.Println("获取本位币字典失败")
	}
	check.CurrencyArr = data
}

// 剂型字典
func (check *Checker) InitDosageUnit() {
	data, err := hospital.GetDict(5302)
	if err != nil {
		log.ErrorLog.Println("获取剂型字典失败")
	}
	check.DosageUnitArr = data
}

// 库存类单位字典
func (check *Checker) InitStoreUnit() {
	data, err := hospital.GetDict(4355)
	if err != nil {
		log.ErrorLog.Println("获取库存字典失败")
	}
	check.StoreUnitArr = data
}

// 保存条件字典
func (check *Checker) InitStoreCondition() {
	data, err := hospital.GetDict(911000)
	if err != nil {
		log.ErrorLog.Println("获取保存条件字典失败")
	}
	check.StoreConditionArr = data
}

func (check *Checker) InitApplication() {
	data, err := hospital.GetDict(1003)
	if err != nil {
		log.ErrorLog.Println("获取适用宠物字典失败")
	}
	check.ApplicationArr = data
}

func (check *Checker) InitTraitsArr() {
	data, err := hospital.GetDict(935458)
	if err != nil {
		log.ErrorLog.Println("获取性状字典失败")
	}
	check.TraitsArr = data
}
