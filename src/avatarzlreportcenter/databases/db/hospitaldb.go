package db

import (
	"avatarzlreportcenter/config"
	"avatarzlreportcenter/modules/log"
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var HospitalDB *gorm.DB

func init() {
	var (
		err                                                   error
		db_type, db_name, user, password, host, port, charset string
	)
	db_type = config.HospitalDBConfig.Dialect
	db_name = config.HospitalDBConfig.Database
	user = config.HospitalDBConfig.User
	password = config.HospitalDBConfig.Password
	host = config.HospitalDBConfig.Host
	port = config.HospitalDBConfig.Port
	host_url := host + ":" + port
	charset = config.HospitalDBConfig.Charset
	//HospitalDB, err = gorm.Open(db_type, "")
	HospitalDB, err = gorm.Open(db_type, fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=%s",
		user,
		password,
		host_url,
		db_name, charset))
	//DB, err := gorm.Open(config.DBConfig.Dialect, config.DBConfig.URL)
	if err != nil {
		log.ErrorLog.Println(err)
	}
	gorm.DefaultTableNameHandler = func(HospitalDB *gorm.DB, defaultTableName string) string {
		return defaultTableName
	}
	HospitalDB.LogMode(true)
	//defer DB.Close()
	fmt.Println(host_url)
	logger := &MyLogger{}

	HospitalDB.LogMode(true)
	HospitalDB.SetLogger(logger)

	HospitalDB.SingularTable(true)
	HospitalDB.DB().SetMaxIdleConns(config.HospitalDBConfig.MaxIdleConns)
	// set max open connections
	HospitalDB.DB().SetMaxOpenConns(config.HospitalDBConfig.MaxOpenConns)

}

