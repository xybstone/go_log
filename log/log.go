package log
import (
	"fmt"
	"os"
	"time"
)
const (
	PREFIX      = "[ST Web Service]"
	TIME_FORMAT = "06-01-02 15:04:05"
)
var (
	LEVEL_FLAGS = [...]string{"DEBUG", " INFO", " WARN", "ERROR", "FATAL"}
)
const (
	DEBUG = iota
	INFO
	WARNING
	ERROR
	FATAL
)
func Print(level int, format string, args ...interface{}) {
	switch level {
	case DEBUG:
		fmt.Printf("%s %s %s %s\n",
			PREFIX, time.Now().Format(TIME_FORMAT), LEVEL_FLAGS[level],
			fmt.Sprintf(format, args...))
	case INFO:
		fmt.Printf("%s %s %s %s\n",
			PREFIX, time.Now().Format(TIME_FORMAT), LEVEL_FLAGS[level],
			fmt.Sprintf(format, args...))
	case WARNING:
		fmt.Printf("%s %s %s %s\n",
			PREFIX, time.Now().Format(TIME_FORMAT), LEVEL_FLAGS[level],
			fmt.Sprintf(format, args...))
	case ERROR:
		fmt.Printf("%s %s %s %s\n",
			PREFIX, time.Now().Format(TIME_FORMAT), LEVEL_FLAGS[level],
			fmt.Sprintf(format, args...))
	case FATAL:
		fmt.Printf("%s %s %s %s\n",
			PREFIX, time.Now().Format(TIME_FORMAT), LEVEL_FLAGS[level],
			fmt.Sprintf(format, args...))
		os.Exit(1)
	default:
		fmt.Printf("%s %s %s %s\n",
			PREFIX, time.Now().Format(TIME_FORMAT), LEVEL_FLAGS[level],
			fmt.Sprintf(format, args...))
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
