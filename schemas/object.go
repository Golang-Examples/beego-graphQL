package schemas

import (
	"github.com/graphql-go/graphql"
	"snaphyAPI/models"
	"github.com/astaxie/beego/orm"

	"github.com/astaxie/beego"

	"time"
)



var typeInterface = graphql.InterfaceConfig{
	Name:"UserInterface",
	Fields: graphql.Fields{
		"ID": &graphql.Field{
			Type:graphql.ID,
		},
	},
}


var authUserSchema = graphql.NewObject(graphql.ObjectConfig{
	Name: "AuthUser",
	Fields:graphql.Fields{
		"ID": &graphql.Field{
			Type:graphql.ID,
		},
		"First": &graphql.Field{
			Type: graphql.NewNonNull(graphql.String),
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



func dynamicAddRootQuery(rootFieldType string) *graphql.Field{
	if(rootFieldType == "latestUser"){
		return &graphql.Field{
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
		}

	}

	return &graphql.Field{
		Type:  graphql.NewList(authUserSchema),

		Resolve:func(p graphql.ResolveParams) (interface{}, error) {
			start := time.Now()
			var users []models.AuthUser
			qb, err := orm.NewQueryBuilder("mysql")

			if err != nil {
				beego.Error(err)
				panic(err)
			}
			// Construct query object
			qb.Select("*").
			From("snaphy_auth_user").
			Where("Email= 'abigailanderson@example.com'").
			OrderBy("Email").Asc().
			Limit(2).Offset(0)

			// export raw query string from QueryBuilder object
			sql := qb.String()
			// execute the raw query string
			o := orm.NewOrm()
			o.Raw(sql).QueryRows(&users)
			beego.Info("Time take to execute query by beego is :", time.Since(start))
			return users, nil
		},
	}

}



func dynamicLoadRootQuery(loadAll bool) *graphql.Object {
	var (
		RootQuery *graphql.Object
	)

	if(loadAll){
		RootQuery = graphql.NewObject(graphql.ObjectConfig{
			Name: "RootQuery",
			Fields:graphql.Fields{
				//load latest users..
				"latestUser": dynamicAddRootQuery("latestUser"),
				//load users..
				"users": dynamicAddRootQuery(""),
			},
		})
	}else{
		beego.Info("I am removing users")
		RootQuery = graphql.NewObject(graphql.ObjectConfig{
			Name: "RootQuery",
			Fields:graphql.Fields{
				//load latest users..
				"latestUser": dynamicAddRootQuery("latestUser"),
				//load users..
				//"users": dynamicAddRootQuery(""),
			},
		})
	}

	return RootQuery
}


