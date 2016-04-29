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

// @router / [get]
func (graphCtrl *GraphQLController) GetAll(){
	graphQL(graphCtrl)
}


// @router / [post]
func (graphCtrl *GraphQLController) Post(){
	graphQL(graphCtrl)
}





func graphQL(graphCtrl *GraphQLController){
	schema := schemas.GetSchemaObject()
	opts := gqlhandler.NewRequestOptions(graphCtrl.Ctx.Request)
	/*defer func(){
		if r := recover(); r != nil {
			//Schema not loaded dynamically
			beego.Error("Schema not loaded dynamically. Now trying to  loading schema dynamically", r)
			schemas.LoadSchema(true)
			schema = schemas.GetSchemaObject()
			serveJson(graphCtrl, schema, opts)
		}
	}()*/

	serveJson(graphCtrl, schema, opts)
}


func serveJson(graphCtrl *GraphQLController, schema graphql.Schema, opts *gqlhandler.RequestOptions){
	params := graphql.Params{Schema: schema, RequestString: opts.Query}
	responseData := graphql.Do(params)
	if len(responseData.Errors) > 0 {
		beego.Trace("Failed to execute graphql operation", responseData.Errors)
		panic(responseData.Errors)
	}

	//beego.Info(responseData.Data)
	graphCtrl.Data["json"] = responseData.Data
	graphCtrl.ServeJSON()
}
