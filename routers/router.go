package routers

import (
	"github.com/astaxie/beego"
	"whiskeybee/controllers"
)

func init() {
	beego.Router("/login", &controllers.LoginControler{}, "*:Get")
	beego.Router("/test", &controllers.TestController{}, "*:Get")
}
