package controllers

import (
	"github.com/astaxie/beego"
	"whiskeybee/models"
)

type LoginControler struct {
	beego.Controller
}

func (c *LoginControler) Get() {
	user := models.GetUser()
	c.Data["json"] = user
	c.ServeJSON()

}
