package fields

import (
	"errors"

	"github.com/graphql-go/graphql"

	"github.com/chris-ramon/golang-scaffolding/domain/internal/services"
)

func servicesFromResolveParams(p graphql.ResolveParams) (*services.Services, error) {
	rootValue := p.Info.RootValue.(map[string]interface{})
	srvs, ok := rootValue["services"].(*services.Services)

	if !ok {
		return nil, errors.New("invalid services type")
	}

	return srvs, nil
}
