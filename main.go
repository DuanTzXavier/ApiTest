package main

import (
	"ApiTest/dbhelper"
	"fmt"
	"ApiTest/dbStructure"
	"net/http"
	"log"
	"strings"
	"encoding/json"
	"ApiTest/processAPI"
)

func main() {
	fmt.Print("Hello World\n")

	//http.HandleFunc("/api", TestServer)
	//err := http.ListenAndServe("localhost:8001", nil)
	//if err != nil {
	//	log.Fatal("ListenAndServe: ", err.Error())
	//}

	http.HandleFunc("/essay", GetEssayServer)
	err := http.ListenAndServe("localhost:8001", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err.Error())
	}

}

func GetEssayServer(writer http.ResponseWriter, request *http.Request) {
	processAPI.ProcessHeader(request.Header)
}
func TestServer(writer http.ResponseWriter, request *http.Request) {
	fmt.Println("Inside HelloServer handler")
	var t dbStructure.UserBaseInfo
	t.User_ID = "asa"
	t.User_Name = "name"
	t.User_Avatar = "avator"
	t.User_Status = 1
	result := dbhelper.QueryData("users_base_info", nil, nil, &t)
	fmt.Println(result)
	var paramSlice []string
	for _, param := range *result {
		b, err := json.Marshal(param)
		if err != nil {
			fmt.Println("error:", err)
		}
		paramSlice = append(paramSlice, string(b))
	}
	aa := strings.Join(paramSlice, "")
	fmt.Fprintf(writer, "Hello,"+request.URL.Path[1:] + "test: " + aa)
	fmt.Println(request.Header)
}

func getName(params ...interface{}) {
	var paramSlice []string
	for _, param := range params {
		paramSlice = append(paramSlice, param.(string))
	}
	aa := strings.Join(paramSlice, "_") // Join 方法第2个参数是 string 而不是 rune
	fmt.Println(aa)
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
