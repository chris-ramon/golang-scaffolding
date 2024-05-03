package fields

import (
	"github.com/graphql-go/graphql"

	"github.com/chris-ramon/golang-scaffolding/domain/auth/mappers"
	"github.com/chris-ramon/golang-scaffolding/domain/gql/types"
	usersMappers "github.com/chris-ramon/golang-scaffolding/domain/users/mappers"
	"github.com/chris-ramon/golang-scaffolding/pkg/ctxutil"
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
	Type: types.CurrentUserType,
	Resolve: func(p graphql.ResolveParams) (interface{}, error) {
		srvs, err := servicesFromResolveParams(p)
		if err != nil {
			return nil, err
		}

		jwtToken, err := ctxutil.AuthHeaderValueFromCtx(p.Context)
		if err != nil {
			return nil, err
		}

		currentUser, err := srvs.AuthService.CurrentUser(p.Context, jwtToken)
		if err != nil {
			return nil, err
		}

		currentUserAPI := mappers.CurrentUserFromTypeToAPI(currentUser)

		return currentUserAPI, nil
	},
}

var AuthUserField = &graphql.Field{
	Type:        types.CurrentUserType,
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
		srvs, err := servicesFromResolveParams(p)
		if err != nil {
			return nil, err
		}

		username, err := fieldFromArgs[string](p.Args, "username")
		if err != nil {
			return nil, err
		}

		password, err := fieldFromArgs[string](p.Args, "password")
		if err != nil {
			return nil, err
		}

		currentUser, err := srvs.AuthService.AuthUser(p.Context, username, password)
		if err != nil {
			return nil, err
		}

		return currentUser, nil
	},
}

var UsersField = &graphql.Field{
	Name: "Users",
	Type: graphql.NewList(types.UserType),
	Resolve: func(p graphql.ResolveParams) (interface{}, error) {
		srvs, err := servicesFromResolveParams(p)
		if err != nil {
			return nil, err
		}

		users, err := srvs.UserService.FindUsers(p.Context)
		if err != nil {
			return nil, err
		}

		usersAPI := usersMappers.UsersFromTypeToAPI(users)

		return usersAPI, nil
	},
}
