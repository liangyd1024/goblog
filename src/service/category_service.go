//AppName: goblog
//Version: V1.0.0
//User: marco
//Date: 2018/11/21

package service

import (
	"goblog/src/dal"
	"goblog/src/model"
	"goblog/src/utils/bizerror"
	"goblog/src/utils/check"
)

type categoryService struct {
	baseService
}

var CategoryBiz categoryService

func init() {
	CategoryBiz = categoryService{}
}

func (categorySer categoryService) GetMapper() *dal.CategoryMapper {
	return new(dal.CategoryMapper)
}

func (categorySer categoryService) GetAllCategory(category *model.Category) []*model.Category {
	categoryMapper := categorySer.GetMapper()
	return categoryMapper.GetByCondition(category)
}

func (categorySer categoryService) CreateCategory(category *model.Category) {
	check.CheckParams(category)
	categoryMapper := categorySer.GetMapper()
	//if categoryMapper.Get(category) != nil {
	//	bizerror.BizError400002.PanicError()
	//}
	categoryMapper.Insert(category)
}

func (categorySer categoryService) DeleteCategory(category *model.Category) {
	categoryMapper := categorySer.GetMapper()
	if categoryMapper.Get(category) == nil {
		return
	}
	articlesTags := BowenBiz.GetMapper().GetArticlesCategorys(&model.ArticlesCategory{CategoryId: category.Id})
	if articlesTags != nil && len(articlesTags) > 0 {
		bizerror.BizError400003.PanicError()
	}

	categoryMapper.Delete(category)
}

func (categorySer categoryService) QueryCategoryBowen(category *model.Category) ([]*model.Articles, model.Paging) {
	paging := category.Paging
	categoryMapper := categorySer.GetMapper()
	if categoryMapper.Get(category) == nil {
		return nil, model.Paging{}
	}

	bowenMapper := BowenBiz.GetMapper()

	articlesCategory := &model.ArticlesCategory{CategoryId: category.Id, Paging: paging}
	articlesCategorys := bowenMapper.GetArticlesCategorys(articlesCategory)

	articlesList := make([]*model.Articles, len(articlesCategorys))
	for index, articlesCategory := range articlesCategorys {
		articles := &model.Articles{Id: articlesCategory.ArticlesId}
		articlesList[index] = bowenMapper.Get(articles).(*model.Articles)
		articles.Tags, articles.ArticlesTags = bowenMapper.GetTags(articles)
		articles.Categorys, articles.ArticlesCategorys = bowenMapper.GetCategorys(articles)
	}

	return articlesList, articlesCategory.Paging
}
