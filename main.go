package main

import (
	"github.com/astaxie/beego"
	_ "goblog/src/config"
	_ "goblog/src/filter"
	_ "goblog/src/routers"
)

func main() {

	beego.Run()

}
