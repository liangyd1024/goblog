//AppName: goblog
//Version: V1.0.0
//User: marco
//Date: 2018/12/29

package service

import (
	"github.com/go-ego/riot"
	"github.com/go-ego/riot/types"
	"goblog/src/component"
	. "goblog/src/logs"
	"goblog/src/model"
	"goblog/src/utils/bizerror"
	"goblog/src/utils/constant"
	"goblog/src/utils/dataconv"
	"goblog/src/utils/datetime"
	"math"
	"reflect"
	"strconv"
	"time"
)

const (
	TAG                = "tag"
	CATEGORY           = "category"
	ARTICLES_FULL_TEXT = "articles"
	PLACE_OF_FILE      = "placeoffile"
)

//全文搜索引擎
var fullTextSearcher = riot.Engine{}

//搜索服务
var SearchBiz searchService

//引擎类型
var engineType map[string]SearchEngine

func init() {
	SearchBiz = searchService{}
	engineType = make(map[string]SearchEngine)
	//标签搜索
	engineType[TAG] = TagSearchEngine{}
	//栏目搜索
	engineType[CATEGORY] = CategorySearchEngine{}
	//博文搜索
	engineType[ARTICLES_FULL_TEXT] = ArticlesSearchEngine{}
	//归档搜索
	engineType[PLACE_OF_FILE] = PlaceOfFileSearchEngine{}

	//component.GoRoutine(func() {
	//	gob.Register(ArticlesScoringFields{})
	//	fullTextSearcher.Init(types.EngineOpts{
	//		Using:       3,
	//		GseDict:     "zh",
	//		UseStore:    true,
	//		StoreFolder: "./indexer",
	//		StoreShards: 8,
	//		StoreEngine: "",
	//	})
	//	fullTextSearcher.Flush()
	//})

}

type searchContent struct {
	Id      int
	Title   string
	Desc    string
	Content string
}

type searchService struct {
	baseService
}

//初始化全文搜索索引
func (searchSer searchService) RefreshFullTextSearcher() {

	articles := new(model.Articles)
	articles.Paging.PageSize = math.MaxInt32
	articles.Status = constant.BOWEN_STATUS_PUBLISH
	articlesList := BowenBiz.GetBowenCondition(articles)

	searchSer.Index(articlesList)
}

func (searchSer searchService) GetSearchEngine(search *model.Search) SearchEngine {
	searchEngine := engineType[search.Stype]
	if searchEngine == nil {
		bizerror.BizError404001.PanicError()
	}
	return engineType[search.Stype]
}

func (searchSer searchService) IndexSingle(articles *model.Articles) {
	//searchSer.Index([]*model.Articles{articles})
}

//索引数据
func (searchSer searchService) Index(articlesList []*model.Articles) {
	component.GoRoutine(func() {
		Log.Info("call initFullTextSearcher start size:%v", len(articlesList))
		for _, article := range articlesList {
			tagLen := len(article.Tags)
			BowenBiz.GetBowen(article)
			labels := make([]string, tagLen+len(article.Categorys))
			for index, tag := range article.Tags {
				labels[index] = tag.TagName
			}
			for index, category := range article.Categorys {
				labels[tagLen+index] = category.CategoryName
			}
			content := searchContent{
				Id:      article.Id,
				Title:   article.Title,
				Desc:    article.Desc,
				Content: article.ArticlesDetails.Content,
			}
			fullTextSearcher.Index(
				strconv.Itoa(article.Id),
				types.DocData{
					Content: dataconv.JsonM2Str(content),
					Labels:  labels,
					Fields: ArticlesScoringFields{
						BrowseNum:   article.BrowseNum,
						PublishTime: article.PublishTime,
						CommentNum:  article.CommentNum,
						PraiseNum:   article.PraiseNum,
					},
				})
		}
		fullTextSearcher.Flush()
		allDocIds := fullTextSearcher.GetAllDocIds()
		Log.Info("call initFullTextSearcher end length:%v,allDocIds:%v", len(allDocIds), allDocIds)
	})
}

