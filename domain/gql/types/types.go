package types

import (
	"github.com/graphql-go/graphql"
)

var CurrentUserType = graphql.NewObject(graphql.ObjectConfig{
	Name: "CurrentUserType",
	Fields: graphql.Fields{
		"id": &graphql.Field{
			Description: "The id of the user.",
			Type:        graphql.String,
		},
		"username": &graphql.Field{
			Description: "The username of the user.",
			Type:        graphql.String,
		},
		"jwt": &graphql.Field{
			Description: "The current JWT of the user.",
			Type:        graphql.String,
		},
	},
})

var UserType = graphql.NewObject(graphql.ObjectConfig{
	Name: "UserType",
	Fields: graphql.Fields{
		"id": &graphql.Field{
			Description: "The id of the user.",
			Type:        graphql.String,
		},
		"username": &graphql.Field{
			Description: "The username of the user.",
			Type:        graphql.String,
		},
	},
})
