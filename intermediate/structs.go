package main

import "fmt"

type Person struct {
	firstName     string
	lastName      string
	age           int
	address       Address //imbedded struct
	PhoneHomeCell         //Anonymouse imbedded struct
}
type PhoneHomeCell struct {
	home string
	cell string
}

type Address struct {
	city    string
	country string
}

func (p Person) fullName() string {
	return p.firstName + " " + p.lastName
}

func (p *Person) increaseAge(x int) {
	p.age += x
}

func structss() {

	// if any field is left without value then default value is assigned automatically
	p := Person{
		firstName: "John",
		lastName:  "doe",
		age:       30,
		address: Address{
			city:    "london",
			country: "U.K.",
		},
		PhoneHomeCell: PhoneHomeCell{
			home: "1234",
			cell: "4321",
		},
	}

	p1 := Person{
		firstName: "Jane",
		age:       10,
	}

	fmt.Println(p.fullName())
	fmt.Println(p1.lastName)

	p1.increaseAge(2)

	fmt.Println(p1.age)

	// Anonymous structs
	user := struct {
		username string
		email    string
	}{
		username: "user123",
		email:    "random@gmail.com",
	}
	user.username = "user1"

}
