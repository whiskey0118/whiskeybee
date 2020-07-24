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
	//username := c.GetString("username")
	//password := c.GetString("password")
	result["test"] = "testhjaha"

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
