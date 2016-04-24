package schemas

import (
	"github.com/graphql-go/graphql"
)

var RootQuery = graphql.NewObject(graphql.ObjectConfig{
	Name: "RootQuery",
	Fields:graphql.Fields{
		"latestUser": &graphql.Field{
			Type:graphql.String,
			Resolve:func(p graphql.ResolveParams) (interface{}, error) {
				return "Hello world!", nil
			},
		},
	},
})

