package hospital

import (
	"avatarzlreportcenter/databases/db"
	"avatarzlreportcenter/modules/log"
)

const (
	ONLINE_HOPSITAL_DATE = "sys_online_date_info"
)

type OnlineHospital struct {
	Id         int    `gorm:"column:id; type:int(11); primary_key; not null" json:"id"`
	OrgId      int    `gorm:"column:org_id; type:int(11); not null" json:"org_id"`
	OrgCode    string `gorm:"column:org_code; type:varchar(50); not null" json:"org_id"`
	OnlineDate string `gorm:"column:online_date; type:datetime" json:"online_date"`
	ReportDate string `gorm:"column:report_date; type:datetime" json:"report_date"`
}

func init() {
	if !db.HospitalDB.HasTable(ONLINE_HOPSITAL_DATE) {
		if err := db.HospitalDB.Set("gorm:table_options", "ENGINE=InnoDB DEFAULT CHARSET=utf8mb4").CreateTable(&OnlineHospital{}).Error; err != nil {
			log.ErrorLog.Println(err)
			return
		}
	}
}

// set table name
func (u OnlineHospital) TableName() string {
	return ONLINE_HOPSITAL_DATE
}

func GetAllOnline() (hospital []OnlineHospital, err error) {
	db.HospitalDB.Select("org_id,org_code,online_date,report_date").Find(&hospital)
	return
}

func GetByCodeOrId(maps interface{}) (hospital []OnlineHospital, err error) {
	db.HospitalDB.Select("org_id,org_code,online_date,report_date").Where(maps).First(&hospital)
	return
}
