package dbhelper

import (
	"database/sql"
	"fmt"
	"container/list"
)

var (
	dbHsotIP  	= "(127.0.0.1:3306)"//IP地址
	dbUserName 	= "root"//用户名
	dbPassword 	= "123456"//密码
	dbName     	= "Test"//表名
)

func Printxx() {

}

func insert(tableName string, insertData list.List) string {
	if insertData.Len() == 0 {
		return ""
	}
	var sourceName = dbUserName + ":" + dbPassword + "@tcp" + dbHsotIP + "/" + dbName + "?charset=utf8";
	db, err := sql.Open("mysql", sourceName)
	checError(err)
	//插入数据
	var setSouce = "";
	stmt, err := db.Prepare("INSERT " + tableName + " SET " + setSouce)
	checError(err)

	res, err := stmt.Exec("码农", "研发部门", "2016-03-06")
	checError(err)

	affect, err := res.RowsAffected()
	checError(err)

	fmt.Println(affect)

	return ""
}

func checError(err error) {
	if err != nil {
		panic(err)
	}
}
