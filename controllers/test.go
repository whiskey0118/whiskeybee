package controllers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"whiskeybee/models"
)

type TestController struct {
	beego.Controller
}

func (c *TestController) Get() {

	result := make(map[string]interface{})
	result["flag"] = "successful"
	result["code"] = "200"

	c.Data["json"] = result
	c.ServeJSON()
}

func (c *TestController) SessionTest() {
	result := make(map[string]interface{})
	if c.Ctx.GetCookie("user") == "" {
		c.Ctx.SetCookie("user", "admin")
		c.Ctx.WriteString("cookie set successful")
	}
	c.Data["json"] = result
	c.ServeJSON()

}
func (c *TestController) UserTest() {
	result := make(map[string]interface{})
	o := orm.NewOrm()
	user := new(models.User)

	user.Id = 1
	user.Email = "test@admin.com"
	user.PasswordHash = "test"
	id, err := o.Insert(user)
	if err != nil {
		result["id"] = id
		result["err"] = err
	}

	c.Data["json"] = result
	c.ServeJSON()
}
