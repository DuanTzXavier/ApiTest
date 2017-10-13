package main

import (
	"ApiTest/dbhelper"
	"fmt"
	"ApiTest/dbStructure"
)

func main() {
	fmt.Print("Hello World\n")
	//dbhelper.QuireData()
	var s []string
	s = append(s, "user_id")
	var b []string
	b = append(b, "123")

	var t dbStructure.UserBaseInfox

	result := dbhelper.QueryBySql("users_base_info", s, b, &t)

	fmt.Println(result)
}
