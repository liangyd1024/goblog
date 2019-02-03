//AppName: goblog
//Version: V1.0.0
//User: marco
//Date: 2018/12/12

package dal

import (
	"github.com/astaxie/beego/orm"
	. "goblog/src/logs"
	"goblog/src/model"
	"goblog/src/utils/bizerror"
)

type UserMapper struct {
	BaseMapper
}

func (userMapper *UserMapper) GetByCondition(user *model.User) []*model.User {
	var userList []*model.User
	pageSize, offset := user.Paging.StartPage()
	querySeter := getOrmer().QueryTable(user)

	cond := orm.NewCondition()
	if user.Id != 0 {
		cond = cond.Or("id", user.Id)
		querySeter = querySeter.SetCond(cond)
	}
	if user.UserName != "" {
		cond = cond.Or("user_name", user.UserName)
		querySeter = querySeter.SetCond(cond)
	}
	if user.UserPwd != "" {
		cond = cond.Or("user_pwd", user.UserPwd)
		querySeter = querySeter.SetCond(cond)
	}
	if user.UserType != "" {
		cond = cond.Or("user_type", user.UserType)
		querySeter = querySeter.SetCond(cond)
	}
	if user.Phone != "" {
		cond = cond.Or("phone", user.Phone)
		querySeter = querySeter.SetCond(cond)
	}

	total, err := querySeter.Count()
	bizerror.Check(err)

	rows, err := querySeter.
		OrderBy("-update_at").
		Limit(pageSize, offset).
		All(&userList)
	bizerror.Check(err)

	user.Paging.CalPages(total)

	Log.Info("call GetByCondition rows:%v", rows)
	return userList
}
