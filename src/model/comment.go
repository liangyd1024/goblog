package model

import "time"

type Comment struct {
	Id          int       `json:"id" orm:"pk;auto"`                                         //评论编号
	ArticlesId  int       `json:"articlesId"`                                               //博文编号
	Content     string    `json:"content" valid:"Required;MaxSize(2048)" orm:"type(text)"`  //评论内容
	Commentator string    `json:"commentator" valid:"Required;MaxSize(64)" orm:"size(32)"`  //评论人
	Email       string    `json:"email" valid:"Required;MaxSize(64);Email" orm:"size(128)"` //邮箱
	CommentTime time.Time `json:"commentTime" orm:"auto_now;type(datetime)"`                //评论时间
	PraiseNum   int       `json:"praiseNum"`                                                //点赞数
	ParentId    int       `json:"parentId" orm:"default(0)"`                                //父级评论编号
	Base
	Paging           `json:"paging"`                              //分页信息
	ReplyCommentList []*Comment `json:"replyCommentList" orm:"-"` //回复列表
}
