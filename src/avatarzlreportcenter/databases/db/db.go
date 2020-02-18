package db

import (
	"avatarzlreportcenter/config"
	"avatarzlreportcenter/modules/log"
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var DB *gorm.DB

func init() {
	var (
		err                                                   error
		db_type, db_name, user, password, host, port, charset string
	)
	db_type = config.DBConfig.Dialect
	db_name = config.DBConfig.Database
	user = config.DBConfig.User
	password = config.DBConfig.Password
	host = config.DBConfig.Host
	port = config.DBConfig.Port
	host_url := host + ":" + port
	charset = config.DBConfig.Charset
	DB, err = gorm.Open(db_type, "")
	DB, err = gorm.Open(db_type, fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=%s",
		user,
		password,
		host_url,
		db_name, charset))
	//DB, err := gorm.Open(config.DBConfig.Dialect, config.DBConfig.URL)
	if err != nil {
		log.ErrorLog.Println(err)
	}
	gorm.DefaultTableNameHandler = func(DB *gorm.DB, defaultTableName string) string {
		return defaultTableName
	}
	DB.LogMode(true)
	//defer DB.Close()

	logger := &MyLogger{}

	DB.LogMode(true)
	DB.SetLogger(logger)

	DB.SingularTable(true)
	DB.DB().SetMaxIdleConns(config.DBConfig.MaxIdleConns)
	// set max open connections
	DB.DB().SetMaxOpenConns(config.DBConfig.MaxOpenConns)

}

type MyLogger struct {
}

func (logger *MyLogger) Print(values ...interface{}) {
	var (
		level           = values[0]
		//source          = values[1]
	)

	if level == "sql" {
		log.InfoLog.Println(values[0])
		log.InfoLog.Println(values[1])
		log.InfoLog.Println(values[2])
		log.InfoLog.Println(values[3].(string))
	} else {
		log.ErrorLog.Println(values[0])
		log.ErrorLog.Println(values[1])
		log.ErrorLog.Println(values[2])
		log.ErrorLog.Println(values[3].(string))

	}
}
