package skadiEngine

import (
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/ribeirosaimon/skadi/domain/config"
)

var (
	skadi       *gin.Engine
	environment string
)

func StartSkadiApi(env string) {
	environment = env
	portString := config.GetPropertiesFile(env).GetString("server.port.src", "0000")

	if err := skadi.Run(fmt.Sprintf(":%s", portString)); err != nil {
		panic(err)
	}
}

func RegisterRouter(routers []*config.SkadiRouter) {
	for _, router := range routers {
		group := skadi.Group(router.Path)
		for index, c := range router.Controllers {
			log.Printf("Add %d to\n", index)

			handlerFunc := gin.HandlerFunc(c.GinFunction)
			group.Handle(c.Method, c.Path, handlerFunc)
		}
	}
}

func GetEnvironment() string {
	return environment
}

func init() {
	skadi = gin.Default()
}
