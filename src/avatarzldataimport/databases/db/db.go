package db

import (
	"avatarzldataimport/config"
	"avatarzldataimport/modules/log"
	"fmt"
	"github.com/go-xorm/xorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"log/syslog"
	"os"
)

var DB *xorm.Engine

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
	DB, err = xorm.NewEngine(db_type, fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=%s",
		user,
		password,
		host_url,
		db_name, charset))
	//DB, err := gorm.Open(config.DBConfig.Dialect, config.DBConfig.URL)
	if err != nil {
		log.ErrorLog.Println(err)
	}
	DB.ShowSQL(true)
	f, err := os.OpenFile("sql.log",os.O_CREATE|os.O_WRONLY|os.O_APPEND, os.ModePerm)
	if err != nil {
		println(err.Error())
		return
	}
	DB.SetLogger(xorm.NewSimpleLogger(f))


	logWriter, err := syslog.New(syslog.LOG_DEBUG, "rest-xorm-example")
	if err != nil {
		log.ErrorLog.Println("Fail to create xorm system logger: %v\n", err)
	}

	logger := xorm.NewSimpleLogger(logWriter)
	logger.ShowSQL(true)
	DB.SetLogger(logger)

	DB.SetMaxIdleConns(config.DBConfig.MaxIdleConns)
	// set max open connections
	DB.SetMaxOpenConns(config.DBConfig.MaxOpenConns)

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
