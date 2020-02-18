package db

import (
	"fmt"
	"github.com/jinzhu/gorm"
	"log"
	"wx_mini/models/setting"
)

var db *gorm.DB

type Mode struct {
	ID int

}

func init(){
	var(
		err error
		db_type,db_name,user,password,host,port string
	)
	sec,err := setting.Cfg.GetSection("database")
	if err != nil {
		log.Fatal(2, "Fail to get section 'database': %v", err)
	}
	db_type = sec.Key("TYPE").String()
	db_name = sec.Key("NAME").String()
	user = sec.Key("USER").String()
	password = sec.Key("PASSWORD").String()
	host = sec.Key("HOST").String()
	port = sec.Key("PORT").String()
	host_url := host+":"+port
	db, err = gorm.Open(db_type, fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8&parseTime=True&loc=Local",
		user,
		password,
		host_url,
		db_name))
	if err != nil {
		log.Println(err)
	}
	gorm.DefaultTableNameHandler = func (db *gorm.DB, defaultTableName string) string  {
		return   defaultTableName
	}
	db.SingularTable(true)
	db.DB().SetMaxIdleConns(10)
	db.DB().SetMaxOpenConns(100)
}
func CloseDB() {
	defer db.Close()
}