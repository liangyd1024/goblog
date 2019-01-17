package model

import (
	"fmt"
	"strconv"
	"time"
)

type User struct {
	Id        int       `json:"id" orm:"pk;auto"`                         //主键
	UserName  string    `json:"userName" valid:"Required" orm:"size(32)"` //用户名
	UserPwd   string    `json:"userPwd" valid:"Required" orm:"size(128)"` //用户密码
	UserType  string    `json:"userType" orm:"size(16)"`                  //用户类型
	Status    string    `json:"status" orm:"size(16)"`                    //用户状态
	NickName  string    `json:"nickName" orm:"size(64);null"`             //昵称
	HeadUrl   string    `json:"headUrl" orm:"size(64);null"`              //头像url
	Phone     string    `json:"phone" orm:"size(32);null"`                //手机号
	Email     string    `json:"email" orm:"size(32);null"`                //邮箱
	LoginTime time.Time `json:"loginTime" orm:"type(datetime);null"`      //登录时间
	LoginIp   string    `json:"loginIp" orm:"size(32);null"`              //登陆IP
	Paging    `json:"paging"`                                             //分页信息
	Base

	Captcha string `json:"captcha" orm:"-"` //验证码
}

func (user User) String() string {
	return fmt.Sprintf("user:{Id:%s,UserName:%s,UserType:%s,Status:%s,NickName:%s,HeadUrl:%s,Phone:%s,Email:%s,LoginTime:%s,LoginIp:%s}",
		strconv.Itoa(user.Id),
		user.UserName,
		user.UserType,
		user.Status,
		user.NickName,
		user.HeadUrl,
		user.Phone,
		user.Email,
		user.LoginTime,
		user.LoginIp,
	)
}
