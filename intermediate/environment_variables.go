package main

import (
	"fmt"
	"os"
)

func envir() {

	user := os.Getenv("USER")
	home := os.Getenv("HOME")

	os.Setenv("FRUIT", "APLLE")

	fmt.Println(user, home)

	for _, e := range os.Environ() {
		fmt.Println(e)
	}
}
