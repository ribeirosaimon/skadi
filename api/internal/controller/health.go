package controller

import (
	"github.com/graphql-go/graphql"
	"github.com/ribeirosaimon/skadi/api/internal/service"
	"github.com/ribeirosaimon/skadi/domain/graphqlobjects"
	"github.com/ribeirosaimon/skadi/domain/sql"
)

type health struct {
	service *service.HealthServiceStruct
}

func (c *health) startController() {
	c.service = service.NewHealthService()
}

func (c *health) GetQueries() map[string]*graphql.Field {

	return map[string]*graphql.Field{
		"open": {
			Name:    "Open",
			Type:    graphqlobjects.HealthObject,
			Resolve: c.OpenHealthController,
		},
		"close": {
			Name:    "Close",
			Type:    graphqlobjects.HealthObject,
			Resolve: addRoleInController(c.CloseHealthController, sql.ADMINISTRATOR, sql.USER),
		},
	}
}

func (c *health) OpenHealthController(p graphql.ResolveParams) (interface{}, error) {
	return c.service.OpenHealth(), nil
}

func (c *health) CloseHealthController(p graphql.ResolveParams) (interface{}, error) {
	return c.service.CloseHealth(), nil
}

func (c *health) GetMutations() map[string]*graphql.Field {
	return nil
}
