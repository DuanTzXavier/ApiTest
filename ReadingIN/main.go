package main

import (
	_ "ApiTest/ReadingIN/routers"
	"github.com/astaxie/beego"
)

func main() {
	beego.SetViewsPath("ReadingIN/views")
	beego.Run()
}

