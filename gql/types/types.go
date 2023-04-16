package types

import (
	"github.com/graphql-go/graphql"

	"github.com/chris-ramon/golang-scaffolding/gql/fields"
)

var QueryType = graphql.NewObject(graphql.ObjectConfig{
	Name: "Query",
	Fields: graphql.Fields{
		"ping": fields.PingField,
	},
})
