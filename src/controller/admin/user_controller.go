//AppName: goblog
//Version: V1.0.0
//User: marco
//Date: 2018/12/11

package admin

import (
	"github.com/astaxie/beego"
	"goblog/src/controller"
	. "goblog/src/logs"
	"goblog/src/model"
	"goblog/src/service"
	"goblog/src/utils/bizerror"
	"goblog/src/utils/constant"
	"goblog/src/utils/crypt"
)

type UserController struct {
	controller.BaseController
}

//初始化用户
func init() {
	user := new(model.User)
	user.UserName = beego.AppConfig.DefaultString("userName", "admin")
	user.UserPwd = beego.AppConfig.DefaultString("userPwd", "goblog")
	user.NickName = constant.SYS
	user.UserType = constant.SYS
	user.Status = constant.USER_NORMAL_STATUS
	user.CreateBy = constant.SYS
	service.UserBiz.CreateUser(user)
}

func (userController *UserController) URLMapping() {
	userController.Mapping("ToLogin", userController.ToLogin)
	userController.Mapping("Captcha", userController.Captcha)
	userController.Mapping("Login", userController.Login)
	userController.Mapping("Logout", userController.Logout)
}

// @router /admin/login [get]
func (userController *UserController) ToLogin() {
	userController.TplName = "admin/user/login.html"
}

// @router /admin/user/login/captcha [get]
func (userController *UserController) Captcha() {
	defer userController.PanicHandler()

	uid, image := crypt.GenerateCaptcha()
	userController.SetSession("CaptchaUid", uid)

	userController.BuildSucResponse(image)
}

// @router /admin/user/login [post]
func (userController *UserController) Login() {
	defer userController.PanicHandler()

	user := &model.User{}
	userController.ParseJson(user)

	Log.Info("call Login user:%v", user)
	captchaUid := userController.GetSession("CaptchaUid")
	if captchaUid == nil || !crypt.VerifyCaptcha(captchaUid.(string), user.Captcha) {
		bizerror.BizError400102.PanicError()
	}

	user = service.UserBiz.CheckPwd(user)
	userController.SetSessionUser(user)

	userController.BuildSucResponse(true)
}

// @router /admin/user/logout [get]
func (userController *UserController) Logout() {
	defer userController.PanicHandler()

	userController.ClearSessionUser()

	userController.BuildSucResponse(true)
}
