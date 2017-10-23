package controllers

import (
	"github.com/astaxie/beego"
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
	var essaysInfo readingIN2.EssaysInfo
	resultSlice := dbhelper.QueryData("essays_info", nil, nil, &essaysInfo)
	if len(*resultSlice) > 0 {
		for _, value := range *resultSlice{
			interfaceToStruct(value, &essaysInfo)
		}
	}

	var dbEssayContent readingIN2.EssaysContents
	var filterField []string
	filterField = append(filterField, "essay_id")
	var filterValue []string
	filterValue = append(filterValue, essaysInfo.Essay_ID)
	resultSlice = dbhelper.QueryData("essays_contents", filterField, filterValue, &dbEssayContent)

	var essayContents []readingIN.EssayContent
	if len(*resultSlice) > 0 {
		for _, value := range *resultSlice{

			interfaceToStruct(value, &dbEssayContent)

			var essayContent readingIN.EssayContent
			essayContent.Content = dbEssayContent.Content
			essayContent.ContentBitMap = dbEssayContent.Content_Bit_Map
			essayContent.ContentName = dbEssayContent.Content_Name
			essayContent.Serial = dbEssayContent.Content_Serial

			essayContents = append(essayContents, essayContent)
		}
	}
	var param readingIN.GETEssayResponse
	param.NextID = essaysInfo.Essay_ID
	param.EssayAuthor = essaysInfo.Essay_Author

	param.EssayContents = essayContents
	param.EssayFrom = essaysInfo.Essay_From
	param.EssayWordsCount, _ = strconv.Atoi(essaysInfo.Essay_Words_Count)
	param.EssayName = essaysInfo.Essay_Name
	c.Data["json"] = param

	c.ServeJSON()
}

func interfaceToStruct(from interface{}, toStruct interface{})  {
	b, err := json.Marshal(from)
	if err != nil {
		fmt.Println("error:", err)
	}
	err = json.Unmarshal(b, toStruct)
	if err != nil {
		fmt.Println("error:", err)
	}
}

