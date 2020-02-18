package config

import (
	"fmt"
	"github.com/go-ini/ini"
	"github.com/goinggo/mapstructure"
	"io/ioutil"
	"os"
	"regexp"
	"wx_mini/modules/log"
	"wx_mini/modules/util"
)

var jsonData map[string]interface{}

func init() {
	readJson()
	initDb()
	initServer()
	initWeApp()
	initRedis()
	initCookie()
	initHttp()
}
func readJson() {
	//var jsonAllData map[string] string
	cfg,err := ini.Load("config/app.ini")
	if err != nil {
		log.ErrorLog.Println(err.Error())
	}
	model := cfg.Section("").Key("appModel").MustString("dev")
	fmt.Println(model)
	GOPATH := os.Getenv("GOPATH")
	file := GOPATH + "/src/wx_mini/config.json"
	bytes, err := ioutil.ReadFile(file)
	if err != nil {
		log.ErrorLog.Println(err.Error())
	}

	configStr := string(bytes[:])

	reg := regexp.MustCompile(`/\*.*\*/`)
	configStr = reg.ReplaceAllString(configStr, "")
	bytes = []byte(configStr)
	if jsonData, err = util.JsonToMap(bytes); err != nil {
		log.ErrorLog.Println(err.Error())
	}
	jsonData = jsonData[model].(map[string]interface{})
}

func initDb() {
	if err := mapstructure.Decode(jsonData["database"].(map[string]interface{}), &DBConfig); err != nil {
		log.ErrorLog.Println(err.Error())
	}
	//set database url
	DBConfig.URL = fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=%s&parseTime=True&loc=Local",
		DBConfig.User, DBConfig.Password, DBConfig.Host, DBConfig.Port, DBConfig.Database, DBConfig.Charset)
}

func initServer() {
	if err := mapstructure.Decode(jsonData["apiServer"].(map[string]interface{}), &ServerConfig); err != nil {
		log.ErrorLog.Println(err.Error())
	}
}

func initWeApp() {
	if err := mapstructure.Decode(jsonData["weApp"].(map[string]interface{}), &WeAppConfig); err != nil {
		log.ErrorLog.Println(err.Error())
	}
}

func initRedis() {
	if err := mapstructure.Decode(jsonData["redis"].(map[string]interface{}), &RedisConfig); err != nil {
		log.ErrorLog.Println(err.Error())
	}
}

func initCookie() {
	if err := mapstructure.Decode(jsonData["cookie"].(map[string]interface{}), &CookieConfig); err != nil {
		log.ErrorLog.Println(err.Error())
	}
}

func initHttp() {
	if err := mapstructure.Decode(jsonData["requestHost"].(map[string]interface{}), &HttpRequestConfig); err != nil {
		log.ErrorLog.Println(err.Error())
	}
}