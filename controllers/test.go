package controllers

import "github.com/astaxie/beego"

type TestController struct {
	beego.Controller
	name string
}

func (this *TestController) Get() {

	result := make(map[string]string)
	result["flag"] = "successful"
	result["code"] = "200"

	this.Data["json"] = result
	this.ServeJSON()
}
