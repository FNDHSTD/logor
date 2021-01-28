package logor

import (
	"fmt"
	"os"
	"time"
)

//FileLogger 往文件中记录日志的结构体
type FileLogger struct {
	level       Level
	logFilePath string
	logFileName string
	logFile     *os.File
	maxSize     int64
}

//NewFileLogger 是一个生成文件日志结构体的构造函数
func NewFileLogger(levelStr string, logFilePath, logFileName string) (fileLoggerPrinter *FileLogger) {
	loglevel := translateLogLevel(levelStr)
	fileLoggerPrinter = &FileLogger{
		level:       loglevel,
		logFileName: logFileName,
		logFilePath: logFilePath,
		maxSize:     10 * 1024 * 1024,
	}
	fileLoggerPrinter.initFileLogger()
	return
}

//用来初始化文件句柄
func (f *FileLogger) initFileLogger() {
	//生成日志文件路径
	filePath := fmt.Sprintf("%s%s", f.logFilePath, f.logFileName)
	//打开日志文件
	file, err := os.OpenFile(filePath, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)
	if err != nil {
		panic(fmt.Errorf("打开日志文件%s失败, %v", filePath, err))
	}
	f.logFile = file
	return
}

func (f *FileLogger) log(level Level, format string, args ...interface{}) {

	//如果设置的日志级别大于DEBUG级别就不记录
	if f.level > level {
		return
	}
	msg := fmt.Sprintf(format, args...)

	//生成当前时间
	now := time.Now().Format("2006-01-02 15:04:05.000")

	//生成调用文件，函数，行号信息
	fileName, funcName, line := getCallerInfo(3)

	//生成日志记录的信息
	logMsg := fmt.Sprintf("[%s][%s:%d][%s][%s] %s", now, fileName, line, funcName, getLevelStr(level), msg)

	//往文件里写之前检查文件大小是否超过maxSize
	fileInfo, _ := f.logFile.Stat()
	fileSize := fileInfo.Size()

	if fileSize >= f.maxSize {
		//切分文件
		fileName := f.logFile.Name() //拿到原文件的绝对路径

		backupName := fmt.Sprintf("%s_%v.back", fileName, time.Now().Unix())
		f.logFile.Close()
		os.Rename(fileName, backupName)
		file, err := os.OpenFile(fileName, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)
		if err != nil {
			panic(fmt.Errorf("打开日志文件失败"))
		}
		f.logFile = file
	}

	//将日志信息写入日志文件
	fmt.Fprintln(f.logFile, logMsg)
}

//Debug 记录日志
func (f *FileLogger) Debug(format string, args ...interface{}) {
	f.log(DEBUG, format, args...)
}

//Info 记录日志
func (f *FileLogger) Info(format string, args ...interface{}) {
	f.log(INFO, format, args...)
}

//Warn 记录日志
func (f *FileLogger) Warn(format string, args ...interface{}) {
	f.log(WARN, format, args...)
}

//Error 记录日志
func (f *FileLogger) Error(format string, args ...interface{}) {
	f.log(ERROR, format, args...)
}

//Fatal 记录日志
func (f *FileLogger) Fatal(format string, args ...interface{}) {
	f.log(FATAL, format, args...)
}

//Close 关闭日志文件
func (f *FileLogger) Close() {
	f.logFile.Close()
}
