//AppName: goblog
//Version: V1.0.0
//User: marco
//Date: 2018/11/20

package service

import (
	"goblog/src/dal"
	. "goblog/src/logs"
	"goblog/src/model"
	"goblog/src/utils/bizerror"
	"goblog/src/utils/check"
	"goblog/src/utils/constant"
	"sync"
	"time"
)

type bowenService struct {
	baseService
}

var BowenBiz bowenService

func init() {
	BowenBiz = bowenService{baseService{Mutex: new(sync.Mutex)}}
}

func (bowenSer bowenService) GetMapper() *dal.BowenMapper {
	return new(dal.BowenMapper)
}

func (bowenSer bowenService) Publish(articles *model.Articles) {
	check.CheckParams(articles)
	check.CheckParams(articles.ArticlesDetails)

	//博文标签集
	if articles.Tags != nil && len(articles.Tags) > 0 {
		articlesTags := make([]*model.ArticlesTag, len(articles.Tags))
		for index, tag := range articles.Tags {
			articlesTags[index] = &model.ArticlesTag{ArticlesId: articles.Id, TagId: tag.Id}
		}
		articles.ArticlesTags = articlesTags
	}

	//博文栏目集
	if articles.Categorys != nil && len(articles.Categorys) > 0 {
		articlesCategorys := make([]*model.ArticlesCategory, len(articles.Categorys))
		for index, category := range articles.Categorys {
			articlesCategorys[index] = &model.ArticlesCategory{Id: articles.Id, CategoryId: category.Id}
		}
		articles.ArticlesCategorys = articlesCategorys
	}

	bowenSer.GetMapper().Publish(articles)

	//索引博文
	SearchBiz.IndexSingle(articles)
}

func (bowenSer bowenService) ModifyStatus(articles *model.Articles) {
	queryArticles := &model.Articles{Id: articles.Id}
	bowenMapper := bowenSer.GetMapper()
	if bowenMapper.Get(queryArticles) == nil {
		bizerror.BizError404001.PanicError()
	}
	if queryArticles.Status != constant.BOWEN_STATUS_INIT {
		bizerror.BizError400004.PanicError()
	}
	queryArticles.Status = articles.Status
	queryArticles.PublishTime = time.Now()
	bowenMapper.Update(queryArticles, "status", "publish_time")
}

func (bowenSer bowenService) ModifyArticles(articles *model.Articles) {
	check.CheckParams(articles)
	check.CheckParams(articles.ArticlesDetails)

	Log.Printf("call ModifyArticles articles:%+v", articles)

	queryArticles := &model.Articles{Id: articles.Id}
	bowenMapper := bowenSer.GetMapper()
	if bowenMapper.Get(queryArticles) == nil {
		bizerror.BizError404001.PanicError()
	}
	queryArticlesDetails := &model.ArticlesDetails{Id: articles.Id}
	if bowenMapper.Get(queryArticlesDetails) == nil {
		bizerror.BizError404001.PanicError()
	}
	queryArticles.ArticlesDetails = queryArticlesDetails

	queryArticles.Title = articles.Title
	queryArticles.Desc = articles.Desc
	queryArticles.Type = articles.Type
	queryArticlesDetails.Content = articles.ArticlesDetails.Content

	//博文标签集
	if articles.Tags != nil && len(articles.Tags) > 0 {
		articlesTags := make([]*model.ArticlesTag, len(articles.Tags))
		for index, tag := range articles.Tags {
			articlesTags[index] = &model.ArticlesTag{ArticlesId: articles.Id, TagId: tag.Id}
		}
		queryArticles.ArticlesTags = articlesTags
	}

	//博文栏目集
	if articles.Categorys != nil && len(articles.Categorys) > 0 {
		articlesCategorys := make([]*model.ArticlesCategory, len(articles.Categorys))
		for index, category := range articles.Categorys {
			articlesCategorys[index] = &model.ArticlesCategory{ArticlesId: articles.Id, CategoryId: category.Id}
		}
		queryArticles.ArticlesCategorys = articlesCategorys
	}

	bowenMapper.Modify(queryArticles)

	//更新博文索引
	//SearchBiz.IndexSingle(queryArticles)
}

func (bowenSer bowenService) DeleteArticles(articles *model.Articles) {
	Log.Printf("call DeleteArticles id:%+v", articles.Id)

	bowenMapper := bowenSer.GetMapper()
	if bowenMapper.Get(&model.Articles{Id: articles.Id}) == nil {
		bizerror.BizError404001.PanicError()
	}
	bowenSer.GetMapper().DeleteArticles(articles)

	//删除博文索引
	SearchBiz.RemoveIndex(articles)
}

