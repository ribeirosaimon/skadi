package main

import (
	"fmt"
	"os"
	"time"

	"github.com/ribeirosaimon/skadi/domain/config"
	"github.com/ribeirosaimon/skadi/domain/nosqldomain"
	"github.com/ribeirosaimon/skadi/domain/repository"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func main() {
	calcularTempo(func() {
		properties := config.GetProperties(os.Args[1])
		skadiRepository := repository.NewSkadiRepository(properties)

		// user := &nosqldomain.User{Name: "teste No Sql"}
		// stock := sqldomain.Stock{Name: "teste stock"}

		v, _ := primitive.ObjectIDFromHex("658767730249eedab563db5d")
		// _, err := skadiRepository.NoSqlTemplate().Save(user)
		teste := &nosqldomain.User{}
		if err := skadiRepository.NoSqlTemplate().FindById(teste, v); err != nil {
			panic(err)
		}

		// skadiRepository.FindById("10")
	})
}

func calcularTempo(f func()) {
	startTime := time.Now()

	f()

	endTime := time.Now()
	duration := endTime.Sub(startTime)
	fmt.Printf("Tempo decorrido: %s\n", duration)
}
