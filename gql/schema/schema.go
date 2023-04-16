package schema

import (
	"github.com/graphql-go/graphql"

	"github.com/chris-ramon/golang-scaffolding/gql/types"
)

func New() (graphql.Schema, error) {
	return graphql.NewSchema(graphql.SchemaConfig{
		Query: types.QueryType,
	})
}
