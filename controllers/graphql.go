package controllers

import (
	gqlhandler "github.com/graphql-go/graphql-go-handler"
	"snaphyAPI/schemas"
	"github.com/astaxie/beego"
	 _ "github.com/graphql-go/graphql"
	 _ "github.com/graphql-go/handler"
	"github.com/graphql-go/graphql"
)


// Operations about Users
type GraphQLController struct {
	beego.Controller
}


// @Title any
// @Description GraphQL for all data type
// @Success 200 {string}
// @router /
func (graphCtrl *GraphQLController) Any(){
	schema, err := schemas.GetSchemaObject()
	if err != nil{
		beego.Trace("Error getting object")
		panic("Error getting object")
	}

	opts := gqlhandler.NewRequestOptions(graphCtrl.Ctx.Request)
	params := graphql.Params{Schema: schema, RequestString: opts.Query}
	responseData := graphql.Do(params)
	if len(responseData.Errors) > 0 {
		beego.Trace("Failed to execute graphql operation", responseData.Errors)
		panic(responseData.Errors)
	}

	beego.Info(responseData.Data)
	graphCtrl.Data["json"] = responseData.Data
	graphCtrl.ServeJSON()
}
