package admin

import (
	"goblog/src/controller"
	"goblog/src/model"
	"goblog/src/service"
	"goblog/src/utils/bizerror"
)

type BowenController struct {
	controller.BaseController
}

func (bowen *BowenController) URLMapping() {
	bowen.Mapping("Forward", bowen.Forward)
	bowen.Mapping("ToPublish", bowen.ToPublish)
	bowen.Mapping("ToDetails", bowen.ToDetails)
	bowen.Mapping("ToModify", bowen.ToModify)

	bowen.Mapping("Publish", bowen.Publish)
	bowen.Mapping("GetBowen", bowen.GetBowen)
	bowen.Mapping("ListComment", bowen.ListComment)
	bowen.Mapping("List", bowen.List)
	bowen.Mapping("ModifyArticles", bowen.ModifyArticles)
	bowen.Mapping("ModifyStatus", bowen.ModifyStatus)
	bowen.Mapping("DeleteArticles", bowen.Delete)
	bowen.Mapping("DeleteArticlesTag", bowen.DeleteArticlesTag)
	bowen.Mapping("DeleteArticlesCategory", bowen.DeleteArticlesCategory)

	bowen.Mapping("CollectStatus", bowen.CollectStatus)
	bowen.Mapping("CollectType", bowen.CollectType)
}

// @router /admin/bowen/forward [get]
func (bowen *BowenController) Forward() {
	bowen.TplName = "admin/bowen/list.html"
}

// @router /admin/bowen/topublish [get]
func (bowen *BowenController) ToPublish() {
	bowen.TplName = "admin/bowen/publish.html"
}

// @router /admin/bowen/todetails [get]
func (bowen *BowenController) ToDetails() {
	id, err := bowen.GetInt("id")
	bizerror.Check(err)
	bowen.Data["id"] = id
	bowen.TplName = "admin/bowen/details.html"
}

// @router /admin/bowen/tomodify [get]
func (bowen *BowenController) ToModify() {
	defer bowen.PanicHandler()

	id, err := bowen.GetInt("id")
	bizerror.Check(err)
	bowen.Data["id"] = id
	bowen.TplName = "admin/bowen/publish.html"
}

// @router /admin/bowen/publish [post]
func (bowen *BowenController) Publish() {
	defer bowen.PanicHandler()

	articles := new(model.Articles)
	bowen.ParseJson(articles)

	articles.CreateBy = bowen.GetUserName()
	articles.Publisher = bowen.GetUserName()
	articles.ArticlesDetails.CreateBy = bowen.GetUserName()

	service.BowenBiz.Publish(articles)

	bowen.BuildSucResponse(articles)
}

// @router /admin/bowen/get [post]
func (bowen *BowenController) GetBowen() {
	defer bowen.PanicHandler()

	articles := new(model.Articles)
	bowen.ParseJson(articles)

	service.BowenBiz.GetBowen(articles)

	bowen.BuildSucResponse(articles)
}

// @router /admin/bowen/comment/list [post]
func (bowenController *BowenController) ListComment() {
	defer bowenController.PanicHandler()

	comment := new(model.Comment)
	bowenController.ParseJson(comment)

	bowenController.BuildSucPagingResponse(service.BowenBiz.ListComment(comment), comment.Paging)
}

// @router /admin/bowen/list [post]
func (bowen *BowenController) List() {
	defer bowen.PanicHandler()

	articles := new(model.Articles)
	bowen.ParseJson(articles)

	bowen.BuildSucPagingResponse(service.BowenBiz.GetBowenCondition(articles), articles.Paging)
}

// @router /admin/bowen/modify [post]
func (bowen *BowenController) ModifyArticles() {
	defer bowen.PanicHandler()

	articles := new(model.Articles)
	bowen.ParseJson(articles)

	articles.UpdateBy = bowen.GetUserName()
	service.BowenBiz.ModifyArticles(articles)

	bowen.BuildSucResponse(articles)
}

// @router /admin/bowen/modifystatus [post]
func (bowen *BowenController) ModifyStatus() {
	defer bowen.PanicHandler()

	articles := new(model.Articles)
	bowen.ParseJson(articles)

	service.BowenBiz.ModifyStatus(articles)

	bowen.BuildSucResponse("success")
}

// @router /admin/bowen/delete [post]
func (bowen *BowenController) Delete() {
	defer bowen.PanicHandler()

	articles := new(model.Articles)
	bowen.ParseJson(articles)

	service.BowenBiz.DeleteArticles(articles)

	bowen.BuildSucResponse("success")
}

// @router /admin/bowen/tag/delete [post]
func (bowen *BowenController) DeleteArticlesTag() {
	defer bowen.PanicHandler()

	articlesTag := new(model.ArticlesTag)
	bowen.ParseJson(articlesTag)

	service.BowenBiz.DeleteArticlesTag(articlesTag)

	bowen.BuildSucResponse("success")
}

// @router /admin/bowen/category/delete [post]
func (bowen *BowenController) DeleteArticlesCategory() {
	defer bowen.PanicHandler()

	articlesCategory := new(model.ArticlesCategory)
	bowen.ParseJson(articlesCategory)

	service.BowenBiz.DeleteArticlesCategory(articlesCategory)

	bowen.BuildSucResponse("success")
}

// @router /admin/bowen/collect/status [get]
func (bowen *BowenController) CollectStatus() {
	defer bowen.PanicHandler()

	bowen.BuildSucResponse(service.BowenBiz.CollectStatus())
}

// @router /admin/bowen/collect/type [get]
func (bowen *BowenController) CollectType() {
	defer bowen.PanicHandler()

	bowen.BuildSucResponse(service.BowenBiz.CollectType())
}
