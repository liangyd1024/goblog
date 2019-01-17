package dal

import (
	"fmt"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
	"goblog/src/config"
	. "goblog/src/logs"
	"goblog/src/model"
)

func init() {
	dataSource := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8&loc=Local", config.DB.DbUser, config.DB.DbPwd, config.DB.DbUrls, config.DB.DbName)
	Log.Printf("call mysql dataSource:%v", dataSource)
	orm.RegisterDriver("mysql", orm.DRMySQL)
	orm.RegisterDataBase("default", "mysql", dataSource, config.DB.DbMaxIdleConns, config.DB.DbMaxConns)
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

	orm.RunSyncdb("default", config.DB.DbForce, true)
}
