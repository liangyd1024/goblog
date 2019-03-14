//AppName: goblog
//Version: V1.0.0
//User: marco
//Date: 2019-01-31

package crypt

import (
	"fmt"
	. "goblog/src/logs"
	"runtime"
	"testing"
)

func TestGetMd5(t *testing.T) {
	fmt.Println(GetMd5("goblog"))
	InitLogs(false)
	go func() {
		defer RecoverLog()
		fmt.Println("test")
		panic("恐慌了")
	}()
	runtime.Gosched()
	panic("结束运行")
}
