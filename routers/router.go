package routers

import (
	"github.com/astaxie/beego"
	"whiskeybee/controllers"
)

func init() {
	//beego.Router("/login", &controllers.LoginControler{}, "*:Login")
	//beego.Router("/test", &controllers.TestController{}, "*:Get")
	//beego.Router("/test1", &controllers.TestController{}, "*:SessionTest")
	//beego.Router("/user", &controllers.TestController{}, "*:UserTest")
	//beego.Router("/searchUser", &controllers.LoginControler{}, "*:SearchUser")

	//beego.Router("/login:test", &controllers.LoginControler{}, "*:Login")
	//ns :=
	//	beego.NewNamespace("/v1",
	//			beego.NSInclude(
	//				&controllers.LoginController{},
	//			),
	//	)
	////注册 namespace
	//beego.AddNamespace(ns)

	beego.Router("/test", &controllers.TestController{}, "*:Get")
	ns := beego.NewNamespace("/v1",
		beego.NSNamespace("/object",
			beego.NSInclude(
				&controllers.LoginController{},
			),
		),
		beego.NSNamespace("/user",
			beego.NSInclude(
				&controllers.LoginController{},
			),
		),
	)
	beego.AddNamespace(ns)
}
