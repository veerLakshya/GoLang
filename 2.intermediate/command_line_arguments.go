package intermediate

import (
	"flag"
	"fmt"
	"os"
)

func cmdLineArguments() {

	fmt.Println("command:", os.Args)

	// Define flags
	var name string
	var age int
	var male bool

	flag.StringVar(&name, "name", "john", "Name of the user")
	flag.IntVar(&age, "age", 18, "Age of the user")
	flag.BoolVar(&male, "male", true, "")

	flag.Parse()

	fmt.Println("Name: ", name)
	fmt.Println("Age: ", age)
	fmt.Println("Male: ", male)

}

// ussage - go run command_line_arguments.go -name Jane Doe -age 14

// --help to get them all printed in terminal
