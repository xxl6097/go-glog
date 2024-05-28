package glog

import (
	"fmt"
	"os"
)

/*
	A global Log handle is provided by default for external use, which can be called directly through the API series.
	The global log object is StdGLog.
	Note: The methods in this file do not support customization and cannot replace the log recording mode.

	If you need a custom logger, please use the following methods:
	zlog.SetLogger(yourLogger)
	zlog.Ins().InfoF() and other methods.

   全局默认提供一个Log对外句柄，可以直接使用API系列调用
   全局日志对象 StdGLog
   注意：本文件方法不支持自定义，无法替换日志记录模式，如果需要自定义Logger:

   请使用如下方法:
   zlog.SetLogger(yourLogger)
   zlog.Ins().InfoF()等方法
*/

var StdGLog = NewGLog(os.Stdout, "", BitDefault)

func Flags() int {
	return StdGLog.Flags()
}

// ResetFlags sets the flags of StdGLog
func ResetFlags(flag int) {
	StdGLog.ResetFlags(flag)
}

func AddFlag(flag int) {
	StdGLog.AddFlag(flag)
}

func SetPrefix(prefix string) {
	StdGLog.SetPrefix(prefix)
}

func SetLogFile(fileDir string, fileName string) {
	StdGLog.SetLogFile(fileDir, fileName)
}

func LogSaveFile() {
	StdGLog.SetLogFile("./", "app.log")
	StdGLog.SetCons(true)
}

// Hook hook log
func Hook(f func([]byte)) {
	StdGLog.SetLogHook(f)
}

// SetMaxAge 最大保留天数
func SetMaxAge(ma int) {
	StdGLog.SetMaxAge(ma)
}

// SetMaxSize 单个日志最大容量 单位：字节
func SetMaxSize(ms int64) {
	StdGLog.SetMaxSize(ms)
}

// SetCons 同时输出控制台
func SetCons(b bool) {
	StdGLog.SetCons(b)
}

func SetLogLevel(logLevel int) {
	StdGLog.SetLogLevel(logLevel)
}

func Debugf(format string, v ...interface{}) {
	StdGLog.Debugf(format, v...)
}

func Debug(v ...interface{}) {
	StdGLog.Debug(v...)
}

func Println(a ...any) {
	StdGLog.Info(a...)
}

func Sprintf(format string, a ...any) string {
	return fmt.Sprintf(format, a...)
}

func Print(a ...any) {
	StdGLog.Info(a...)
}

func Printf(format string, a ...any) {
	StdGLog.Infof(format, a...)
}

func Infof(format string, v ...interface{}) {
	StdGLog.Infof(format, v...)
}

func Info(v ...interface{}) {
	StdGLog.Info(v...)
}

func Warnf(format string, v ...interface{}) {
	StdGLog.Warnf(format, v...)
}

func Warn(v ...interface{}) {
	StdGLog.Warn(v...)
}

func Errorf(format string, v ...interface{}) {
	StdGLog.Errorf(format, v...)
}

func Error(v ...interface{}) {
	StdGLog.Error(v...)
}

func Fatalf(format string, v ...interface{}) {
	StdGLog.Fatalf(format, v...)
}

func Fatal(v ...interface{}) {
	StdGLog.Fatal(v...)
}

func Panicf(format string, v ...interface{}) {
	StdGLog.Panicf(format, v...)
}

func Panic(v ...interface{}) {
	StdGLog.Panic(v...)
}

func Stack(v ...interface{}) {
	StdGLog.Stack(v...)
}

func init() {
	// (因为StdGLog对象 对所有输出方法做了一层包裹，所以在打印调用函数的时候，比正常的logger对象多一层调用
	// 一般的gLogger对象 calldDepth=2, StdGLog的calldDepth=3)
	StdGLog.calldDepth = 3
}
