package controllers

import (
	"github.com/astaxie/beego"
	"ReadingIN/base/communication/apiStructure/readingIN"
	"ReadingIN/base/db/dbhelper"
	readingIN2 "ReadingIN/base/db/dbStructure/readingIN"
	"strconv"
	"ReadingIN/base/communication/apiStructure/fail"
)

type EssayBySerial struct {
	beego.Controller
}

func (c *EssayBySerial) Get() {
	var filterField []string
	filterField = append(filterField, "essay_serial")
	var filterValue []string
	filterValue = append(filterValue, c.Ctx.Input.Param(":serial"))

	var essayTodayIssue readingIN2.EssayTodayIssue
	resultSlice := dbhelper.QueryData("essay_today_issue", filterField, filterValue, &essayTodayIssue)
	if len(*resultSlice) > 0 {
		for _, value := range *resultSlice{
			interfaceToStruct(value, &essayTodayIssue)
		}
	}else {
		var failMsg fail.FailStructure
		failMsg.ResultMessage = "文章丢失啦"
		failMsg.ResultCode = 1403
		c.Data["json"] = failMsg
		c.ServeJSON()
		return
	}

	filterField = nil
	filterField = append(filterField, "essay_id")
	filterValue = nil
	filterValue = append(filterValue, essayTodayIssue.Essay_ID)

	var essaysInfo readingIN2.EssaysInfo
	resultSlice = dbhelper.QueryData("essays_info", filterField, filterValue, &essaysInfo)
	if len(*resultSlice) > 0 {
		for _, value := range *resultSlice{
			interfaceToStruct(value, &essaysInfo)
		}
	}else {
		var failMsg fail.FailStructure
		failMsg.ResultMessage = "文章丢失啦"
		failMsg.ResultCode = 1404
		c.Data["json"] = failMsg
		c.ServeJSON()
		return
	}

	var dbEssayContent readingIN2.EssaysContents
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
	}else {
		var failMsg fail.FailStructure
		failMsg.ResultMessage = "文章丢失啦"
		failMsg.ResultCode = 1405
		c.Data["json"] = failMsg
		c.ServeJSON()
		return
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

