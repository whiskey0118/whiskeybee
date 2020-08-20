package controllers

import (
	"encoding/json"
	"github.com/astaxie/beego"
)

type RegisterController struct {
	beego.Controller
}

func (c *RegisterController) Register() {
	result := make(map[string]interface{})
	data := map[string]interface{}{"username": ""}
	var (
		err error
	)

	if err = json.Unmarshal(c.Ctx.Input.RequestBody, &data); err != nil {

		result["user"] = data
		result["body"] = c.Ctx.Input.RequestBody
	} else {
		result["err"] = err
	}

	c.Data["json"] = result
	c.ServeJSON()

}
