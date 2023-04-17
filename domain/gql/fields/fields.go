package fields

import (
	"github.com/graphql-go/graphql"

	"github.com/chris-ramon/golang-scaffolding/domain/auth/mappers"
	"github.com/chris-ramon/golang-scaffolding/domain/gql/types"
	"github.com/chris-ramon/golang-scaffolding/domain/internal/services"
)

var PingField = &graphql.Field{
	Name: "Ping",
	Type: graphql.String,
	Resolve: func(p graphql.ResolveParams) (interface{}, error) {
		return "ok", nil
	},
}

var CurrentUserField = &graphql.Field{
	Name: "CurrentUser",
	Type: types.UserType,
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

		currentUserAPI := mappers.CurrentUserFromTypeToAPI(currentUser)

		return currentUserAPI, nil
	},
}

var AuthUserField = &graphql.Field{
	Type:        types.UserType,
	Description: "Authenticates and authorizes an user.",
	Args: graphql.FieldConfigArgument{
		"username": &graphql.ArgumentConfig{
			Type: graphql.NewNonNull(graphql.String),
		},
		"password": &graphql.ArgumentConfig{
			Type: graphql.NewNonNull(graphql.String),
		},
	},
	Resolve: func(p graphql.ResolveParams) (interface{}, error) {
		rootValue := p.Info.RootValue.(map[string]interface{})
		srvs, ok := rootValue["services"].(*services.Services)
		if !ok {
			return nil, nil
		}

		username, ok := p.Args["username"].(string)
		if !ok {
			return nil, nil
		}
		password, ok := p.Args["password"].(string)
		if !ok {
			return nil, nil
		}

		currentUser, err := srvs.AuthService.AuthUser(p.Context, username, password)
		if err != nil {
			return nil, err
		}

		return currentUser, nil
	},
}
