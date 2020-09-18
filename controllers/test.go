package controllers

import (
	"encoding/json"
	"github.com/astaxie/beego"
	"time"
	"whiskeybee/models"
	"whiskeybee/tools"
)

type TestController struct {
	beego.Controller
}

func (c *TestController) Get() {

	result := make(map[string]interface{})
	ud := models.UserRd{
		Username:        "test",
		LoginErrorTimes: 0,
		Status:          "lock",
	}

	uds, _ := json.Marshal(ud)
	tools.RedisConn.Set("user", string(uds), 30*time.Minute)
	result["res"] = "haha"
	result["uds"] = tools.RedisConn.Get("user").Val()
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
	user.Role = ""
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
