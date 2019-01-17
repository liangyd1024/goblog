//AppName: goblog
//Version: V1.0.0
//User: marco
//Date: 2018/11/29

package model

type ArticlesCollect struct {
	Total  int    `json:"total"`  //总数
	Status string `json:"status"` //博文状态
	Type   string `json:"type"`   //博文类型
}
