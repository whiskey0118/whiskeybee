package controllers

import (
	"crypto/sha1"
	"fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/utils"
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

func SendEmail() error {
	config := `{"username":"1043374586@qq.com","password":"bvlgboktybivbfdc","host":"smtp.qq.com","port":587}`
	email := utils.NewEMail(config)
	email.To = []string{"243495484@qq.com"}
	email.From = "1043374586@qq.com"
	email.Text = "this is beego test"
	err := email.Send()
	return err
}

func (c *RegisterController) RegisterTest() {
	result := make(map[string]interface{})
	err := SendEmail()
	if err != nil {
		result["err"] = err
	} else {
		result["flag"] = "success"
	}
	c.Data["json"] = &result
	c.ServeJSON()
}
