package controllers

import (
	"github.com/astaxie/beego"
	//"encoding/json"
	//"fmt"
	"ApiTest/ReadingIN/base/communication/apiStructure/readingIN"
	"ApiTest/ReadingIN/base/db/dbhelper"
	readingIN2 "ApiTest/ReadingIN/base/db/dbStructure/readingIN"
	"encoding/json"
	"fmt"
	"strconv"
)

type EssayToday struct {
	beego.Controller
}

func (c *EssayToday) Get() {
	var t readingIN2.EssaysInfo
	result := dbhelper.QueryData("essays_info", nil, nil, &t)




	if len(*result) > 0 {
		for _, value := range *result{
			b, err := json.Marshal(value)
			if err != nil {
				fmt.Println("error:", err)
			}

			err = json.Unmarshal(b, &t)
			if err != nil {
				fmt.Println("error:", err)
			}
		}
	}
	var param readingIN.GETEssayResponse
	param.NextID = t.Essay_ID
	param.EssayAuthor = t.Essay_Author
	param.EssayContent = t.Essay_Content
	param.EssayFrom = t.Essay_From
	param.EssayWordsCount, _ = strconv.Atoi(t.Essay_Words_Count)
	param.EssayName = t.Essay_Name
	c.Data["json"] = param

	c.ServeJSON()
}

