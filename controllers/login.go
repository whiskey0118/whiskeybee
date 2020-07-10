package controllers

import (
	"github.com/astaxie/beego"
	"whiskeybee/models"
)

type LoginControler struct {
	beego.Controller
}

func (c *LoginControler) Login() {
	result := make(map[string]interface{})
	//username := c.GetString("username")
	//password := c.GetString("password")

	c.Data["json"] = result
	c.ServeJSON()

}

func (c *LoginControler) Register() {

}

//search username whether exist
func (c *LoginControler) SearchUser(user models.User) {
	result := make(map[string]interface{})
	//o := orm.NewOrm()
	//qs := o.QueryTable(user)
	//result["flag"] = qs.Filter("username",c.GetString("username")).Exist()
	username := c.GetString("username")
	result["flag"] = user.FindUserByName(username)

	c.Data["json"] = result
	c.ServeJSON()
}
