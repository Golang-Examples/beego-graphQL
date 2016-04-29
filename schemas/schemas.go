package schemas

import (
	"github.com/graphql-go/graphql"
	"github.com/astaxie/beego"
)

var (
	schema graphql.Schema
)


func GetSchemaObject() (graphql.Schema) {
	/*if schema {
		return schema
	}else{
		beego.Info("Error generating schema. Schama variable is null.Regenerating schema into memory.")
		mySchema, _ := LoadSchema(true)
		schema = mySchema
		return schema
	}*/
	return schema
}


func LoadSchema(loadAll bool) (graphql.Schema, error){
	schemaConfig := graphql.SchemaConfig{Query: dynamicLoadRootQuery(loadAll)}
	mySchema, err := graphql.NewSchema(schemaConfig)
	if err != nil {
		beego.Error("Error generating schema", err)
		panic(err)
	}

	//Now make schema global..
	schema = mySchema
	return schema, err
}

