package main

import (
	"github.com/astaxie/beego"
	_ "goblog/src/config"
	_ "goblog/src/filter"
	"goblog/src/logs"
	_ "goblog/src/routers"
)

func main() {

	defer logs.RecoverLog()

	beego.Run()

}
