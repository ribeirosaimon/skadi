package main

import (
	"fmt"

	"github.com/ribeirosaimon/skadi/domain/repository"
)

func main() {
	repository.NewSkadiRepository()
	pings := make(chan string, 1)
	pongs := make(chan string, 1)
	ping(pings, "enviando a mensagem")
	pong(pings, pongs)
	fmt.Println(<-pongs)
}

func ping(pings chan<- string, msg string) {
	pings <- msg
}

func pong(pings <-chan string, pongs chan<- string) {
	msg := <-pings
	pongs <- msg
}
