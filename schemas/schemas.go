package schemas

import (
	"github.com/graphql-go/graphql"
)

func GetSchemaObject() (graphql.Schema, error) {
	schemaConfig := graphql.SchemaConfig{Query: RootQuery}
	schema, err := graphql.NewSchema(schemaConfig)
	return schema, err
}

