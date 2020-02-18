package log

import (
	"log"
	"os"
)

var (
	InfoLog  *log.Logger
	ErrorLog *log.Logger
)

func init() {
	GOPATH := os.Getenv("GOPATH")
	dname := GOPATH + "/src/wx_mini/storage/log"
	os.MkdirAll(dname, os.ModeDir|os.ModePerm)
	errFile, err := os.OpenFile(dname+"/errors.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, os.ModePerm)

	if err != nil {
		log.Fatalln("cannot open log File",err)
	}

	infoFile ,err := os.OpenFile(dname+"/info.log",os.O_CREATE|os.O_WRONLY|os.O_APPEND, os.ModePerm)
	if err != nil {
		log.Fatalln("cannot open log File",err)
	}
	InfoLog = log.New(infoFile, "Info:", log.Ldate|log.Ltime|log.Lshortfile)
	ErrorLog = log.New(errFile, "Error:", log.Ldate|log.Ltime|log.Lshortfile)

}
