package main

import (
	"fmt"
	"github.com/xxl6097/go-glog/glog"
	"time"
)

func main() {
	fmt.Println("aaa")
	glog.LogSaveFile()

	for {
		glog.Info("fajswehfasef")
		glog.Error("fajswehfasef")
		glog.Warn("fajswehfasef")
		glog.Debug("fajswehfasef")
		time.Sleep(time.Second)
	}
}
