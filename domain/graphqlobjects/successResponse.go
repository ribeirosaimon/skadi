package graphqlobjects

import "github.com/graphql-go/graphql"

var SuccessObject = graphql.NewObject(graphql.ObjectConfig{
	Name: "Success",
	Fields: graphql.Fields{
		"success": &graphql.Field{
			Type: graphql.Boolean,
		},
	},
})

func Resolver() (interface{}, error) {
	return map[string]interface{}{
		"success": true,
	}, nil
}
