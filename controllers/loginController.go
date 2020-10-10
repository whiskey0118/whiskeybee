package controllers

import (
	"crypto/md5"
	"fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"log"
	"net/http"
	"time"
	"whiskeybee/models"
	"whiskeybee/tools/helper"
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

	user.Username = c.GetString("username")
	err := user.FindUserByName()
	if err == orm.ErrNoRows {
		result["result"] = "账号或密码错误"
	} else if err == orm.ErrMultiRows {
		result["result"] = "用户名重复"
		log.Fatal("数据库字段重复")
	} else {
		//账号是否被锁定
		//tools.RedisConn.Get("UserRd")
		if true {
			result["result"] = "账号被锁定"
		} else {
			password := c.GetString("password")
			tempHash := fmt.Sprintf("%x", md5.Sum([]byte(password+user.Salt)))
			if tempHash == user.PasswordHash {

				result["result"] = "login successful"
			} else {
				//密码错误计数器
				result["result"] = "账号或密码错误"
			}
		}
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
	result["flag"] = user.FindUserByNameExist(username)

	c.Data["json"] = result
	c.ServeJSON()
}

func (c *LoginController) GetUTCTime() {
	result := make(map[string]interface{})
	result["flag"] = helper.SuccessFlag
	result["time"] = time.Now().UTC().Format(http.TimeFormat)
	c.Data["json"] = &result
	c.ServeJSON()
}
