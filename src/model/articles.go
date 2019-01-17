package model

import (
	"time"
)

//文章
type Articles struct {
	Id          int       `json:"id" orm:"pk;auto"`                              //博文编号
	Title       string    `json:"title" valid:"Required" orm:"size(64)"`         //标题
	Desc        string    `json:"desc" valid:"Required" orm:"type(text)"`        //文章描述
	ImgUrl      string    `json:"imgUrl" orm:"size(64);null"`                    //博文封面url
	Publisher   string    `json:"publisher" orm:"size(64)"`                      //发布人
	PublishTime time.Time `json:"publishTime" orm:"auto_now_add;type(datetime)"` //发布时间
	BrowseNum   int       `json:"browseNum"`                                     //浏览数
	CommentNum  int       `json:"commentNum"`                                    //评论数
	PraiseNum   int       `json:"praiseNum"`                                     //点赞数
	Status      string    `json:"status" orm:"size(16);default(INIT)"`
	Type        string    `json:"type" orm:"size(16);default(ORIGINAL)"` //文章类型
	Base
	Paging `json:"paging"` //分页信息

	ArticlesDetails *ArticlesDetails `json:"articlesDetails" orm:"-"` //博文详情
	Tags            []*Tag           `json:"tags" orm:"-"`            //博文标签集
	Categorys       []*Category      `json:"categorys" orm:"-"`       //博文分类集

	ArticlesTags      []*ArticlesTag      `json:"articlesTags" orm:"-"`
	ArticlesCategorys []*ArticlesCategory `json:"articlesCategorys" orm:"-"`
}

//表名
func (art *Articles) TableName() string {
	return "articles"
}
