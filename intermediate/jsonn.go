package main

import (
	"encoding/json"
	"fmt"
)

type person struct {
	FirstName    string  `json:"first_name"`
	Age          int     `json:"age,omitempty"` // adding omitempty to leave if not provided while marshalling unmarshalling
	EmailAddress string  `json:"emal"`
	Address      address `json:"address"`
}
type address struct {
	City  string `json:"city"`
	State string `json:"state"`
}

func main() {
	person1 := person{FirstName: "john"}
	jsonData, err := json.Marshal(person1)

	if err != nil {
		fmt.Println("error marshalling data: ", err)
		return
	}
	fmt.Println(string(jsonData))

	person2 := person{FirstName: "John", EmailAddress: "sample@example.com", Address: address{City: "new york"}, Age: 0}
	jsonData1, err := json.Marshal(person2)

	if err != nil {
		fmt.Println("error marshalling data: ", err)
		return
	}
	fmt.Println(string(jsonData1))

	jsonData2 := `{"full_name":"Lakshya veer", "emp_id":"009", "age" : 30, "address":{"city": "bareilly", "state": "up"}}`
	var employeeFromJson Employee
	err = json.Unmarshal([]byte(jsonData2), &employeeFromJson)

	if err != nil {
		fmt.Println("error unmarshalling data: ", err)
		return
	}
	fmt.Println(employeeFromJson)
}

type Employee struct {
	FullName string  `json:"full_name"`
	EmpID    string  `json:"emp_id"`
	Age      int     `json:"age"`
	Address  address `json:"address"`
}
