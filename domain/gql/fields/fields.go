package fields

import (
	"github.com/graphql-go/graphql"

	"github.com/chris-ramon/golang-scaffolding/domain/internal/services"
)

var PingField = &graphql.Field{
	Type: graphql.String,
	Resolve: func(p graphql.ResolveParams) (interface{}, error) {
		return "ok", nil
	},
}

var CurrentUserField = &graphql.Field{
	Type: graphql.String,
	Resolve: func(p graphql.ResolveParams) (interface{}, error) {
		rootValue := p.Info.RootValue.(map[string]interface{})
		srvs, ok := rootValue["services"].(*services.Services)
		if !ok {
			return nil, nil
		}

		currentUser, err := srvs.AuthService.CurrentUser()
		if err != nil {
			return nil, err
		}

		return currentUser.Username, nil
	},
}
