package controllers

import (
	"github.com/astaxie/beego"
	"whiskeybee/models"
)

type LoginController struct {
	beego.Controller
}

// @router / [get]
func (c *LoginController) Login() {
	result := make(map[string]interface{})
	v := c.GetSession("asta")
	if v == nil {
		c.SetSession("asta", int(1))
		result["num"] = 0
		result["session"] = c.GetSession("asta")
	} else {
		c.SetSession("asta", v.(int)+1)
		result["num"] = v.(int)
		result["session"] = c.GetSession("asta")
	}

	c.Data["json"] = result
	c.ServeJSON()

}

func (c *LoginController) Register() {

}

//search username whether exist
func (c *LoginController) SearchUser() {
	result := make(map[string]interface{})
	var (
		user models.User
	)
	username := c.GetString("username")
	result["flag"] = user.FindUserByName(username)

	c.Data["json"] = result
	c.ServeJSON()
}
