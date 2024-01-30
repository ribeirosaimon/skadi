package config

import (
	"github.com/gin-gonic/gin"
	"github.com/ribeirosaimon/skadi/domain/noSql"
)

type SkadiRouter struct {
	Path        string
	Controllers []*skadiController
}

type skadiController struct {
	Method      string
	Path        string
	GinFunction func(ctx *gin.Context)
	Roles       []noSql.Role
}

func NewSkadiController(method, path string, handlerFunc gin.HandlerFunc, roles ...noSql.Role) *skadiController {
	return &skadiController{Method: method, Path: path, GinFunction: handlerFunc, Roles: roles}
}

func NewSkadiRouter(p string) *SkadiRouter {
	return &SkadiRouter{
		Path: p,
	}
}

func (s *SkadiRouter) AddController(c *skadiController) *SkadiRouter {
	s.Controllers = append(s.Controllers, c)
	return s
}
