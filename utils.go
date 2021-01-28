package logor

import (
	"path"
	"runtime"
)

//存放一些公用的工具函数

//获得调用的文件名、行号、函数名
func getCallerInfo(skip int) (fileName, funcName string, line int) {
	pc, fileName, line, ok := runtime.Caller(skip)
	if !ok {
		return
	}
	//从fileName路径中剥离出文件名
	fileName = path.Base(fileName)
	//根据pc拿到函数名
	funcName = path.Base(runtime.FuncForPC(pc).Name())
	return
}
