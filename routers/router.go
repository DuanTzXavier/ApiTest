package routers

import (
	"ReadingIN/controllers"
	"github.com/astaxie/beego"
)

func init() {
    beego.Router("/", &controllers.MainController{})
	beego.Router("/essay", &controllers.EssayToday{})
	beego.Router("/essay/:id", &controllers.EssayByID{})
	beego.Router("/essayserial/:serial", &controllers.EssayBySerial{})
	beego.Router("/essay/random", &controllers.EssayRandom{})
}
