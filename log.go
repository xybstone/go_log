package log

import (
	"fmt"
	"os"
	"reflect"
	"strings"
	"time"
)

const (
	PREFIX      = "[ST Web Service]"
	TIME_FORMAT = "2006-01-02 15:04:05"
	DATE_FORMAT = "2006-01-02"
)

var (
	LEVEL_FLAGS = [...]string{"DEBUG", " INFO", " WARN", "ERROR", "FATAL"}
	PATH        = ""
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
	path := time.Now().Format(DATE_FORMAT) + ".txt"
	var f *os.File
	if exist(path) {
		f, _ = os.OpenFile(path, os.O_RDWR, os.ModeAppend)
	} else {
		f, _ = os.Create(path)
	}

	defer f.Close()
	f.WriteString(format)
}

func printf(format string, args ...interface{}) string {
	var res []string
	for _, arg := range args {
		t := reflect.TypeOf(arg)
		v := reflect.ValueOf(arg)
		var obj []string
		if t.Kind() == reflect.Struct {
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
	return strings.Join(res, "\n")
}

func Print(level int, format string, args ...interface{}) {
	info := fmt.Sprintf("%s %s %s %s\n",
		PREFIX, time.Now().Format(TIME_FORMAT), LEVEL_FLAGS[level],
		printf(format, args...))
	switch level {
	case DEBUG:
		fmt.Println(info)
	case INFO:
		fmt.Println(info)
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

func Debug(format string, args ...interface{}) {
	Print(DEBUG, format, args...)
}
func Warn(format string, args ...interface{}) {
	Print(WARNING, format, args...)
}
func Info(format string, args ...interface{}) {
	Print(INFO, format, args...)
}
func Error(format string, args ...interface{}) {
	Print(ERROR, format, args...)
}
func Fatal(format string, args ...interface{}) {
	Print(FATAL, format, args...)
}
