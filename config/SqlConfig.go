package config

import (
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
	"project/model"
)

const debug = false

func SetOrm() {
	once.Do(func() {
		orm.Debug = debug
		_ = orm.RegisterDataBase("default", "mysql", "root:root(127.0.0.1:3306)/zhifa?charset=utf8")
		orm.RegisterModel(new(model.SftLawMainBody))
	})
}
