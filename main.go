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
	t.User_ID = "asa"
	t.User_Name = "name"
	t.User_Avatar = "avator"
	t.User_Status = 1

	dbhelper.DeleteData("users_base_info", s, b)
	//dbhelper.InsertData("users_base_info", &t)
	//fmt.Print(dbhelper.BuildInsertCommand("users_base_info", &t))
	result := dbhelper.QueryData("users_base_info", nil, nil, &t)

	fmt.Println(result)
}
