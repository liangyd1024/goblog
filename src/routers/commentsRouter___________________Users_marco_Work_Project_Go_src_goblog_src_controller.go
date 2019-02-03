package routers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context/param"
)

func init() {

	beego.GlobalControllerRouter["goblog/src/controller:ArticlesController"] = append(beego.GlobalControllerRouter["goblog/src/controller:ArticlesController"],
		beego.ControllerComments{
			Method: "ToArticles",
			Router: `/articles/:id`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["goblog/src/controller:ArticlesController"] = append(beego.GlobalControllerRouter["goblog/src/controller:ArticlesController"],
		beego.ControllerComments{
			Method: "Browse",
			Router: `/articles/browse`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["goblog/src/controller:ArticlesController"] = append(beego.GlobalControllerRouter["goblog/src/controller:ArticlesController"],
		beego.ControllerComments{
			Method: "ListCategory",
			Router: `/articles/category/list`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["goblog/src/controller:ArticlesController"] = append(beego.GlobalControllerRouter["goblog/src/controller:ArticlesController"],
		beego.ControllerComments{
			Method: "ListComment",
			Router: `/articles/comment/list`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["goblog/src/controller:ArticlesController"] = append(beego.GlobalControllerRouter["goblog/src/controller:ArticlesController"],
		beego.ControllerComments{
			Method: "PubComment",
			Router: `/articles/comment/pub`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["goblog/src/controller:ArticlesController"] = append(beego.GlobalControllerRouter["goblog/src/controller:ArticlesController"],
		beego.ControllerComments{
			Method: "ReplyComment",
			Router: `/articles/comment/reply`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["goblog/src/controller:ArticlesController"] = append(beego.GlobalControllerRouter["goblog/src/controller:ArticlesController"],
		beego.ControllerComments{
			Method: "GetArticles",
			Router: `/articles/get`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["goblog/src/controller:ArticlesController"] = append(beego.GlobalControllerRouter["goblog/src/controller:ArticlesController"],
		beego.ControllerComments{
			Method: "ListArticles",
			Router: `/articles/list`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["goblog/src/controller:ArticlesController"] = append(beego.GlobalControllerRouter["goblog/src/controller:ArticlesController"],
		beego.ControllerComments{
			Method: "ListPlaceOfFile",
			Router: `/articles/place-of-file/list`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["goblog/src/controller:ArticlesController"] = append(beego.GlobalControllerRouter["goblog/src/controller:ArticlesController"],
		beego.ControllerComments{
			Method: "Praise",
			Router: `/articles/praise`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["goblog/src/controller:ArticlesController"] = append(beego.GlobalControllerRouter["goblog/src/controller:ArticlesController"],
		beego.ControllerComments{
			Method: "ListRecommendArticles",
			Router: `/articles/recommend/list`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["goblog/src/controller:ArticlesController"] = append(beego.GlobalControllerRouter["goblog/src/controller:ArticlesController"],
		beego.ControllerComments{
			Method: "ListTag",
			Router: `/articles/tag/list`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["goblog/src/controller:ArticlesController"] = append(beego.GlobalControllerRouter["goblog/src/controller:ArticlesController"],
		beego.ControllerComments{
			Method: "ToSearch",
			Router: `/articles/tosearch`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["goblog/src/controller:ArticlesController"] = append(beego.GlobalControllerRouter["goblog/src/controller:ArticlesController"],
		beego.ControllerComments{
			Method: "TypeSearch",
			Router: `/articles/typesearch`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

}
