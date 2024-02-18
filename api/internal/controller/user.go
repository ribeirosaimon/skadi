package controller

import (
	"github.com/graphql-go/graphql"
	"github.com/ribeirosaimon/skadi/api/internal/service"
	"github.com/ribeirosaimon/skadi/domain/graphqlobjects"
)

type user struct {
	userService *service.UserService
}

func (u *user) startController() {
	u.userService = service.NewUserService()
}

func (u *user) GetMutations() map[string]*graphql.Field {

	signUpArgs := graphql.FieldConfigArgument{
		"name":       &graphql.ArgumentConfig{Type: graphql.String},
		"familyName": &graphql.ArgumentConfig{Type: graphql.String},
		"email":      &graphql.ArgumentConfig{Type: graphql.String},
		"password":   &graphql.ArgumentConfig{Type: graphql.String},
	}

	return map[string]*graphql.Field{
		"signUp": {
			Name:    "SignUp",
			Type:    graphqlobjects.SuccessObject,
			Args:    signUpArgs,
			Resolve: u.signUp,
		},
		"signIn": {
			Name:    "SignIn",
			Type:    graphqlobjects.HealthObject,
			Resolve: u.SignIn,
		},
	}
}

func (u *user) GetQueries() map[string]*graphql.Field {
	return nil
}

func (u *user) signUp(p graphql.ResolveParams) (interface{}, error) {
	name := p.Args["name"].(string)
	familyName := p.Args["familyName"].(string)
	email := p.Args["email"].(string)
	password := p.Args["password"].(string)
	_, err := u.userService.SignUp(name, familyName, email, password)
	if err != nil {
		return nil, err
	}
	return graphqlobjects.Resolver()
}

func (u *user) SignIn(p graphql.ResolveParams) (interface{}, error) {
	email := p.Args["email"].(string)
	password := p.Args["password"].(string)

	return u.userService.SignIn(email, password)
}
