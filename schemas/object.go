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
		"users": &graphql.Field{
			Type:  graphql.NewList(authUserSchema),
			Resolve:func(p graphql.ResolveParams) (interface{}, error) {

				var users []models.AuthUser

				// Get a QuerySeter object. User is table name
				// Get a QueryBuilder object. Takes DB driver name as parameter
				// Second return value is error, ignored here
				qb, err := orm.NewQueryBuilder("mysql")

				if err != nil{
					beego.Error(err)
					panic(err)
				}
				// Construct query object
				qb.Select("*").
				From("snaphy_auth_user").
				OrderBy("First").Desc().
				Limit(10).Offset(0)
				// export raw query string from QueryBuilder object
				sql := qb.String()
				beego.Info(sql)
				// execute the raw query string
				o := orm.NewOrm()
				o.Raw(sql).QueryRows(&users)
				beego.Info(users)
				return users, nil
			},
		},
	},
})

