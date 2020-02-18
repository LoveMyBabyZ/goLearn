package report

import (
	"avatarzlreportcenter/databases/db"
	"avatarzlreportcenter/modules/log"
)

const (
	HOSPITALREPOT = "erp_org_report"
)

type HospitalReport struct {
	Id                    int     `gorm:"column:id; type:int(11); primary_key; not null" json:"id"`
	OrgId                 int     `gorm:"column:org_id; type:int(11); not null" json:"org_id"`
	ReportDate            string  `gorm:"column:report_id;type:varchar(50);not null" json:"org_id"`
	ProductTypeName       string  `gorm:"column:product_type_name;type:varchar(30)" json:"product_type_name"`
	StockAmountBegin      float64 `gorm:"column:stock_amount_begin;type:decimal(18,8)" json:"stock_amount_begin"`
	PurchasingInAmount    float64 `gorm:"column:purchasing_in_amount;type:decimal(18,8)" json:"purchasing_in_amount"`
	ReturnOutAmount       float64 `gorm:"column:return_out_amount;type:decimal(18,8)" json:"return_out_amount"`
	PurchasingTotalAmount float64 `gorm:"column:purchasing_total_amount;type:decimal(18,8)" json:"purchasing_total_amount"`
	AllotInAmount         float64 `gorm:"column:allot_in_amount;type:decimal(18,8)" json:"allot_in_amount"`
	InTotalAmount         float64 `gorm:"column:in_total_amount;type:decimal(18,8)" json:"in_total_amount"`
	SellOutAmount         float64 `gorm:"column:sell_out_amount;type:decimal(18,8)" json:"sell_out_amount"`
	UseOutAmount          float64 `gorm:"column:use_out_amount;type:decimal(18,8)" json:"use_out_amount"`
	SurplusInAmount       float64 `gorm:"column:surplus_in_amount;type:decimal(18,8)" json:"surplus_in_amount"`
	SellReturnAmount      float64 `gorm:"column:sell_return_amount;type:decimal(18,8)" json:"sell_return_amount"`
	AutTotalAmount        float64 `gorm:"column:out_total_amount;type:decimal(18,8)" json:"out_total_amount"`
	AllotOutAmount        float64 `gorm:"column:allot_out_amount;type:decimal(18,8)" json:"allot_out_amount"`
	PastOutAmount         float64 `gorm:"column:past_out_amount;type:decimal(18,8)" json:"past_out_amount"`
	DamageOutAmount       float64 `gorm:"column:damage_out_amount;type:decimal(18,8)" json:"damage_out_amount"`
	CheckInAmount         float64 `gorm:"column:check_in_amount;type:decimal(18,8)" json:"check_in_amount"`
	CheckOutAmount        float64 `gorm:"column:check_out_amount;type:decimal(18,8)" json:"check_out_amount"`
	StockAmountEnd        float64 `gorm:"column:stock_amount_end;type:decimal(18,8)" json:"stock_amount_end"`
	CreateDate            string  `gorm:"column:create_date;type:datetime" json:"create_date"`
}

func init() {
	if !db.DB.HasTable(HOSPITALREPOT) {
		if err := db.DB.Set("gorm:table_options", "ENGINE=InnoDB DEFAULT CHARSET=utf8mb4").CreateTable(&HospitalReport{}).Error; err != nil {
			log.ErrorLog.Println(err)
			return
		}
	}
}

//set table name
func (u HospitalReport) TableName() string {
	return HOSPITALREPOT
}

func GetHospital(maps interface{}) (report []HospitalReport, err error) {
	db.DB.Where(maps).Find(&report)
	return
}


