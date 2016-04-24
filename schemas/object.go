package schemas

import (
	"github.com/graphql-go/graphql"
	"snaphyAPI/models"
	"github.com/astaxie/beego/orm"

	"github.com/astaxie/beego"


)

var authUserSchema = graphql.NewObject(graphql.ObjectConfig{
	Name: "AuthUser",
	Fields:graphql.Fields{
		"ID": &graphql.Field{
			Type:graphql.Int,

		},
		"First": &graphql.Field{
			Type:graphql.String,
		},
		"Last": &graphql.Field{
			Type:graphql.String,
		},
		"Email": &graphql.Field{
			Type:graphql.String,
		},
		"Password": &graphql.Field{
			Type:graphql.String,
		},
		"Reg_key": &graphql.Field{
			Type:graphql.String,
		},
		"Reg_date": &graphql.Field{
			Type:graphql.String,
		},
	},
})



var RootQuery = graphql.NewObject(graphql.ObjectConfig{
	Name: "RootQuery",
	Fields:graphql.Fields{
		"latestUser": &graphql.Field{
			Type: authUserSchema,
			Resolve:func(p graphql.ResolveParams) (interface{}, error) {
				//Fetch data from AuthUser
				o := orm.NewOrm()
				auth_user := models.AuthUser{Id:1}
				err := o.Read(&auth_user)
				if err == orm.ErrNoRows {
					beego.Error("No result found.")
				} else if err == orm.ErrMissPK {
					beego.Error("No result found.")
				}
				beego.Info(auth_user)
				return auth_user, nil
			},
		},
	},
})

