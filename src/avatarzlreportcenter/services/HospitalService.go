package services

import (
	"avatarzlreportcenter/models/report"
	"avatarzlreportcenter/modules/util"
	"fmt"
	"github.com/goinggo/mapstructure"
)

func GetHospitalReport(where interface{}) (returnData []report.HospitalReport, err error) {
	data, err := report.GetHospital(where)
	count := report.HospitalReport{}
	value := util.GetFieldName(report.HospitalReport{})
	fmt.Println(value)
	var structHospital report.HospitalReport

	testMap := make(map[string]interface{})
	testMap["OrgId"] = 8000
	mapstructure.Decode(testMap,&structHospital)
	fmt.Println(structHospital)
	for _, item := range (data) {
		//testMap[k] = make(map[string]string)
		//for _, keys := range (value) {
		//	fmt.Println(keys)
		//	testMap[k][keys] = "dsdf "
		//}
		count.AllotInAmount = (item.AllotInAmount*100 + count.AllotInAmount*100) / 100
		count.AutTotalAmount += item.AutTotalAmount
		count.AllotOutAmount += item.AllotOutAmount
		count.PastOutAmount += item.PastOutAmount
		count.PurchasingInAmount += item.PurchasingInAmount
		count.PurchasingTotalAmount += item.PurchasingTotalAmount
		count.CheckInAmount += item.CheckInAmount
		count.CheckOutAmount += item.CheckOutAmount
		count.DamageOutAmount += item.DamageOutAmount
		count.ReturnOutAmount += item.ReturnOutAmount
		count.StockAmountBegin = (item.StockAmountBegin*100 + count.StockAmountBegin*100) / 100
		count.StockAmountEnd += item.StockAmountEnd
		count.SellOutAmount += item.SellOutAmount
		count.SellReturnAmount += item.SellReturnAmount
		count.UseOutAmount += item.UseOutAmount
		count.SurplusInAmount += item.SurplusInAmount
		count.InTotalAmount += item.InTotalAmount
	}
	fmt.Println(testMap, "===")
	count.ProductTypeName = "合计"
	returnData = append(data, count)
	return

}
