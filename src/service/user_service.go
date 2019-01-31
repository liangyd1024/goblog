//AppName: goblog
//Version: V1.0.0
//User: marco
//Date: 2018/12/12

package service

import (
	"goblog/src/dal"
	"goblog/src/model"
	"goblog/src/utils/bizerror"
	"goblog/src/utils/check"
	"goblog/src/utils/crypt"
	"log"
)

var UserBiz userBiz

type userBiz struct {
}

func init() {
	UserBiz = userBiz{}
}

func (userSer userBiz) GetMapper() *dal.UserMapper {
	return new(dal.UserMapper)
}

func (userSer userBiz) CreateUser(user *model.User) *model.User {
	check.CheckParams(user)
	newUser := userSer.GetUser(user.UserName)
	if newUser == nil {
		user.UserPwd = crypt.GetMd5(user.UserPwd)
		log.Printf("call CreateUser user:%v", user)
		userSer.GetMapper().Insert(user)
	}
	return user
}

func (userSer userBiz) GetUser(userName string) *model.User {
	user := &model.User{UserName: userName}
	userList := userSer.ListUser(user)
	if userList != nil && len(userList) > 0 {
		if len(userList) > 1 {
			bizerror.BizError400004.PanicError()
		}
		return userList[0]
	}
	return nil
}

func (userSer userBiz) ListUser(user *model.User) []*model.User {
	return userSer.GetMapper().GetByCondition(user)
}

func (userSer userBiz) CheckPwd(user *model.User) *model.User {
	findUser := userSer.GetUser(user.UserName)
	md5Pwd := crypt.GetMd5(user.UserPwd)
	if findUser == nil || findUser.UserPwd != md5Pwd {
		bizerror.BizError400101.PanicError()
	}
	return findUser
}
