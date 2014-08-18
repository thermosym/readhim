package routers

import (
	"github.com/thermosym/readhim/controllers"
	"github.com/astaxie/beego"
)

func init() {
    beego.Router("/", &controllers.MainController{})
}
