package model

//文章栏目
type ArticlesCategory struct {
	Id         int `json:"id" orm:"pk;auto"`       //主键
	ArticlesId int `json:"articlesId" orm:"index"` //博文编号
	CategoryId int `json:"categoryId" orm:"index"` //栏目编号
	Base
	Paging
}

//多字段唯一索引
func (art *ArticlesCategory) TableUnique() [][]string {
	return [][]string{
		{"ArticlesId", "CategoryId"},
	}
}
