package graphqlobjects

import "github.com/graphql-go/graphql"

var HealthObject = graphql.NewObject(graphql.ObjectConfig{
	Name: "Health",
	Fields: graphql.Fields{
		"isUp":  &graphql.Field{Type: graphql.Boolean},
		"since": &graphql.Field{Type: graphql.DateTime},
	},
})
