package routers

import (
	"github.com/astaxie/beego"
)

func init() {

	beego.GlobalControllerRouter["snaphyAPI/controllers:GraphQLController"] = append(beego.GlobalControllerRouter["snaphyAPI/controllers:GraphQLController"],
		beego.ControllerComments{
			"Any",
			`/`,
			[]string{"get"},
			nil})

}
