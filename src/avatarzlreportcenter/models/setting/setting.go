package setting

import (
	"github.com/go-ini/ini"
	"log"
	"time"
)

var (
	Cfg *ini.File

	RunModel     string
	HttpPort     int
	ReadTimeOut  time.Duration
	WriteTimeOit time.Duration
	//PageSize     int
)

func init() {
	var err error
	Cfg, err = ini.Load("config/app.ini")
	if err != nil {
		log.Fatalf("Fail to parse 'conf/app.ini': %v", err)
	}
	LoadBase()
	LoadServer()

}

func LoadBase() {
	RunModel = Cfg.Section("").Key("RUN_MODE").MustString("debug")
}

func LoadServer() {
	sec, err := Cfg.GetSection("server")
	if err != nil {
		log.Fatalf("Fail to get section server : %v", err)
	}
	HttpPort = sec.Key("HTTP_PORT").MustInt(8086)
	ReadTimeOut = time.Duration(sec.Key("READ_TIMEOUT").MustInt(60)) * time.Second
	WriteTimeOit = time.Duration(sec.Key("WRITE_TIMEOUT").MustInt(60)) * time.Second

}
