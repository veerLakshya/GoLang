package intermediate

func main() {
	type Person struct {
		firstName string
		lastName  string
		age       int
	}

	// if any field is left without value then default value is assigned automatically
	p := Person{
		firstName: "John",
		lastName: "doe",
		age: 30,
	}

	p1 := Person{
		firstName : "Jane",
		age : 10
	}
}
