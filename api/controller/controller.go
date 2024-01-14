package controller

import "github.com/ribeirosaimon/skadi/domain/config"

var routers []*config.SkadiRouter

func AddRouters(router *config.SkadiRouter) {
	StartRouters()
	routers = append(routers, router)
}

func StartRouters() {
	if routers == nil {
		routers = make([]*config.SkadiRouter, 0)
	}
}

func GetRouters() []*config.SkadiRouter {
	return routers
}
