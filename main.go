package main

import (
	"ApiTest/dbhelper"
	"fmt"
	"ApiTest/dbStructure"
	"net/http"
	"log"
)

func main() {
	fmt.Print("Hello World\n")

	http.HandleFunc("/api", TestServer)
	err := http.ListenAndServe("localhost:8001", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err.Error())
	}

}
func TestServer(writer http.ResponseWriter, request *http.Request) {
	fmt.Println("Inside HelloServer handler")

	fmt.Fprintf(writer, "Hello,"+request.URL.Path[1:] + " header: ")
	fmt.Println(request.Header)
}

func testDB() {
	//dbhelper.QuireData()
	var s []string
	s = append(s, "user_id")
	var b []string
	b = append(b, "user_id")

	var c []string
	c = append(c, "123")
	c = append(c, "asa")

	var t dbStructure.UserBaseInfo
	t.User_ID = "asa"
	t.User_Name = "name"
	t.User_Avatar = "avator"
	t.User_Status = 1

	dbhelper.UpdateData("users_base_info", s, b, c)
	//dbhelper.InsertData("users_base_info", &t)
	//fmt.Print(dbhelper.BuildInsertCommand("users_base_info", &t))
	result := dbhelper.QueryData("users_base_info", nil, nil, &t)

	fmt.Println(result)
}
