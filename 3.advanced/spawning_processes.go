package advanced

import (
	"fmt"
	"io"
	"os/exec"
)

func spawning_processes() {

	pr, pw := io.Pipe()
	cmd := exec.Command("grep", "foo")
	cmd.Stdin = pr

	go func() {
		defer pw.Close()
		pw.Write([]byte("food is good\nbar\nbaz"))
	}()

	output, err := cmd.Output()
	if err != nil {
		fmt.Println("Error: ", err)
		return
	}
	fmt.Println("Output: ", string(output))

	// cmd := exec.Command("printenv", "SHELL")
	// output, err := cmd.Output()
	// if err != nil {
	// 	fmt.Println("Error: ", err)
	// 	return
	// }
	// fmt.Println("Output: ", string(output))

	// cmd := exec.Command("sleep", "60")
	// err := cmd.Start()

	// if err != nil {
	// 	fmt.Println("error: ", err)
	// 	return
	// }

	// time.Sleep(time.Second * 2)

	// err = cmd.Process.Kill()
	// if err != nil {
	// 	fmt.Println("error killing process: ", err)
	// 	return
	// }

	// fmt.Println("Process killed ")

	// err = cmd.Wait()
	// if err != nil {
	// 	fmt.Println("error: ", err)
	// 	return
	// }
	// fmt.Println("Process is complete")

	// cmd := exec.Command("echo", "Hello world")
	// output, err := cmd.Output()
	// if err != nil {
	// 	panic(err)
	// }
	// println(string(output))

	// cmd := exec.Command("grep", "foo")
	// // Set input for the command
	// cmd.Stdin = strings.NewReader("foo is good\nbar\nbaz\n")

	// output, err := cmd.Output()
	// if err != nil {
	// 	panic(err)
	// }

	// println("output: ", string(output))

}

/*
Process spawning refers to creating and managing separate processes from a parent program.
This is useful for running tasks concurrently, isolating workloads, and managing system resources effectively.

Why use process spawning
	- soncurrency
	- isolation
	- resource management

os/exec package
	- exec.Command
	- cmd.Stdin / cmd.Stdout / cmd.Stderr
	- cmd.Start / cmd.Wait
	- cmd.Output
*/
