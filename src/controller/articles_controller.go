//AppName: goblog
//Version: V1.0.0
//User: marco
//Date: 2018/12/14

package controller

import (
	"encoding/base64"
	. "goblog/src/logs"
	"goblog/src/model"
	"goblog/src/service"
	"strconv"
)

const (
	COMMENT_COMMENTATOR = "commentCommentator"
	COMMENT_EMAIL       = "commentEmail"
)

type ArticlesController struct {
	BaseController
}

func (articlesController *ArticlesController) URLMapping() {
	articlesController.Mapping("ListArticles", articlesController.ListArticles)
	articlesController.Mapping("ListRecommendArticles", articlesController.ListRecommendArticles)
	articlesController.Mapping("ListTag", articlesController.ListTag)
	articlesController.Mapping("ListCategory", articlesController.ListCategory)
	articlesController.Mapping("ListPlaceOfFile", articlesController.ListPlaceOfFile)
	articlesController.Mapping("ToArticles", articlesController.ToArticles)
	articlesController.Mapping("GetArticles", articlesController.GetArticles)
	articlesController.Mapping("Browse", articlesController.Browse)
	articlesController.Mapping("Praise", articlesController.Praise)
	articlesController.Mapping("ListComment", articlesController.ListComment)
	articlesController.Mapping("PubComment", articlesController.PubComment)
	articlesController.Mapping("ReplyComment", articlesController.ReplyComment)
	articlesController.Mapping("ToSearch", articlesController.ToSearch)
	articlesController.Mapping("TypeSearch", articlesController.TypeSearch)
}

// @router /articles/list [post]
func (articlesController *ArticlesController) ListArticles() {
	defer articlesController.PanicHandler()

	articles := new(model.Articles)
	articlesController.ParseJson(articles)

	articlesController.BuildSucPagingResponse(service.BowenBiz.GetBowenCondition(articles), articles.Paging)
}

// @router /articles/recommend/list [get]
func (articlesController *ArticlesController) ListRecommendArticles() {
	defer articlesController.PanicHandler()

	articlesController.BuildSucResponse(service.BowenBiz.ListRecommendBowen(3))
}

// @router /articles/tag/list [get]
func (articlesController *ArticlesController) ListTag() {
	defer articlesController.PanicHandler()

	articlesController.BuildSucResponse(service.TagBiz.GetAllTag(new(model.Tag)))
}

// @router /articles/category/list [get]
func (articlesController *ArticlesController) ListCategory() {
	defer articlesController.PanicHandler()

	articlesController.BuildSucResponse(service.CategoryBiz.GetAllCategory(new(model.Category)))
}

// @router /articles/place-of-file/list [get]
func (articlesController *ArticlesController) ListPlaceOfFile() {
	defer articlesController.PanicHandler()

	articlesController.BuildSucResponse(service.BowenBiz.CollectPlaceOfFile())
}

// @router /articles/:id [get]
func (articlesController *ArticlesController) ToArticles() {
	Log.Printf("call ToDetails id:%v", articlesController.Ctx.Input.Param(":id"))
	articlesController.Data["id"] = articlesController.Ctx.Input.Param(":id")

	stdEncoding := base64.StdEncoding
	commentator, _ := stdEncoding.DecodeString(articlesController.Ctx.GetCookie(COMMENT_COMMENTATOR))
	email, _ := stdEncoding.DecodeString(articlesController.Ctx.GetCookie(COMMENT_EMAIL))
	articlesController.Data["commentCommentator"] = string(commentator)
	articlesController.Data["commentEmail"] = string(email)
	articlesController.TplName = "details.html"
}

// @router /articles/get [post]
func (articlesController *ArticlesController) GetArticles() {
	defer articlesController.PanicHandler()

	articles := new(model.Articles)
	articlesController.ParseJson(articles)

	service.BowenBiz.GetBowen(articles)

	articlesController.BuildSucResponse(articles)
}

