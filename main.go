package main

import (
	"fmt"

	"github.com/bxcodec/faker"
)

// SomeStruct ...
type SomeStruct struct {
	FirstName string `faker:"first_name"`
	LastName  string `faker:"last_name"`
}

func main() {
	a := SomeStruct{}

	err := faker.FakeData(&a)

	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Printf("%#v", a)
}