//删除索引数据
func (searchSer searchService) RemoveIndex(articles *model.Articles) {
	//component.GoRoutine(func() {
	//	fullTextSearcher.RemoveDoc(strconv.Itoa(articles.Id), true)
	//	fullTextSearcher.Flush()
	//})
}

type ArticlesScoringFields struct {
	BrowseNum   int
	PublishTime time.Time
	CommentNum  int
	PraiseNum   int
}

type ArticlesScoringCriteria struct {
}

//博文被搜索出来后评分排序规则
func (score ArticlesScoringCriteria) Score(doc types.IndexedDoc, fields interface{}) []float32 {
	if reflect.TypeOf(fields) != reflect.TypeOf(ArticlesScoringFields{}) {
		return []float32{}
	}
	result := make([]float32, 5)
	scoring := fields.(ArticlesScoringFields)
	result[0] = float32(scoring.BrowseNum)
	result[1] = float32(int(doc.BM25))
	result[2] = float32(scoring.PublishTime.UnixNano())
	result[3] = float32(scoring.PraiseNum)
	result[4] = float32(scoring.CommentNum)
	return result
}

type SearchEngine interface {
	Search(search *model.Search) ([]*model.Articles, model.Paging)
}

type TagSearchEngine struct {
}

func (TagSearchEngine) Search(search *model.Search) ([]*model.Articles, model.Paging) {
	return TagBiz.QueryTagBowen(&model.Tag{Id: search.Id, Paging: search.Paging})
}

type CategorySearchEngine struct {
}

func (CategorySearchEngine) Search(search *model.Search) ([]*model.Articles, model.Paging) {
	return CategoryBiz.QueryCategoryBowen(&model.Category{Id: search.Id, Paging: search.Paging})
}

type ArticlesSearchEngine struct {
}

//全文检索
func (ArticlesSearchEngine) Search(search *model.Search) ([]*model.Articles, model.Paging) {
	Log.Info("call ArticlesSearchEngine start search:%v", search.Content)
	articles := &model.Articles{Title: search.Content, Desc: search.Content, Paging: search.Paging}
	articlesList := BowenBiz.GetBowenCondition(articles)
	paging := articles.Paging
	return articlesList, paging

	//paging := search.Paging
	//pageSize, offset := paging.StartPage()
	//searchDoc := fullTextSearcher.SearchDoc(types.SearchReq{
	//	Text: search.Content,
	//	//Labels: []string{search.Content},
	//	RankOpts: &types.RankOpts{
	//		OutputOffset:    offset,
	//		MaxOutputs:      pageSize,
	//		ScoringCriteria: ArticlesScoringCriteria{},
	//	},
	//})
	//Log.Info("call ArticlesSearchEngine Docs:%v", searchDoc.Docs)
	//docLen := len(searchDoc.Docs)
	//articlesList := make([]*model.Articles, docLen)
	//for index, searchDoc := range searchDoc.Docs {
	//	Log.Info("call ArticlesSearchEngine DocId:%v", searchDoc.DocId)
	//	articles := new(model.Articles)
	//	articles.Id, _ = strconv.Atoi(searchDoc.DocId)
	//	articles.Status = constant.BOWEN_STATUS_PUBLISH
	//	articlesList[index] = BowenBiz.GetBowenCondition(articles)[0]
	//}
	//paging.CalPages(int64(searchDoc.NumDocs))
	//
	//Log.Info("call ArticlesSearchEngine end docLength:%v", docLen)
	//return articlesList, paging
}

type PlaceOfFileSearchEngine struct {
}

func (PlaceOfFileSearchEngine) Search(search *model.Search) ([]*model.Articles, model.Paging) {
	articles := &model.Articles{PublishTime: datetime.ParseTime(datetime.FM_DATE_MOUNTH, search.Content), Paging: search.Paging}
	articlesList := BowenBiz.GetBowenCondition(articles)
	paging := articles.Paging
	return articlesList, paging
}
