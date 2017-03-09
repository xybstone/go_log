package log

import (
	"fmt"
	golog "log"
	"os"
	"reflect"
	"strings"
	"time"
)

const (
	//PREFIX 前缀
	PREFIX = "[ST LOG]"
	//TIME_FORMAT 日期格式
	TIME_FORMAT = "2006-01-02 15:04:05"
	//DATE_FORMAT 日期格式
	DATE_FORMAT = "2006-01-02"
)

var (
	LEVEL_FLAGS = [...]string{"DEBUG", " INFO", " WARN", "ERROR", "FATAL"}
	PATH        = ""
	IsDebug     = false
)

const (
	DEBUG = iota
	INFO
	WARNING
	ERROR
	FATAL
)

func exist(filename string) bool {
	_, err := os.Stat(filename)
	return err == nil || os.IsExist(err)
}

func saveLog(format string) {

	path := "logs/" + time.Now().Format(DATE_FORMAT) + ".txt"
	var f *os.File
	if !exist(path) {
		f, _ = os.Create(path)
	} else {
		f, _ = os.OpenFile(path, os.O_CREATE|os.O_RDWR|os.O_APPEND, 0666)
	}
	golog.SetOutput(f)
	//golog.SetFlags(golog.Ldate | golog.Ltime | golog.Lshortfile)
	golog.Output(0, format)
}

func printf(format string, args ...interface{}) string {
	var res []string
	for _, arg := range args {
		t := reflect.TypeOf(arg)
		v := reflect.ValueOf(arg)
		if t.Kind() == reflect.String {
			res = append(res, fmt.Sprintf("%s", arg))
		} else if t.Kind() == reflect.Struct {
			var obj []string
			obj = append(obj, fmt.Sprintf("Object Name:%s", t.Name()))
			for i := 0; i < t.NumField(); i++ {
				f := t.Field(i)
				//此时的 Field 必须为public
				val := v.Field(i).Interface()
				if valStr := fmt.Sprintf("%v", val); valStr != "" {
					obj = append(obj, fmt.Sprintf("%s:%v", f.Name, val))
				}
			}
			res = append(res, strings.Join(obj, "\n"))
		} else {
			res = append(res, fmt.Sprintf("%v", arg))
		}
	}
	return fmt.Sprintln(format, strings.Join(res, "\n"))
}

//Println 打印输出
func Println(a ...interface{}) {
	saveLog(fmt.Sprintf("%s %s %s\n", PREFIX, LEVEL_FLAGS[INFO], fmt.Sprintln(a)))
}

//Print 打印对象，对象属性会被依次打印
func Print(level int, format string, args ...interface{}) {
	info := fmt.Sprintf("%s %s %s\n",
		PREFIX, LEVEL_FLAGS[level],
		printf(format, args...))
	switch level {
	case DEBUG:
		fmt.Println(info)
	case INFO:
		fmt.Println(info)
		saveLog(info)
	case WARNING:
		fmt.Println(info)
	case ERROR:
		fmt.Println(info)
	case FATAL:
		fmt.Println(info)
		saveLog(info)
	default:
		fmt.Println(info)
	}
}

//Debug 调试
func Debug(format string, args ...interface{}) {
	if IsDebug {
		Print(DEBUG, format, args...)
	}
}

//Warn 警告
func Warn(format string, args ...interface{}) {
	Print(WARNING, format, args...)
}

//Info 信息，持久化
func Info(format string, args ...interface{}) {
	Print(INFO, format, args...)
}

//Error 醋味
func Error(format string, args ...interface{}) {
	Print(ERROR, format, args...)
}

//Fatal 异常
func Fatal(format string, args ...interface{}) {
	Print(FATAL, format, args...)
}
