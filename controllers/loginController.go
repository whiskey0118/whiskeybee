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
	var (
		user models.User
	)
	username := c.GetString("username")
	userExist := user.FindUserByName(username)
	if !userExist {
		result["result"] = "authentication fail"
	} else {
		pass := c.GetString("password")
	}

	c.Data["json"] = &result
	c.ServeJSON()
}

func (c *LoginController) LoginInfo() {
	result := make(map[string]interface{})
	result["result"] = "authentication failure, please login"
	c.Data["json"] = &result
	c.ServeJSON()
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
