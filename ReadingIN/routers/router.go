package routers

import (
	"ApiTest/ReadingIN/controllers"
	"github.com/astaxie/beego"
)

func init() {
    beego.Router("/", &controllers.MainController{})
	beego.Router("/essay", &controllers.EssayToday{})
}