// @router /articles/browse [post]
func (articlesController *ArticlesController) Browse() {
	defer articlesController.PanicHandler()

	articles := new(model.Articles)
	articlesController.ParseJson(articles)

	browseCookieKey := "browseCookie_" + strconv.Itoa(articles.Id)
	browseCookie := articlesController.Ctx.Input.Cookie(browseCookieKey)
	Log.Printf("call Browse browseCookieKey:%v,browseCookie:%v", browseCookieKey, browseCookie)
	if browseCookie == "" {
		service.BowenBiz.Browse(articles)
		articlesController.Ctx.SetCookie(browseCookieKey, "true", 3600)
		articlesController.BuildSucResponse(true)
	} else {
		articlesController.BuildSucResponse(false)
	}

}

// @router /articles/praise [post]
func (articlesController *ArticlesController) Praise() {
	defer articlesController.PanicHandler()

	articles := new(model.Articles)
	articlesController.ParseJson(articles)

	praiseCookieKey := "praiseCookie_" + strconv.Itoa(articles.Id)
	praiseCookie := articlesController.Ctx.GetCookie(praiseCookieKey)
	Log.Printf("call Praise praiseCookieKey:%v,praiseCookie:%v", praiseCookieKey, praiseCookie)
	if praiseCookie == "" {
		service.BowenBiz.Praise(articles)
		articlesController.Ctx.SetCookie(praiseCookieKey, "true", 3600)
		articlesController.BuildSucResponse(true)
	} else {
		articlesController.BuildSucResponse(false)
	}
}

// @router /articles/comment/list [post]
func (articlesController *ArticlesController) ListComment() {
	defer articlesController.PanicHandler()

	comment := new(model.Comment)
	articlesController.ParseJson(comment)

	articlesController.BuildSucPagingResponse(service.BowenBiz.ListComment(comment), comment.Paging)
}

// @router /articles/comment/pub [post]
func (articlesController *ArticlesController) PubComment() {
	defer articlesController.PanicHandler()

	comment := new(model.Comment)
	articlesController.ParseJson(comment)

	comment.CreateBy = comment.Commentator
	service.BowenBiz.PubComment(comment)

	stdEncoding := base64.StdEncoding
	articlesController.Ctx.SetCookie(COMMENT_EMAIL, stdEncoding.EncodeToString([]byte(comment.Email)))
	articlesController.Ctx.SetCookie(COMMENT_COMMENTATOR, stdEncoding.EncodeToString([]byte(comment.Commentator)))

	articlesController.BuildSucResponse(comment)
}

// @router /articles/comment/reply [post]
func (articlesController *ArticlesController) ReplyComment() {
	defer articlesController.PanicHandler()

	comment := new(model.Comment)
	articlesController.ParseJson(comment)

	comment.CreateBy = comment.Commentator
	service.BowenBiz.PubComment(comment)

	stdEncoding := base64.StdEncoding
	articlesController.Ctx.SetCookie(COMMENT_EMAIL, stdEncoding.EncodeToString([]byte(comment.Email)))
	articlesController.Ctx.SetCookie(COMMENT_COMMENTATOR, stdEncoding.EncodeToString([]byte(comment.Commentator)))

	articlesController.BuildSucResponse(comment)
}

// @router /articles/tosearch [get]
func (articlesController *ArticlesController) ToSearch() {
	defer articlesController.PanicHandler()

	articlesController.Data["id"] = articlesController.GetString("id")
	articlesController.Data["type"] = articlesController.GetString("type")
	articlesController.Data["content"] = articlesController.GetString("content")

	articlesController.TplName = "search.html"

}

// @router /articles/typesearch [post]
func (articlesController *ArticlesController) TypeSearch() {
	defer articlesController.PanicHandler()

	search := new(model.Search)
	articlesController.ParseJson(search)

	searchEngine := service.SearchBiz.GetSearchEngine(search)

	articlesController.BuildSucPagingResponse(searchEngine.Search(search))
}
