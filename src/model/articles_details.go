package model

import (
	"fmt"
	"strconv"
)

//博文详情
type ArticlesDetails struct {
	Id         int    `json:"id" orm:"pk;auto"`                          //博文编号
	EditorType string `json:"editorType" valid:"Required"`               //博文内容编辑类型
	Content    string `json:"content" valid:"Required" orm:"type(text)"` //博文内容
	Base
}

//表名
func (articlesDetails ArticlesDetails) TableName() string {
	return "articles_details"
}

func (articlesDetails ArticlesDetails) String() string {
	content := ""
	if articlesDetails.Content != "" {
		content = string([]byte(articlesDetails.Content)[:30])
	}
	return fmt.Sprintf("user:{Id:%s,EditorType:%s,Content:%s}",
		strconv.Itoa(articlesDetails.Id),
		articlesDetails.EditorType,
		content,
	)
}
