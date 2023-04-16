package fields

import (
	"github.com/graphql-go/graphql"
)

var PingField = &graphql.Field{
	Type: graphql.String,
	Resolve: func(p graphql.ResolveParams) (interface{}, error) {
		return "ok", nil
	},
}
