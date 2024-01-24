package main

import (
	"os"

	"github.com/ribeirosaimon/skadi/api/controller"
	"github.com/ribeirosaimon/skadi/api/internal/skadiEngine"
)

func main() {
	skadiEngine.RegisterRouter(controller.GetRouters())

	skadiEngine.StartSkadiApi(os.Args[1])
}
