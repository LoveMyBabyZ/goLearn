package db

import (
	"avatarzldataimport/config"
	"avatarzldataimport/modules/log"
	"fmt"
	"github.com/go-xorm/xorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"log/syslog"
)

var HospitalDb *xorm.Engine

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
	fmt.Println(host_url)
	charset = config.HospitalDBConfig.Charset
	HospitalDb, err = xorm.NewEngine(db_type, fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=%s",
		user,
		password,
		host_url,
		db_name, charset))
	if err != nil {
		log.ErrorLog.Println(err)
	}

	logWriter, err := syslog.New(syslog.LOG_DEBUG, "rest-xorm-example")
	if err != nil {
		log.ErrorLog.Println("Fail to create xorm system logger: %v\n", err)
	}

	logger := xorm.NewSimpleLogger(logWriter)
	logger.ShowSQL(true)
	HospitalDb.SetLogger(logger)

	HospitalDb.SetMaxIdleConns(config.HospitalDBConfig.MaxIdleConns)
	// set max open connections
	HospitalDb.SetMaxOpenConns(config.HospitalDBConfig.MaxOpenConns)

}
