# go-glog的使用方式

```shell
git add .
git commit -m "release v0.0.6"
git tag -a v0.0.6 -m "release v0.0.6"
git push origin v0.0.6
```
## 一、添加依赖

```go
go get -u github.com/xxl6097/go-glog@v0.0.5
```

## 二、使用示例
```go

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
		time.Sleep(time.Second)
	}
}

```


