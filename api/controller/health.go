package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/ribeirosaimon/skadi/api/internal/entity"
	"github.com/ribeirosaimon/skadi/api/service"
	"github.com/ribeirosaimon/skadi/domain/config"
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
	entity.HandleSuccess(c, 220, healthService.OpenHealth())
}

func Close(c *gin.Context) {
	entity.ThrowError(c, 402, "TESTE DE ERRO")
}

func init() {
	NewHealthController()
}
