package controller

import (
	"github.com/astaxie/beego"
	. "goblog/src/logs"
	"goblog/src/model"
	"goblog/src/utils/bizerror"
	"goblog/src/utils/constant"
	"goblog/src/utils/dataconv"
	"reflect"
	"runtime/debug"
	"strconv"
)

type BaseController struct {
	beego.Controller
}

//获取当前应用路径
func (base *BaseController) Site(path string) string {
	site := base.Ctx.Input.Site() + ":" + strconv.Itoa(base.Ctx.Input.Port())
	if path != "" {
		site = site + "/" + path
	}
	return site
}

//解析请求体中json数据
func (base *BaseController) ParseJson(body interface{}) interface{} {
	return dataconv.JsonByte2M(base.Ctx.Input.RequestBody, body)
}

//设置当前用户会话
func (base *BaseController) SetSessionUser(user *model.User) {
	base.SetSession("user", user)
}

//获取当前用户会话
func (base *BaseController) GetSessionUser() *model.User {
	user := base.GetSession("user")
	if user != nil{
		return user.(*model.User)
	}
	return nil
}

//清除当前用户会话
func (base *BaseController) ClearSessionUser() {
	if base.GetSessionUser() != nil {
		base.DelSession("user")
	}
}

//获取系统当前用户
func (base *BaseController) GetUserName() string {
	user := base.GetSessionUser()
	if user != nil {
		return user.UserName
	}
	return constant.SYS
}

//异常处理
func (base *BaseController) PanicHandler() {
	if err := recover(); err != nil {
		Log.Printf("call PanicHandler errInfo:%v,err_type:%v,stack:%v", err, reflect.TypeOf(err),string(debug.Stack()))
		bizError := bizerror.BizError500100
		err, ok := err.(bizerror.BizError)
		if ok {
			bizError = err
		}
		base.SysErrResponse(bizError)
	}
}

//构建业务异常
func (base *BaseController) BuildBizErrorHandler(bizError bizerror.BizError) {
	base.BuildErrResponse(bizError.ErrCode, bizError.ErrMsg)
}

//构建异常响应
func (base *BaseController) BuildErrResponse(code, msg string) {
	resp := new(model.Response)
	response := resp.FailAll(code, msg)
	base.Data["json"] = response
	base.ServeJSON()
	Log.Printf("call BuildErrResponse response:%+v", response)
}

//构建系统异常响应
func (base *BaseController) SysErrResponse(bizError bizerror.BizError) {
	base.BuildErrResponse(bizError.ErrCode, bizError.ErrMsg)
}

//构建分页成功响应
func (base *BaseController) BuildSucPagingResponse(result interface{}, paging model.Paging) {
	resp := new(model.Response)
	response := resp.SucPage(result, paging)
	base.Data["json"] = response
	base.ServeJSON()
	Log.Printf("call BuildSucResponse url:%+v,response:%+v", base.Ctx.Request.URL, response)
}

//构建成功响应
func (base *BaseController) BuildSucResponse(result interface{}) {
	resp := new(model.Response)
	response := resp.Suc(result)
	base.Data["json"] = response
	base.ServeJSON()
	Log.Printf("call BuildSucResponse url:%+v,response:%+v", base.Ctx.Request.URL, response)
}
