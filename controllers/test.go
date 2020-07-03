package controllers

import "github.com/astaxie/beego"

type TestController struct {
	beego.Controller
	name string
}

func (c *TestController) Get() {

	result := make(map[string]interface{})
	result["flag"] = "successful"
	result["code"] = "200"

	c.Data["json"] = result
	c.ServeJSON()
}

func (c TestController) SessionTest() {
	result := make(map[string]interface{})
	//result["flag"] = "successful"
	//result["code"] = "200"
	//
	//v := c.GetSession("asta")
	//if v ==nil{
	//	c.SetSession("asta", int(1))
	//	result["num"] = 0
	//} else {
	//	c.SetSession("asta",v.(int)+1)
	//	result["num"] = v.(int)
	//}
	c.Data["json"] = result
	c.ServeJSON()

}
