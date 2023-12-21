package main

import (
	"fmt"
	"github.com/xxl6097/go-glog/glog"
	"time"
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

func main() {
	for {
		fmt.Println("aaa")
		glog.Info("只有使用这个log打印才能记录日志哦")
		glog.Info("Info。。。。")
		glog.Error("Error。。。。")
		glog.Warn("Warn。。。")
		glog.Debug("Debug。。。")
		time.Sleep(time.Second)
	}
}
