package config

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	. "goblog/src/logs"
	"goblog/src/utils/bizerror"
	"goblog/src/utils/constant"
)

type dbConfig struct {
	DbUser         string
	DbPwd          string
	DbUrls         string
	DbName         string
	DbMaxIdleConns int
	DbMaxConns     int
	DbForce        bool
	DbDebug        bool
}

var (
	DB *dbConfig
)

func init() {
	err := beego.LoadAppConfig("ini", "conf/beego.conf")
	bizerror.Check(err)

	//日志配置
	logConf()
	//数据库信息
	dbConf()
	//函数导出信息
	funcConf()

	Log.Sys("goblog starup successful env:%v,appPath:%v", beego.AppConfig.String("runmode"), beego.AppPath)
}

func dbConf() {

	DB = new(dbConfig)
	DB.DbUser = beego.AppConfig.DefaultString("mysqlUser", "root")
	DB.DbPwd = beego.AppConfig.DefaultString("mysqlPass", "lyd")
	DB.DbUrls = beego.AppConfig.DefaultString("mysqlUrls", "localhost:3306")
	DB.DbName = beego.AppConfig.DefaultString("mysqlDb", "goblog")
	DB.DbMaxIdleConns = beego.AppConfig.DefaultInt("mysqlMaxIdleConns", 10)
	DB.DbMaxConns = beego.AppConfig.DefaultInt("mysqlMaxOpenConns", 50)
	DB.DbForce = beego.AppConfig.DefaultBool("mysqlForce", false)
	DB.DbDebug = beego.AppConfig.DefaultBool("mysqlDebug", true)
	orm.Debug = DB.DbDebug

	Log.Sys("call Config init DB:%+v", DB)
}

func logConf() {
	InitLogs(false)
}

func funcConf() {
	cfg := beego.AppConfig
	err := beego.AddFuncMap("getValue", constant.GetValue)
	bizerror.Check(err)
	err = beego.AddFuncMap("appName", func() string { return cfg.String("appname") })
	bizerror.Check(err)
}
