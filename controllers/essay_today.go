package controllers

import (
	"github.com/astaxie/beego"
	"ReadingIN/base/communication/apiStructure/readingIN"
	"ReadingIN/base/db/dbhelper"
	readingIN2 "ReadingIN/base/db/dbStructure/readingIN"
	"encoding/json"
	"fmt"
	"strconv"
	"ReadingIN/base"
	"time"
)

type EssayToday struct {
	beego.Controller
}

func (c *EssayToday) Post() {
	var postEssayRequest readingIN.PostEssayRequest
	json.Unmarshal(c.Ctx.Input.RequestBody, &postEssayRequest)
	var essaysInfo readingIN2.EssaysInfo

	essaysInfo.Essay_ID 		= base.GenerateUniqueCode(10)
	essaysInfo.Essay_Name 		= postEssayRequest.EssayName
	essaysInfo.Essay_Words_Count= 100 //TODO 自动计算Content值
	essaysInfo.Essay_From 		= postEssayRequest.EssayFrom
	essaysInfo.Essay_Author 	= postEssayRequest.EssayAuthor
	essaysInfo.Essay_Creat_Time = strconv.FormatInt(time.Now().UTC().UnixNano(), 10)
	essaysInfo.Essay_Status		= 0
	//essaysInfo.Essay_Tags
	dbhelper.InsertData("essays_info", essaysInfo)

	var essayContent readingIN2.EssaysContents

	essayContent.Essay_ID			= essaysInfo.Essay_ID
	essayContent.Content_ID 		= base.GenerateUniqueCode(10)
	essayContent.Content_Name		= ""
	essayContent.Content			= postEssayRequest.EssayContents
	essayContent.Content_Bit_Map	= 0
	essayContent.Content_Serial		= 0

	dbhelper.InsertData("essays_contents", essayContent)

	var resultMsg readingIN.PostEssayResponse
	resultMsg.ResultCode = 0
	resultMsg.ResultMessage = "保存成功"
	c.Data["json"] = resultMsg
	c.ServeJSON()
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
	param.EssayWordsCount = essaysInfo.Essay_Words_Count
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

