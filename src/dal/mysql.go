package dal

import (
	"fmt"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
	"goblog/src/config"
	. "goblog/src/logs"
	"goblog/src/model"
	"goblog/src/utils/bizerror"
	"time"
)

const (
	AliasName = "default"
)

func init() {
	dataSource := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8&loc=Local", config.DB.DbUser, config.DB.DbPwd, config.DB.DbUrls, config.DB.DbName)
	Log.Sys("call mysql dataSource:%v", dataSource)
	err := orm.RegisterDriver("mysql", orm.DRMySQL)
	bizerror.Check(err)

	err = orm.RegisterDataBase(AliasName, "mysql", dataSource, config.DB.DbMaxIdleConns, config.DB.DbMaxConns)
	bizerror.Check(err)

	orm.RegisterModelWithPrefix("t_goblog_",
		new(model.Articles),
		new(model.ArticlesDetails),
		new(model.Tag),
		new(model.ArticlesTag),
		new(model.Category),
		new(model.ArticlesCategory),
		new(model.Comment),
		new(model.User),
	)

	orm.SetMaxIdleConns(AliasName, config.DB.DbMaxIdleConns)
	orm.SetMaxOpenConns(AliasName, config.DB.DbMaxConns)

	db, err := orm.GetDB(AliasName)
	bizerror.Check(err)
	db.SetConnMaxLifetime(time.Duration(time.Second * time.Duration(config.DB.ConnMaxLifetime)))

	err = orm.RunSyncdb(AliasName, config.DB.DbForce, true)
	bizerror.Check(err)
}
