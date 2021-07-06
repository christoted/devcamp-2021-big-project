package gql

import "github.com/graphql-go/graphql"

type SchemaWrapper struct {
	Schema graphql.Schema
}

func NewSchemaWrapper() *SchemaWrapper {
	return &SchemaWrapper{}
}

func (s *SchemaWrapper) Init() error {
	schema, err := graphql.NewSchema(graphql.SchemaConfig{
		Query : graphql.NewObject(graphql.ObjectConfig{
			Name : "Query",
			
		})
	})
}