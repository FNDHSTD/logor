package mylogger

import "strings"

//自定义日志库,实现日志记录的功能

//日志分级别
//DEBUG  INFO WARING ERROR

//Level 是一个日志的自定义的日志级别
type Level uint8

// Logger 定义一个logger接口
type Logger interface {
	Debug(format string, args ...interface{})
	Info(format string, args ...interface{})
	Warn(format string, args ...interface{})
	Error(format string, args ...interface{})
	Fatal(format string, args ...interface{})
	Close()
}

//定义具体的日志级别
const (
	DEBUG Level = iota
	INFO
	WARN
	ERROR
	FATAL
)

// 写一个格局传进来的Level获取对应的字符串
func getLevelStr(level Level) (levelStr string) {
	switch level {
	case DEBUG:
		return "DEBUG"
	case INFO:
		return "INFO"
	case WARN:
		return "WARN"
	case ERROR:
		return "ERROR"
	case FATAL:
		return "FATAL"
	default:
		return "DEBUG"
	}

}

// 将用户传入的字符串类型的日志级别，解析出对应的Level
func translateLogLevel(levelStr string) (level Level) {
	levelStr = strings.ToLower(levelStr)
	switch levelStr {
	case "debug":
		return DEBUG
	case "info":
		return INFO
	case "warn":
		return WARN
	case "error":
		return ERROR
	case "fatal":
		return FATAL
	default:
		return DEBUG
	}
}
