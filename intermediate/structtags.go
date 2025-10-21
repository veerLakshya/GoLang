package main

import (
	"encoding/json"
	"fmt"
	"log"
)

type person1 struct {
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name,omitempty"` // omitted only when not present
	Age       int    `json:"-"`                   // always omited evern when its present
}

func structtags() {
	p := person1{FirstName: "lakshya", LastName: "veer singh", Age: 50}

	jsonP, err := json.Marshal(p)

	if err != nil {
		log.Fatalln("error marshalling struct", err)
	}

	fmt.Println(string(jsonP))
}
