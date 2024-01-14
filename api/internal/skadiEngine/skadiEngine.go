package skadiEngine

import (
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/ribeirosaimon/skadi/domain/config"
)

var skadi *gin.Engine

func StartSkadiApi() {
	portString := GetProperties().GetString("server.port.src", "0000")

	if err := skadi.Run(fmt.Sprintf(":%s", portString)); err != nil {
		panic(err)
	}
}

func RegisterRouter(routers []*config.SkadiRouter) {
	for _, router := range routers {
		group := skadi.Group(router.Path)
		for index, c := range router.Controllers {
			log.Printf("Add %s to\n", index)

			handlerFunc := gin.HandlerFunc(c.GinFunction)
			group.Handle(c.Method, c.Path, handlerFunc)
		}
	}
}

func init() {
	skadi = gin.Default()
}
