package controllers

import (
	"crypto/sha1"
	"fmt"
	"github.com/astaxie/beego"
	"log"
	"regexp"
	"time"
	"whiskeybee/models"
	"whiskeybee/tools"
)

type RegisterController struct {
	beego.Controller
}

func (c *RegisterController) Register() {
	result := make(map[string]interface{})
	var (
		user models.User
		err  error
	)

	user.Username = c.GetString("username")
	user.Email = c.GetString("email")
	exist, err := user.FindUserExistRaw()
	if err != nil {
		result["result"] = "error"
		log.Fatal(err)
		return
	}
	if exist {
		result["result"] = "username or email is exist"
		return
	}
	if IsEmail(user.Email) == false {
		result["result"] = "email format error"
		log.Fatal("email format error")
		return
	}

	user.IsEmailVerified = 0
	user.Salt, err = tools.GenerateRandom()
	if err != nil {
		log.Fatal("generate salt fail")
		return
	}
	newSha1 := sha1.New().Sum([]byte(c.GetString("password")))
	user.PasswordHash = fmt.Sprintf("%x", newSha1)
	user.Role = ""
	user.FromLdap = 0
	user.Status = ""
	user.UpdatedTime = time.Now()
	user.CreatedTime = time.Now()
	num, err := user.InsertUser()
	if err != nil {
		log.Fatal(err)
		result["err"] = err
		result["result"] = "fail"
		return
	}
	result["result"] = "successful"
	result["num"] = num

	c.Data["json"] = result
	c.ServeJSON()

}

//邮箱匹配
func IsEmail(str ...string) bool {
	var b bool
	for _, s := range str {
		b, _ = regexp.MatchString("^([a-z0-9_\\.-]+)@([\\da-z\\.-]+)\\.([a-z\\.]{2,6})$", s)
		if false == b {
			return b
		}
	}
	return b
}
