package model

import "math"

//博文标签关联
type ArticlesTag struct {
	Id         int `json:"id" orm:"pk;auto"`       //主键
	ArticlesId int `json:"articlesId" orm:"index"` //博文编号
	TagId      int `json:"tagId" orm:"index"`      //标签编号
	Base
	Paging
}

//多字段唯一索引
func (art *ArticlesTag) TableUnique() [][]string {
	return [][]string{
		{"ArticlesId", "TagId"},
	}
}

func (art *ArticlesTag) InitPaging() {
	if art.PageSize == 0 {
		art.PageSize = math.MaxInt32
	}
}
