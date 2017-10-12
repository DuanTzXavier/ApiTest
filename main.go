package main

import (
	"ApiTest/dbhelper"
	"fmt"
)

func main() {
	fmt.Print("Hello World\n")
	//dbhelper.QuireData()
	var s []string
	s = append(s, "user_name")
	var b []string
	b = append(b, "FirstUser")

	fmt.Print(dbhelper.QueryBySql(s, "users_base_info", s, b))
}
