package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)
import (
	"github.com/bxcodec/faker"
)

// SomeStruct ...
type SomeStruct struct {
	FirstName string `faker:"first_name"`
	LastName  string `faker:"last_name"`
}

func main() {
	l := 100000
	obj := SomeStruct{}
	var err error
	var data strings.Builder

	for i := 0; i < l; i++ {
		err = faker.FakeData(&obj)

		if err != nil {
			fmt.Printf("Error: %s\n", err.Error())
			break
		}

		data.WriteString(fmt.Sprintf("%s:%s:1000:1000:jun:/home/jun:/bin/bash\n", obj.FirstName, "x"))
	}

	//fmt.Printf("%s\n", data.String())

	if err = ioutil.WriteFile("data.txt", []byte(data.String()), 0640); err != nil {
		fmt.Printf("Error: %s\n", err.Error())
		return
	}
}
