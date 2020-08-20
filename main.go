package main

import (
	"fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
	"log"
	"os"
	_ "whiskeybee/routers"
)

type UserTest struct {
	Id   int
	Name string
	Age  int
}

func init() {
	dbUser := beego.AppConfig.String("mysqlUser")
	dbPassword := beego.AppConfig.String("mysqlPassword")
	dbHost := beego.AppConfig.String("mysqlHost")
	dbPort := beego.AppConfig.String("mysqlPort")
	dbName := beego.AppConfig.String("mysqlDb")

	dbLink := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8", dbUser, dbPassword, dbHost, dbPort, dbName)

	//orm.RegisterModel(new(UserTest))
	err := orm.RegisterDriver("mysql", orm.DRMySQL)
	if err != nil {
		log.Fatal("err", err)
		os.Exit(2)
		return
	}
	orm.RegisterDataBase("default", "mysql", dbLink)
	//orm.RunSyncdb("default",false,true)
}

func main() {
	//if beego.BConfig.RunMode == "dev" {
	//	beego.BConfig.WebConfig.DirectoryIndex = true
	//	beego.BConfig.WebConfig.StaticDir["/swagger"] = ""
	//}

	//o := orm.NewOrm()
	//user := UserTest{Id: 1}
	//
	//err := o.Read(&user)
	//fmt.Println("aaa", err, user.Name, user.Age)
	if beego.BConfig.RunMode == "dev" {
		beego.BConfig.WebConfig.DirectoryIndex = true
		beego.BConfig.WebConfig.StaticDir["/swagger"] = "swagger"

	}

	//开启session
	beego.BConfig.WebConfig.Session.SessionOn = true

	//开启登录过滤
	//var FilterUser = func(ctx *context.Context) {
	//	_, ok := ctx.Input.Session("username").(int)
	//	if !ok && ctx.Request.RequestURI != "/login" && ctx.Request.RequestURI != "/loginInfo" && ctx.Request.RequestURI != "register" {
	//		ctx.Redirect(302, "/loginInfo")
	//	}
	//}
	//beego.InsertFilter("/*", beego.BeforeRouter, FilterUser)

	beego.Run()
}
