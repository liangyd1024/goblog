package filter

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context"
	. "goblog/src/logs"
	"goblog/src/model"
)

func init() {
	beego.InsertFilter("/*", beego.BeforeExec, LogFilter)
	beego.InsertFilter("/admin/*", beego.BeforeRouter, LoginFilter)
}

var (
	LOGIN_FILTER_URL_MAP = map[string]byte{
		"/admin/login":              1,
		"/admin/user/login/captcha": 1,
		"/admin/user/login":         1,
	}
)

//登陆过滤器
func LoginFilter(ctx *context.Context) {
	user, ok := ctx.Input.Session("user").(*model.User)
	Log.Info("call LoginFilter url:%v,ok:%v,user:%v", ctx.Request.RequestURI, ok, user)
	if (!ok || user == nil) && LOGIN_FILTER_URL_MAP[ctx.Request.RequestURI] == 0 {
		ctx.Redirect(302, "/admin/login")
	}
}

//日志过滤器
func LogFilter(ctx *context.Context) {
	Log.Info("call LogFilter url:%s,form:%s,requestBody:%+v", ctx.Request.URL, ctx.Request.Form, string(ctx.Input.RequestBody[:]))
}
