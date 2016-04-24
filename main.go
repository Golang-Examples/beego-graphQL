package main

import (
	_ "snaphyAPI/docs"
	_ "snaphyAPI/routers"
	"github.com/astaxie/beego"
	_ "snaphyAPI/models"
)







func main() {
	if beego.BConfig.RunMode == "dev" {
		beego.BConfig.WebConfig.DirectoryIndex = true
		beego.BConfig.WebConfig.StaticDir["/swagger"] = "swagger"

	}
	configLog()
	addDoc()
	beego.Run()
}

func addDoc()  {
	beego.SetStaticPath("/doc","static")
	beego.SetStaticPath("/css","static/css")
	beego.SetStaticPath("/js","static/js")
}



func configLog(){
	beego.SetLogger("file", `{"filename":"logs/backend.log"}`)
	if beego.BConfig.RunMode != "dev"{
		//Remove data from console..
		beego.BeeLogger.DelLogger("console")
	}


	beego.SetLogFuncCall(true)

}

