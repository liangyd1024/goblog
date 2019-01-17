//AppName: goblog
//Version: V1.0.0
//User: marco
//Date: 2018/12/15

package model

type Search struct {
	Id      int    `json:"id"`
	Stype   string `json:"type"`    //类型：tag/category
	Content string `json:"content"` //搜索内容
	Paging  `json:"paging"`
}


