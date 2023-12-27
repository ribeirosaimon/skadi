package main

import (
	"fmt"
	"reflect"
	"time"
)

type User struct {
	Name string
	Age  int
}

func (u User) GetName() string {
	return u.Name
}

type human interface {
	GetName() string
}

func newHuman() human {
	return User{"human", 40}
}

func main() {
	var x float64 = 3.14
	var u User = User{"Bob", 10}
	var uSlice = []User{
		{"1", 1},
		{"2", 2},
		{"3", 3},
	}
	h := newHuman()

	valorDoX := reflect.ValueOf(x)
	typoDoX := reflect.TypeOf(x)

	sliceTypoDoX := reflect.MakeSlice(reflect.SliceOf(typoDoX), 10, 10)

	fmt.Printf("VALOR E TIPO DO X %s ----- %s ------ %s \n", valorDoX, typoDoX, sliceTypoDoX)

	valorDoU := reflect.ValueOf(u)
	typoDoU := reflect.TypeOf(u)
	fmt.Printf("VALOR E TIPO DO U %s ----- %s \n", valorDoU, typoDoU)

	valorDoUSlice := reflect.ValueOf(uSlice)
	typoDoUSlice := reflect.TypeOf(uSlice)
	fmt.Printf("VALOR E TIPO DO U %s ----- %s \n", valorDoUSlice, typoDoUSlice)

	valorDoH := reflect.ValueOf(h)
	typoDoH := reflect.TypeOf(h)
	fmt.Printf("VALOR E TIPO DO U %s ----- %s \n", valorDoH, typoDoH)

	// calcularTempo(func() {
	// 	properties := config.GetProperties(os.Args[1])
	// 	skadiRepository := repository.NewSkadiRepository(properties)
	//
	// 	// user := &nosqldomain.User{Name: "teste No Sql"}
	// 	// stock := sqldomain.Stock{Name: "teste stock"}
	//
	// 	v, _ := primitive.ObjectIDFromHex("658767730249eedab563db5d")
	// 	// _, err := skadiRepository.NoSqlTemplate().Save(user)
	// 	teste := &nosqldomain.User{}
	// 	if err := skadiRepository.NoSqlTemplate().FindById(teste, v); err != nil {
	// 		panic(err)
	// 	}
	//
	// 	// skadiRepository.FindById("10")
	// })
}

func calcularTempo(f func()) {
	startTime := time.Now()

	f()

	endTime := time.Now()
	duration := endTime.Sub(startTime)
	fmt.Printf("Tempo decorrido: %s\n", duration)
}