func (bowenSer bowenService) DeleteArticlesTag(articlesTag *model.ArticlesTag) {
	Log.Printf("call DeleteArticlesTag id:%+v", articlesTag.Id)

	bowenMapper := bowenSer.GetMapper()

	articlesTag.InitPaging()
	articleTags := bowenMapper.GetArticlesTags(articlesTag)
	if len(articleTags) > 0 {
		//删除博文对应标签
		bowenMapper.DeleteArticlesTag(articleTags[0])
	}
}

func (bowenSer bowenService) DeleteArticlesCategory(articlesCategory *model.ArticlesCategory) {
	Log.Printf("call DeleteArticlesCategory id:%+v", articlesCategory.Id)

	bowenMapper := bowenSer.GetMapper()

	articleCategorys := bowenMapper.GetArticlesCategorys(articlesCategory)
	if len(articleCategorys) > 0 {
		//删除博文对应栏目
		bowenMapper.DeleteArticlesCategory(articleCategorys[0])
	}
}

func (bowenSer bowenService) GetBowenCondition(articles *model.Articles) []*model.Articles {
	Log.Printf("call GetBowenCondition articles:%+v", articles)
	bowenMapper := bowenSer.GetMapper()
	articlesList := bowenMapper.GetByCondition(articles)
	for _, articles := range articlesList {
		articles.Tags, articles.ArticlesTags = bowenMapper.GetTags(articles)
		articles.Categorys, articles.ArticlesCategorys = bowenMapper.GetCategorys(articles)
	}
	return articlesList
}

func (bowenSer bowenService) ListRecommendBowen(size int) []*model.Articles {
	bowenMapper := bowenSer.GetMapper()
	return bowenMapper.ListRecommendArticles(size)
}

func (bowenSer bowenService) GetBowen(articles *model.Articles) {
	bowenMapper := bowenSer.GetMapper()
	bowenMapper.Get(articles)
	articlesDetails := &model.ArticlesDetails{Id: articles.Id}
	bowenMapper.Get(articlesDetails)
	articles.ArticlesDetails = articlesDetails
	articles.Tags, articles.ArticlesTags = bowenMapper.GetTags(articles)
	articles.Categorys, articles.ArticlesCategorys = bowenMapper.GetCategorys(articles)

	Log.Printf("call GetBowen articles:%+v", articles)
}

func (bowenSer bowenService) CollectStatus() []*model.ArticlesCollect {
	bowenMapper := bowenSer.GetMapper()

	return bowenMapper.CollectGroup("status")
}

func (bowenSer bowenService) CollectType() []*model.ArticlesCollect {
	bowenMapper := bowenSer.GetMapper()

	return bowenMapper.CollectGroup("type")
}

func (bowenSer bowenService) CollectPlaceOfFile() []*model.ArticlesCollect {
	bowenMapper := bowenSer.GetMapper()

	return bowenMapper.CollectPlaceOfFile()
}

func (bowenSer bowenService) Browse(articles *model.Articles) {
	bowenSer.Lock()
	defer bowenSer.Unlock()

	Log.Printf("call Browse id:%v", articles.Id)
	bowenMapper := bowenSer.GetMapper()
	bowenMapper.Get(articles)

	articles.BrowseNum = articles.BrowseNum + 1
	bowenMapper.Update(articles, "browse_num", "update_at")
	Log.Printf("call Browse end id:%v", articles.Id)
}

func (bowenSer bowenService) Praise(articles *model.Articles) {
	bowenSer.Lock()
	defer bowenSer.Unlock()

	Log.Printf("call Praise id:%v", articles.Id)
	bowenMapper := bowenSer.GetMapper()
	bowenMapper.Get(articles)

	articles.PraiseNum = articles.PraiseNum + 1
	bowenMapper.Update(articles, "praise_num", "update_at")
	Log.Printf("call Praise end id:%v", articles.Id)
}

func (bowenSer bowenService) ListComment(comment *model.Comment) []*model.Comment {
	bowenMapper := bowenSer.GetMapper()
	commentList := bowenMapper.GetComments(comment)
	for _, comment := range commentList {
		replyComment := new(model.Comment)
		replyComment.ParentId = comment.Id
		replyComment.ArticlesId = comment.ArticlesId
		comment.ReplyCommentList = bowenMapper.GetComments(replyComment)
	}
	return commentList
}

func (bowenSer bowenService) PubComment(comment *model.Comment) {
	check.CheckParams(comment)

	bowenSer.Lock()
	defer bowenSer.Unlock()

	Log.Printf("call PubComment ArticlesId:%v", comment.ArticlesId)

	articles := &model.Articles{Id: comment.ArticlesId}
	bowenMapper := bowenSer.GetMapper()
	bowenMapper.Get(articles)

	bowenMapper.PubComment(comment, articles)

	Log.Printf("call PubComment end ArticlesId:%v", comment.ArticlesId)
}
