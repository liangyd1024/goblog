package routers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context/param"
)

func init() {

	beego.GlobalControllerRouter["goblog/src/controller/admin:BowenController"] = append(beego.GlobalControllerRouter["goblog/src/controller/admin:BowenController"],
		beego.ControllerComments{
			Method: "DeleteArticlesCategory",
			Router: `/admin/bowen/category/delete`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["goblog/src/controller/admin:BowenController"] = append(beego.GlobalControllerRouter["goblog/src/controller/admin:BowenController"],
		beego.ControllerComments{
			Method: "CollectStatus",
			Router: `/admin/bowen/collect/status`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["goblog/src/controller/admin:BowenController"] = append(beego.GlobalControllerRouter["goblog/src/controller/admin:BowenController"],
		beego.ControllerComments{
			Method: "CollectType",
			Router: `/admin/bowen/collect/type`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["goblog/src/controller/admin:BowenController"] = append(beego.GlobalControllerRouter["goblog/src/controller/admin:BowenController"],
		beego.ControllerComments{
			Method: "ListComment",
			Router: `/admin/bowen/comment/list`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["goblog/src/controller/admin:BowenController"] = append(beego.GlobalControllerRouter["goblog/src/controller/admin:BowenController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/admin/bowen/delete`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["goblog/src/controller/admin:BowenController"] = append(beego.GlobalControllerRouter["goblog/src/controller/admin:BowenController"],
		beego.ControllerComments{
			Method: "Forward",
			Router: `/admin/bowen/forward`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["goblog/src/controller/admin:BowenController"] = append(beego.GlobalControllerRouter["goblog/src/controller/admin:BowenController"],
		beego.ControllerComments{
			Method: "GetBowen",
			Router: `/admin/bowen/get`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["goblog/src/controller/admin:BowenController"] = append(beego.GlobalControllerRouter["goblog/src/controller/admin:BowenController"],
		beego.ControllerComments{
			Method: "List",
			Router: `/admin/bowen/list`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["goblog/src/controller/admin:BowenController"] = append(beego.GlobalControllerRouter["goblog/src/controller/admin:BowenController"],
		beego.ControllerComments{
			Method: "ModifyArticles",
			Router: `/admin/bowen/modify`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["goblog/src/controller/admin:BowenController"] = append(beego.GlobalControllerRouter["goblog/src/controller/admin:BowenController"],
		beego.ControllerComments{
			Method: "ModifyStatus",
			Router: `/admin/bowen/modifystatus`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["goblog/src/controller/admin:BowenController"] = append(beego.GlobalControllerRouter["goblog/src/controller/admin:BowenController"],
		beego.ControllerComments{
			Method: "Publish",
			Router: `/admin/bowen/publish`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["goblog/src/controller/admin:BowenController"] = append(beego.GlobalControllerRouter["goblog/src/controller/admin:BowenController"],
		beego.ControllerComments{
			Method: "DeleteArticlesTag",
			Router: `/admin/bowen/tag/delete`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["goblog/src/controller/admin:BowenController"] = append(beego.GlobalControllerRouter["goblog/src/controller/admin:BowenController"],
		beego.ControllerComments{
			Method: "ToDetails",
			Router: `/admin/bowen/todetails`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["goblog/src/controller/admin:BowenController"] = append(beego.GlobalControllerRouter["goblog/src/controller/admin:BowenController"],
		beego.ControllerComments{
			Method: "ToModify",
			Router: `/admin/bowen/tomodify`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["goblog/src/controller/admin:BowenController"] = append(beego.GlobalControllerRouter["goblog/src/controller/admin:BowenController"],
		beego.ControllerComments{
			Method: "ToPublish",
			Router: `/admin/bowen/topublish`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["goblog/src/controller/admin:CategoryController"] = append(beego.GlobalControllerRouter["goblog/src/controller/admin:CategoryController"],
		beego.ControllerComments{
			Method: "Add",
			Router: `/admin/category/add`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["goblog/src/controller/admin:CategoryController"] = append(beego.GlobalControllerRouter["goblog/src/controller/admin:CategoryController"],
		beego.ControllerComments{
			Method: "BowenList",
			Router: `/admin/category/bowen/list`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["goblog/src/controller/admin:CategoryController"] = append(beego.GlobalControllerRouter["goblog/src/controller/admin:CategoryController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/admin/category/delete`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["goblog/src/controller/admin:CategoryController"] = append(beego.GlobalControllerRouter["goblog/src/controller/admin:CategoryController"],
		beego.ControllerComments{
			Method: "List",
			Router: `/admin/category/list`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["goblog/src/controller/admin:FileUploadController"] = append(beego.GlobalControllerRouter["goblog/src/controller/admin:FileUploadController"],
		beego.ControllerComments{
			Method: "UploadEditorImage",
			Router: `/admin/file/upload/editor`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["goblog/src/controller/admin:SysController"] = append(beego.GlobalControllerRouter["goblog/src/controller/admin:SysController"],
		beego.ControllerComments{
			Method: "RefreshIndexer",
			Router: `/admin/sys/refresh/indexer`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["goblog/src/controller/admin:TagController"] = append(beego.GlobalControllerRouter["goblog/src/controller/admin:TagController"],
		beego.ControllerComments{
			Method: "Add",
			Router: `/admin/tag/add`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["goblog/src/controller/admin:TagController"] = append(beego.GlobalControllerRouter["goblog/src/controller/admin:TagController"],
		beego.ControllerComments{
			Method: "BowenList",
			Router: `/admin/tag/bowen/list`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["goblog/src/controller/admin:TagController"] = append(beego.GlobalControllerRouter["goblog/src/controller/admin:TagController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/admin/tag/delete`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["goblog/src/controller/admin:TagController"] = append(beego.GlobalControllerRouter["goblog/src/controller/admin:TagController"],
		beego.ControllerComments{
			Method: "List",
			Router: `/admin/tag/list`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["goblog/src/controller/admin:UserController"] = append(beego.GlobalControllerRouter["goblog/src/controller/admin:UserController"],
		beego.ControllerComments{
			Method: "ToLogin",
			Router: `/admin/login`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["goblog/src/controller/admin:UserController"] = append(beego.GlobalControllerRouter["goblog/src/controller/admin:UserController"],
		beego.ControllerComments{
			Method: "Login",
			Router: `/admin/user/login`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["goblog/src/controller/admin:UserController"] = append(beego.GlobalControllerRouter["goblog/src/controller/admin:UserController"],
		beego.ControllerComments{
			Method: "Captcha",
			Router: `/admin/user/login/captcha`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["goblog/src/controller/admin:UserController"] = append(beego.GlobalControllerRouter["goblog/src/controller/admin:UserController"],
		beego.ControllerComments{
			Method: "Logout",
			Router: `/admin/user/logout`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

}
