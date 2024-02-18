package controller

import (
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/graphql-go/graphql"
	"github.com/ribeirosaimon/skadi/api/internal/repository"
	"github.com/ribeirosaimon/skadi/domain/sql"
)

type SkadiControllerQuery interface {
	GetMutations() map[string]*graphql.Field
	GetQueries() map[string]*graphql.Field
	startController()
}

type skadiRouter struct {
	Queries   graphql.Fields
	Mutations graphql.Fields
}

var (
	sRouters *skadiRouter
	upSince  time.Time
)

func GetRouters() *skadiRouter {
	return sRouters
}

func AddInRouter(c SkadiControllerQuery) {
	c.startController()
	for key, value := range c.GetQueries() {
		sRouters.Queries[key] = value
	}
	for key, value := range c.GetMutations() {
		sRouters.Mutations[key] = value
	}
}

func addRoleInController(controller graphql.FieldResolveFn, roles ...sql.Role) graphql.FieldResolveFn {
	return func(p graphql.ResolveParams) (interface{}, error) {
		if roles != nil {
			func(p graphql.ResolveParams) {
				header := p.Context.Value("header")

				if header, ok := header.(http.Header); ok {
					token := strings.ReplaceAll(header.Get("Authorization"), "Bearer ", "")
					session := sql.Session{}
					repository.GetRepository().
						SqlTemplate().
						CreateNativeQuery(fmt.Sprintf("select s.user_id from session s where s.expired = false and s.token = %s", token),
							&session)
				}
			}(p)
		}
		return controller(p)
	}
}

func init() {
	upSince = time.Now()
	sRouters = &skadiRouter{
		Queries:   graphql.Fields{},
		Mutations: graphql.Fields{},
	}
	AddInRouter(&health{})
	AddInRouter(&user{})
}
