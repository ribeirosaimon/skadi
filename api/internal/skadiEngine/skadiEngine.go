package skadiEngine

import (
	"context"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/graphql-go/graphql"
	"github.com/graphql-go/handler"
	"github.com/ribeirosaimon/skadi/api/internal/controller"
	"github.com/ribeirosaimon/skadi/domain/config"
)

var (
	skadi       *gin.Engine
	environment string
)

func StartSkadiApi(env string) {
	environment = env
	portString := config.GetPropertiesFile(env).GetString("server.port.src", "0000")
	apiVersion := config.GetPropertiesFile(env).GetString("api.version", "v1")

	routers := controller.GetRouters()
	var schema, err = graphql.NewSchema(graphql.SchemaConfig{
		Query:    graphql.NewObject(graphql.ObjectConfig{Name: "Query", Fields: routers.Queries}),
		Mutation: graphql.NewObject(graphql.ObjectConfig{Name: "Mutations", Fields: routers.Mutations}),
	})

	if err != nil {
		panic(err)
	}
	h := handler.New(&handler.Config{
		Schema: &schema,
		Pretty: true,
	})

	skadi.POST(fmt.Sprintf("/api/%s/graphql", apiVersion), GraphqlHandler(h))
	if err := skadi.Run(fmt.Sprintf(":%s", portString)); err != nil {
		panic(err)
	}
}

func GraphqlHandler(h *handler.Handler) gin.HandlerFunc {
	return func(c *gin.Context) {
		httpHeaderMiddleware(h)(c.Writer, c.Request)
		h.ServeHTTP(c.Writer, c.Request)
	}
}

func httpHeaderMiddleware(next *handler.Handler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx := context.WithValue(r.Context(), "header", r.Header)
		next.ContextHandler(ctx, w, r)
	}
}

func init() {
	skadi = gin.Default()
}
