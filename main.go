package main

import (
	"fmt"
	"github.com/xxl6097/go-glog/glog"
	"os"
	"path"
	"runtime"
	"time"
)

func hook(data []byte) {
	fmt.Println(string(data))
}

func init() {
	//开启日志保存文件
	//glog.LogSaveFile()
	//glog.SetNoHeader(true)
	//拦截日志
	//glog.Hook(hook)
	//glog.SetCons(true)
	//glog.SetNoHeader(true)
	//glog.SetNoColor(true)

	//glog.SetMaxSize(1 * 1024 * 1024)
	//glog.SetMaxAge(15)
	//glog.SetCons(true)
	//glog.SetNoHeader(true)
	//glog.SetNoColor(true)
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
	//glog.SetLogFile("/usr/local/AATEST/logs", "normal.log")
	glog.Info("只有使用这个log打印才能记录日志哦", time.Now().Format("2006-01-02 15:04:05"))
	//glog.Flush()
	//testlog()
	//for {
	//	glog.Info("只有使用这个log打印才能记录日志哦", time.Now().Format("2006-01-02 15:04:05"))
	//	//glog.Info("Info。。。。")
	//	//glog.Error("Error。。。。")
	//	//glog.Warn("Warn。。。")
	//	//glog.Debug("Debug。。。")
	//	time.Sleep(5 * time.Second)
	//}
	fmt.Scanln()
	os.Exit(0)
}
