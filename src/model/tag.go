package model

type Tag struct {
	Id          int    `json:"id" orm:"pk;auto"`                                           //标签编号
	TagName     string `json:"tagName" valid:"Required;MaxSize(64)" orm:"size(32);unique"` //标签名称
	ArticlesNum int    `json:"articlesNum" orm:""`                                         //文章数
	Base
	Paging
	ArticlesTags []*ArticlesTag `json:"articlesTags" orm:"-"` //标签博文集合
}
