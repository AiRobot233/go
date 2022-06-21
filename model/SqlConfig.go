package model

import (
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
	"project/config"
)

const debug = false

func SetOrm() {
	config.Once.Do(func() {
		orm.Debug = debug
		_ = orm.RegisterDataBase("default", "mysql", "root:root(127.0.0.1:3306)/zhifa?charset=utf8")
		orm.RegisterModel(new(SftLawMainBody))
	})
}
