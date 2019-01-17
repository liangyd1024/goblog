//AppName: goblog
//Version: V1.0.0
//User: marco
//Date: 2018/11/21

package admin

import (
	"goblog/src/controller"
	"goblog/src/model"
	"goblog/src/service"
)

type TagController struct {
	controller.BaseController
}

func (tagController *TagController) URLMapping() {
	tagController.Mapping("List", tagController.List)
	tagController.Mapping("Add", tagController.Add)
	tagController.Mapping("DeleteArticles", tagController.Delete)

	tagController.Mapping("BowenList", tagController.BowenList)
}

// @router /admin/tag/list [get]
func (tagController *TagController) List() {
	defer tagController.PanicHandler()

	tagController.BuildSucResponse(service.TagBiz.GetAllTag(new(model.Tag)))
}

// @router /admin/tag/add [post]
func (tagController *TagController) Add() {
	defer tagController.PanicHandler()

	tag := new(model.Tag)
	tagController.ParseJson(tag)

	tag.CreateBy = tagController.GetUserName()
	service.TagBiz.CreateTag(tag)

	tagController.BuildSucResponse(tag)
}

// @router /admin/tag/delete [post]
func (tagController *TagController) Delete() {
	defer tagController.PanicHandler()

	tag := new(model.Tag)
	tagController.ParseJson(tag)

	tag.UpdateBy = tagController.GetUserName()
	service.TagBiz.DeleteTag(tag)

	tagController.BuildSucResponse("success")
}

// @router /admin/tag/bowen/list [post]
func (tagController *TagController) BowenList(){
	defer tagController.PanicHandler()

	tag := new(model.Tag)
	tagController.ParseJson(tag)

	tagController.BuildSucPagingResponse(service.TagBiz.QueryTagBowen(tag))
}
