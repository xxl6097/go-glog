package main

import (
	"fmt"
	"github.com/xxl6097/go-glog/glog"
	"path"
	"runtime"
)

func hook(data []byte) {
	fmt.Println(string(data))
}

func init() {
	//开启日志保存文件
	glog.LogSaveFile()
	//拦截日志
	//glog.Hook(hook)
}
func getCallerInfo(skip int) (info string) {

	pc, file, lineNo, ok := runtime.Caller(skip)
	if !ok {

		info = "runtime.Caller() failed"
		return
	}
	funcName := runtime.FuncForPC(pc).Name()
	fileName := path.Base(file) // Base函数返回路径的最后一个元素
	return fmt.Sprintf("FuncName:%s, file:%s, line:%d ", funcName, fileName, lineNo)
}

func testlog() {
	// 打印出getCallerInfo函数自身的信息
	//fmt.Println(getCallerInfo(0))
	// 打印出getCallerInfo函数的调用者的信息
	fmt.Println(getCallerInfo(1))
}

func main() {
	glog.Println("hello glog...")
	glog.Flush()
	//testlog()
	//for {
	//	fmt.Println("aaa")
	//	glog.Info("只有使用这个log打印才能记录日志哦")
	//	glog.Info("Info。。。。")
	//	glog.Error("Error。。。。")
	//	glog.Warn("Warn。。。")
	//	glog.Debug("Debug。。。")
	//	time.Sleep(5 * time.Second)
	//}
}
