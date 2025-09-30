package main

import (
	"fmt"
	"os"
)

func CheckError(err error) {
	if err != nil {
		panic(err)
	}
}

func temporaryFiles() {
	file, err := os.CreateTemp("", "temporaryFile")
	CheckError(err)
	fmt.Println("Temp file created", file.Name())

	defer file.Close()
	defer os.Remove(file.Name())

}
