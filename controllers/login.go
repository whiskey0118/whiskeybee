package controllers

import (
	"github.com/astaxie/beego"
)

type LoginControler struct {
	beego.Controller
}

func (c *LoginControler) Login() {
	result := make(map[string]interface{})
	username := c.GetString("username")
	password := c.GetString("password")

	c.Data["json"] = result
	c.ServeJSON()

}
