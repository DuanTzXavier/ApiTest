package dbhelper

import (
	"database/sql"
	"fmt"
	"container/list"
	_ "github.com/go-sql-driver/mysql"
)

var (
	dbHsotIP  	= "(127.0.0.1:3306)"//IP地址
	dbUserName 	= "readingin"//用户名
	dbPassword 	= "soccer"//密码
	dbName     	= "users"//数据库名
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

func QuireData() {
	var sourceName = dbUserName + ":" + dbPassword + "@tcp" + dbHsotIP + "/" + dbName + "?charset=utf8";
	db, err := sql.Open("mysql", sourceName)
	checError(err)

	rows, err := db.Query("SELECT user_id FROM users_base_info")
	checError(err)

	for rows.Next() {
		var uid string
		err = rows.Scan(&uid)
		checError(err)
		fmt.Println(uid)
	}
}

func checError(err error) {
	if err != nil {
		panic(err)
	}
}
