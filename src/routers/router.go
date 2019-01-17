package routers

import (
	"github.com/astaxie/beego"
	"goblog/src/controller"
	"goblog/src/controller/admin"
)

func init() {
	//编程式路由
	beego.Include(
		&admin.SysController{},
		&admin.BowenController{},
		&admin.TagController{},
		&admin.CategoryController{},
		&admin.FileUploadController{},
		&admin.UserController{},

		&controller.ArticlesController{},
	)

	//声明式路由
	beego.Router("/", &controller.IndexController{}, "get,post:Index")
	beego.Router("/test", &controller.IndexController{}, "get,post:Test")
	beego.Router("/details", &controller.IndexController{}, "get,get:Details")
	beego.Router("/about", &controller.IndexController{}, "get,get:About")
	beego.Router("/contact", &controller.IndexController{}, "get,get:Contact")
	beego.Router("/admin", &admin.HomeController{}, "get,post:Home")
	beego.Router("/admin/main", &admin.HomeController{}, "get,post:Main")
	beego.Router("/admin/test", &admin.HomeController{}, "get,post:Test")

	//404错误处理
	beego.ErrorController(&controller.ErrorController{})
}
