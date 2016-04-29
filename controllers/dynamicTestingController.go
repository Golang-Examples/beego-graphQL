package controllers

import (
	"github.com/astaxie/beego"
	"snaphyAPI/schemas"
)



// Operations about Testing dynamic loading of controllers for generating a schema..
type TestController struct {
	beego.Controller
}



// @router / [get]
func (testCtrl *TestController) GetAll(){
	schemas.LoadSchema(false)
	beego.Info("Succefully removed schemas")
	testCtrl.Data["json"] = "Successfully removed user query dynamically."
	testCtrl.ServeJSON()
}


// Operations about Testing dynamic loading of controllers for generating a schema..
type LoadAllController struct {
	beego.Controller
}



// @router / [get]
func (loadAllCtrl *LoadAllController) GetAll(){
	schemas.LoadSchema(true)
	beego.Info("Succefully loaded all schemas")
	loadAllCtrl.Data["json"] = "Successfully loaded query dynamically."
	loadAllCtrl.ServeJSON()
}




