package logor

import (
	"fmt"
	"time"
)

//往终端打印日志

//ConsoleLogger 是一个往终端写日志的结构体
type ConsoleLogger struct {
	level Level
}

//NewConsoleLogger 是一个生成文件日志结构体的构造函数
func NewConsoleLogger(levelStr string) (consoleLoggerPrinter *ConsoleLogger) {
	loglevel := translateLogLevel(levelStr)
	consoleLoggerPrinter = &ConsoleLogger{
		level: loglevel,
	}
	return
}

func (c *ConsoleLogger) log(level Level, format string, args ...interface{}) {

	//如果设置的日志级别大于DEBUG级别就不记录
	if c.level > level {
		return
	}
	msg := fmt.Sprintf(format, args...)

	//生成当前时间
	now := time.Now().Format("2006-01-02 15:04:05.000")

	//生成调用文件，函数，行号信息
	fileName, funcName, line := getCallerInfo(3)

	//生成日志记录的信息
	logMsg := fmt.Sprintf("[%s][%s:%d][%s][%s] %s", now, fileName, line, funcName, getLevelStr(level), msg)

	//将日志信息打印到终端
	fmt.Println(logMsg)
}

//Debug 记录日志
func (c *ConsoleLogger) Debug(format string, args ...interface{}) {
	c.log(DEBUG, format, args...)
}

//Info 记录日志
func (c *ConsoleLogger) Info(format string, args ...interface{}) {
	c.log(INFO, format, args...)
}

//Warn 记录日志
func (c *ConsoleLogger) Warn(format string, args ...interface{}) {
	c.log(WARN, format, args...)
}

//Error 记录日志
func (c *ConsoleLogger) Error(format string, args ...interface{}) {
	c.log(ERROR, format, args...)
}

//Fatal 记录日志
func (c *ConsoleLogger) Fatal(format string, args ...interface{}) {
	c.log(FATAL, format, args...)
}

// Close console不需要关闭
func (c *ConsoleLogger) Close() {
	return
}
