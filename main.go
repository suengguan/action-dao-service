package main

import (
	_ "dao-service/action-dao-service/routers"

	"model"

	"github.com/astaxie/beego"
)

func main() {
	err := model.InitEnv()

	if err != nil {
		beego.Debug(err)
		return
	}

	if beego.BConfig.RunMode == "dev" {
		beego.BConfig.WebConfig.DirectoryIndex = true
		beego.BConfig.WebConfig.StaticDir["/swagger"] = "swagger"
	}
	beego.Run()
}
