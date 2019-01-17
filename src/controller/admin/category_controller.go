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

type CategoryController struct {
	controller.BaseController
}

func (categoryController *CategoryController) URLMapping() {
	categoryController.Mapping("List", categoryController.List)
	categoryController.Mapping("Add", categoryController.Add)
	categoryController.Mapping("DeleteArticles", categoryController.Delete)

	categoryController.Mapping("BowenList", categoryController.BowenList)
}

// @router /admin/category/list [get]
func (categoryController *CategoryController) List() {
	defer categoryController.PanicHandler()

	categoryController.BuildSucResponse(service.CategoryBiz.GetAllCategory(new(model.Category)))
}

// @router /admin/category/add [post]
func (categoryController *CategoryController) Add() {
	defer categoryController.PanicHandler()

	category := new(model.Category)
	categoryController.ParseJson(category)

	category.CreateBy = categoryController.GetUserName()
	service.CategoryBiz.CreateCategory(category)

	categoryController.BuildSucResponse(category)
}

// @router /admin/category/delete [post]
func (categoryController *CategoryController) Delete() {
	defer categoryController.PanicHandler()

	category := new(model.Category)
	categoryController.ParseJson(category)

	category.UpdateBy = categoryController.GetUserName()
	service.CategoryBiz.DeleteCategory(category)

	categoryController.BuildSucResponse("success")
}

// @router /admin/category/bowen/list [post]
func (categoryController *CategoryController) BowenList() {
	defer categoryController.PanicHandler()

	category := new(model.Category)
	categoryController.ParseJson(category)

	categoryController.BuildSucPagingResponse(service.CategoryBiz.QueryCategoryBowen(category))
}
