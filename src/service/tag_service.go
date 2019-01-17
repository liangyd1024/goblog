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

type tagService struct {
	baseService
}

var TagBiz tagService

func init() {
	TagBiz = tagService{}
}

func (tagSer tagService) GetMapper() *dal.TagMapper {
	return new(dal.TagMapper)
}

func (tagSer tagService) GetAllTag(tag *model.Tag) []*model.Tag {
	tagMapper := tagSer.GetMapper()
	return tagMapper.GetByCondition(tag)
}

func (tagSer tagService) CreateTag(tag *model.Tag) {
	check.CheckParams(tag)
	tagMapper := tagSer.GetMapper()
	//通过名字过滤
	//if tagMapper.Get(tag) != nil {
	//	bizerror.BizError400002.PanicError()
	//}
	tagMapper.Insert(tag)
}

func (tagSer tagService) DeleteTag(tag *model.Tag) {
	tagMapper := tagSer.GetMapper()
	if tagMapper.Get(tag) == nil {
		return
	}

	articlesTag := &model.ArticlesTag{TagId: tag.Id}
	articlesTag.InitPaging()
	articlesTags := BowenBiz.GetMapper().GetArticlesTags(articlesTag)
	if articlesTags != nil && len(articlesTags) > 0 {
		bizerror.BizError400003.PanicError()
	}

	tagMapper.Delete(tag)
}

func (tagSer tagService) QueryTagBowen(tag *model.Tag) ([]*model.Articles, model.Paging) {
	paging := tag.Paging
	tagMapper := tagSer.GetMapper()
	if tagMapper.Get(tag) == nil {
		return nil, model.Paging{}
	}

	bowenMapper := BowenBiz.GetMapper()

	articlesTag := &model.ArticlesTag{TagId: tag.Id, Paging: paging}
	articlesTags := bowenMapper.GetArticlesTags(articlesTag)

	articlesList := make([]*model.Articles, len(articlesTags))
	for index, articlesTag := range articlesTags {
		articles := &model.Articles{Id: articlesTag.ArticlesId}
		articlesList[index] = bowenMapper.Get(articles).(*model.Articles)
		articles.Tags, articles.ArticlesTags = bowenMapper.GetTags(articles)
		articles.Categorys, articles.ArticlesCategorys = bowenMapper.GetCategorys(articles)
	}

	return articlesList, articlesTag.Paging
}
