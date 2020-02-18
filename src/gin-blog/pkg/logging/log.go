package logging

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"runtime"
)

type Level int

var (
	F                  *os.File
	DefaultPrefix      = ""
	DefaultCallerDepth = 2
	logger             *log.Logger
	logPrefix          = ""
	levelFlags         = []string{"DEBUG", "INFO", "WARN", "ERROR", "FATAL"}
)

const (
	DEBUG Level = iota
	INFO
	WARNING
	ERROR
	FATAL
)

func init() {
	filePath := getLogFileFullPath()
	F = openLogFile(filePath)
	logger = log.New(F, DefaultPrefix, log.LstdFlags)
}

func Debug(v ...interface{}) {
	setPrefix(DEBUG)
	logger.Println(v)
}

func Warn(v ...interface{}) {
	setPrefix(WARNING)
	logger.Println(v)

}

func Error(v ...interface{}) {
	setPrefix(ERROR)
	logger.Println(v)
}

func Info(v ...interface{}) {
	setPrefix(INFO)
	fmt.Println(v)
	panic(v)
	logger.Println(v)
}

func Fata(v ...interface{}) {
	setPrefix(FATAL)
	logger.Fatalln(v)
}

//- `log.New`：创建一个新的日志记录器。`out`定义要写入日志数据的`IO`句柄。`prefix`定义每个生成的日志行的开头。`flag`定义了日志记录属性
func setPrefix(level Level) {
	_, file, line, ok := runtime.Caller(DefaultCallerDepth)
	if ok {
		logPrefix = fmt.Sprint("[%s][%s:%d]", levelFlags[level], filepath.Base(file), line)
	} else {
		logPrefix = fmt.Sprint("[%s]", levelFlags[level])
	}

	logger.SetPrefix(logPrefix)
}
