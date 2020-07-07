package controllers

import (
	"github.com/astaxie/beego"
)

type LoginControler struct {
	beego.Controller
}

func (c *LoginControler) Get() {
	result := make(map[string]interface{})
	c.Data["json"] = result
	c.ServeJSON()

}
