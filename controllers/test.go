package controllers

import (
	"fmt"
	"github.com/astaxie/beego"
	"time"
	"whiskeybee/models"
)

type TestController struct {
	beego.Controller
}

func (c *TestController) Get() {

	result := make(map[string]interface{})
	var user models.User
	user.Username = c.GetString("username")
	user.Email = c.GetString("email")
	num, err := user.FindUserExistRaw()

	if err != nil {
		result["err"] = err
	} else {
		result["numType"] = fmt.Sprintf("%T", num)
		result["num"] = num
	}

	c.Data["json"] = &result
	c.ServeJSON()
}

func (c *TestController) SessionTest() {
	jsoninfo := c.GetString("jsoninfo")
	if jsoninfo == "" {
		c.Ctx.WriteString("jsoninfo is empty")
		return
	} else {
		c.Ctx.WriteString(jsoninfo)
		return
	}

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
