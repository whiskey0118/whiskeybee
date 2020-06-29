package controllers

import (
	"fmt"
	"github.com/astaxie/beego"
)

type BaseController struct {
	beego.Controller
}

func test() {
	a := 1
	fmt.Println(a)
}
