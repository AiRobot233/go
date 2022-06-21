package model

import (
	"strings"
)

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
