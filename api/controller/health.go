package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/ribeirosaimon/skadi/api/repository"
	"github.com/ribeirosaimon/skadi/api/service"
	"github.com/ribeirosaimon/skadi/domain/config"
	"github.com/ribeirosaimon/skadi/domain/sqldomain"
)

const healthPath = "/health"

var healthService = service.NewHealthService()

func NewHealthController() {

	router := config.NewSkadiRouter(healthPath).
		AddController(config.NewSkadiController(http.MethodGet, "/open", Open)).
		AddController(config.NewSkadiController(http.MethodGet, "/close", Close))

	AddRouters(router)
}

func Open(c *gin.Context) {
	var stocks []sqldomain.Stock
	err := repository.GetRepository().SqlTemplate().FindAll(&stocks)
	if err != nil {
		panic(err)
	}
	c.JSON(200, healthService.OpenHealth())
}

func Close(c *gin.Context) {
	c.JSON(200, healthService.CloseHealth())
}

func init() {
	NewHealthController()
}
