package routers

import (
	"github.com/astaxie/beego"
	"whiskeybee/controllers"
)

func init() {
	beego.Router("/login", &controllers.LoginControler{}, "*:Get")
	beego.Router("/test", &controllers.TestController{}, "*:Get")
	beego.Router("/test1", &controllers.TestController{}, "*:SessionTest")
	beego.Router("/user", &controllers.TestController{}, "*:UserTest")
}
