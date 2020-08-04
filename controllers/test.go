package controllers

import (
	"github.com/astaxie/beego"
	"time"
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
	var user models.User

	user.Id = 2
	user.Email = "test@admin.com"
	user.PasswordHash = "test"
	user.Username = "whiskey"
	user.CreatedTime = time.Now()
	user.UpdatedTime = time.Now()
	user.Role = 1
	id, err := user.InsertUser()
	if err != nil {
		result["id"] = id
		result["err"] = err
		result["flag"] = "fail"
	} else {
		result["id"] = id
		result["flag"] = "successful"
	}

	c.Data["json"] = result
	c.ServeJSON()
}
