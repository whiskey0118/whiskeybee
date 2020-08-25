package controllers

import (
	"github.com/astaxie/beego"
	"log"
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
	user.IsEmailVerified = 0
	user.Salt, err = tools.GenerateRandom()
	if err != nil {
		log.Fatal("generate salt fail")
		return
	}
	//user.PasswordHash = sha1.New().Write([]byte("sdf"))
	user.Role = 0
	user.FromLdap = 0
	user.Status = 0

	c.Data["json"] = result
	c.ServeJSON()

}
