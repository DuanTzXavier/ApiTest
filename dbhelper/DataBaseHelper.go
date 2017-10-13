package dbhelper

import (
	"database/sql"
	"fmt"
	"container/list"
	_ "github.com/go-sql-driver/mysql"
	"strings"
	"reflect"
)

var (
	dbHsotIP  	= "(127.0.0.1:3306)"//IP地址
	dbUserName 	= "readingin"//用户名
	dbPassword 	= "soccer"//密码
	dbName     	= "users"//数据库名
	connectString = dbUserName + ":" + dbPassword + "@tcp" + dbHsotIP + "/" + dbName + "?charset=utf8"
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
	defer db.Close()
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
	db, err := sql.Open("mysql", connectString)
	checError(err)
	defer db.Close()

	rows, err := db.Query("SELECT user_id FROM users_base_info")
	checError(err)
	defer rows.Close()

	for rows.Next() {
		var uid string
		err = rows.Scan(&uid)
		checError(err)
		fmt.Println(uid)
		fmt.Print(rows.Columns())
	}
}

func QueryBySql(tableName string, filterField []string, filterValue []string, struc interface{}) *[]interface{} {
	query := buildQueryCommand(tableName, filterField)
	var d []interface{}
	for _, value := range filterValue {
		d = append(d, value)
	}

	return processQueryCommand(query, struc, d...)
}

func processQueryCommand(query string, struc interface{}, cond ...interface{}) *[]interface{} {
	db, err := sql.Open("mysql", connectString)
	checError(err)
	defer db.Close()

	stmt, err := db.Prepare(query)
	checError(err)
	defer stmt.Close()

	rows, err := stmt.Query(cond...)
	checError(err)
	defer rows.Close()

	result := make([]interface{}, 0)

	s := reflect.ValueOf(struc).Elem()
	leng := s.NumField()
	onerow := make([]interface{}, leng)
	for i := 0; i < leng; i++ {
		onerow[i] = s.Field(i).Addr().Interface()
	}
	for rows.Next() {
		err = rows.Scan(onerow...)
		if err != nil {
			panic(err)
		}
		result = append(result, s.Interface())
	}

	return &result
}

func buildQueryCommand(tableName string, filterField []string) string {
	var queryCommand string
	queryCommand += "SELECT * FROM " + tableName + " WHERE "
	if filterField != nil {
		queryCommand += buildFilterField(filterField)
	}
	return queryCommand
}

func buildFilterField(filterField []string) string{
	var result string

	for i := 0; i < len(filterField); i++ {
		result += filterField[i] + "=?" + " AND "
	}

	result = strings.TrimRight(result, " AND ")

	return result
}

func checError(err error) {
	if err != nil {
		panic(err)
	}
}
