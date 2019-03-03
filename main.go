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
	l := 100
	data := make([]SomeStruct, l)
	obj := SomeStruct{}
	var err error

	for i := 0; i < l; i++ {
		err = faker.FakeData(&obj)

		if err != nil {
			fmt.Printf("Error: %s\n", err.Error())
			break
		}

		data = append(data, obj)
	}

	fmt.Printf("%#v", data)
}
