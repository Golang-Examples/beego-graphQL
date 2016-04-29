package routers

import (
	"github.com/astaxie/beego"
)

func init() {

	beego.GlobalControllerRouter["snaphyAPI/controllers:GraphQLController"] = append(beego.GlobalControllerRouter["snaphyAPI/controllers:GraphQLController"],
		beego.ControllerComments{
			"GetAll",
			`/`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["snaphyAPI/controllers:GraphQLController"] = append(beego.GlobalControllerRouter["snaphyAPI/controllers:GraphQLController"],
		beego.ControllerComments{
			"Post",
			`/`,
			[]string{"post"},
			nil})

	beego.GlobalControllerRouter["snaphyAPI/controllers:LoadAllController"] = append(beego.GlobalControllerRouter["snaphyAPI/controllers:LoadAllController"],
		beego.ControllerComments{
			"GetAll",
			`/`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["snaphyAPI/controllers:TestController"] = append(beego.GlobalControllerRouter["snaphyAPI/controllers:TestController"],
		beego.ControllerComments{
			"GetAll",
			`/`,
			[]string{"get"},
			nil})

}
