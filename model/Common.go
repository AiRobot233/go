package model

import (
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
	"project/config"
	"strings"
)

const debug = false

func SetOrm() {
	config.Once.Do(func() {
		orm.Debug = debug
		_ = orm.RegisterDataBase("default", "mysql", "root:root@(127.0.0.1:3306)/zhifa?charset=utf8")
		orm.RegisterModel(new(SftLawMainBody))
		_ = orm.NewOrm().Using("default")
	})
}

//创建查询语句
func buildSql(tableName string, where [][]string, field string, limit string, deleteTime string) string {
	var str = "SELECT " + field + " FROM " + tableName
	//增加删除时间的字段
	if deleteTime != "" {
		str += " WHERE " + deleteTime + " IS NULL"
	}
	if len(where) != 0 {
		if deleteTime != "" {
			str += " AND "
		} else {
			str += " WHERE "
		}
		for _, value := range where {
			for _, v := range value {
				str += " " + v + " "
			}
			str += " AND "
		}
		str = strings.TrimRight(str, " AND ")
	}
	if limit != "" {
		str += " LIMIT " + limit
	}
	return str
}
