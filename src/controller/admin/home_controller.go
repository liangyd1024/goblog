package admin

import (
	"goblog/src/controller"
)

type HomeController struct {
	controller.BaseController
}

func (home *HomeController) Home() {
	//home.TplName = "admin/home.html"
	//home.LayoutSections
	//home.Data["ctxPath"] = home.Ctx.Request.Site
	home.TplName = "admin/index.html"
}

func (home *HomeController) Main() {
	//home.Data["ctxPath"] = home.Ctx.Request.Site
	home.TplName = "admin/main.html"
}

func (home *HomeController) Test() {
	//home.Data["ctxPath"] = home.Ctx.Request.Site
	home.TplName = "admin/test.html"
}
