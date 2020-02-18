package config

import (
	"avatarzlreportcenter/modules/log"
	"avatarzlreportcenter/modules/util"
	"fmt"
	"github.com/Unknwon/goconfig"
	"github.com/goinggo/mapstructure"
	"io/ioutil"
	"os"
	"regexp"
)

var jsonData map[string]interface{}

func init() {
	readJson()
	initDb()
	initHospitalDb()
	initServer()
	initWeApp()
	initRedis()
	initCookie()
	initHttp()
}
func readJson() {
	//var jsonAllData map[string] string
	//cfg, err := ini.Load("app.ini")
	cfg,err := goconfig.LoadConfigFile("config/app.ini")
	if err != nil {
		log.ErrorLog.Println(err.Error())
	}
	//model := cfg.Section("").Key("appModel").MustString("dev")
	model,err := cfg.GetSection("appModel")
	dir := model["setting"]
	fmt.Println(model)
	GOPATH := os.Getenv("GOPATH")
	file := GOPATH + "/src/avatarzlreportcenter/config/" + dir + "/config.json"
	bytes, err := ioutil.ReadFile(file)
	fmt.Println(file)
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
}

func initDb() {
	if err := mapstructure.Decode(jsonData["database"].(map[string]interface{}), &DBConfig); err != nil {
		log.ErrorLog.Println(err.Error())
	}
	//set database url
	DBConfig.URL = fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=%s&parseTime=True&loc=Local",
		DBConfig.User, DBConfig.Password, DBConfig.Host, DBConfig.Port, DBConfig.Database, DBConfig.Charset)
}
func initHospitalDb() {
	if err := mapstructure.Decode(jsonData["hospitalDb"].(map[string]interface{}), &HospitalDBConfig); err != nil {
		log.ErrorLog.Println(err.Error())
	}
	//set database url
	HospitalDBConfig.URL = fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=%s&parseTime=True&loc=Local",
		HospitalDBConfig.User, HospitalDBConfig.Password, HospitalDBConfig.Host, HospitalDBConfig.Port, HospitalDBConfig.Database, HospitalDBConfig.Charset)
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
