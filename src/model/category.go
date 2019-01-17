package model

//分类
type Category struct {
	Id           int    `json:"id" orm:"pk;auto"`                   //栏目编号
	CategoryName string `json:"categoryName" valid:"Required;MaxSize(64)" orm:"size(64);unique"` //栏目名称
	ArticlesNum  int    `json:"articlesNum"`                        //文章数
	Base
	ArticlesCategorys []*ArticlesCategory `json:"articlesCategorys" orm:"-"` //分类博文集
	Paging
}

