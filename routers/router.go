package routers

import (
	"github.com/astaxie/beego"
	"whiskeybee/controllers"
)

func init() {
	beego.Router("/login", &controllers.LoginController{}, "*:Login")
	beego.Router("/loginInfo", &controllers.LoginController{}, "get:LoginInfo")
	beego.Router("/test", &controllers.TestController{}, "*:Get")
	beego.Router("/test1", &controllers.TestController{}, "*:SessionTest")
	beego.Router("/user", &controllers.TestController{}, "*:UserTest")
	beego.Router("/searchUser", &controllers.LoginController{}, "*:SearchUser")

	//beego.Router("/test", &controllers.TestController{}, "*:Get")
	//ns := beego.NewNamespace("/v1",
	//	beego.NSNamespace("/object",
	//		beego.NSInclude(
	//			&controllers.LoginController{},
	//		),
	//	),
	//	beego.NSNamespace("/user",
	//		beego.NSInclude(
	//			&controllers.LoginController{},
	//		),
	//	),
	//)
	//beego.AddNamespace(ns)
}
