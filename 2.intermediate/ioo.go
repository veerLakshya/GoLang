package intermediate

import (
	"bytes"
	"fmt"
	"io"
	"log"
	"os"
	"strings"
)

func readFromReader(r io.Reader) {
	buf := make([]byte, 1024)
	n, err := r.Read(buf)
	if err != nil {
		log.Fatalln("error reading from reader. ", err)
	}
	fmt.Println(string(buf[:n]))
}

func writeToWriter(w io.Writer, data string) {
	_, err := w.Write([]byte(data))
	if err != nil {
		log.Fatalln("error writing to writer. ", err)
	}
}

func closeResource(c io.Closer) {
	err := c.Close()
	if err != nil {
		log.Fatalln("error writing to writer. ", err)
	}
}

func bufferExample() {
	var buf bytes.Buffer // value type creation, created in stack
	buf.WriteString("Hello buffer")
	fmt.Println(buf.String())
}

func multiReaderExample() {
	r1 := strings.NewReader("Hello ")
	r2 := strings.NewReader("World!")
	mr := io.MultiReader(r1, r2)
	buf := new(bytes.Buffer) // pointer type creation, created in heap
	_, err := buf.ReadFrom(mr)
	if err != nil {
		log.Fatalln("error reading from mulireader. ", err)
	}
	fmt.Println(buf.String())
}

func pipeEx() {
	pr, pw := io.Pipe()
	go func() {
		pw.Write([]byte("Hello pipe 123"))
		pw.Close()
	}()

	buf := new(bytes.Buffer)
	buf.ReadFrom(pr)
	fmt.Println(buf.String())
}

func writeToFile(filepath string, data string) {
	file, err := os.OpenFile(filepath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatalln("error openning/creating file ", err)
	}
	_, err = file.Write([]byte(data))
	if err != nil {
		log.Fatalln("error writing to file: ", err)
	}
	closeResource(file)

	// writer := io.Writer(file)

	// _, err = writer.Write([]byte("Hello file "))
	// if err != nil {
	// 	log.Fatalln("error writing to file: ", err)
	// }

}

func ioo() {
	fmt.Println("=== read from reader ===")
	readFromReader(strings.NewReader("Hello Reader!"))

	fmt.Println("=== Write to writer ===")
	var writer bytes.Buffer
	writeToWriter(&writer, "Hello writter")
	fmt.Println(writer.String())

	fmt.Println("=== Buffer example ===")
	bufferExample()

	fmt.Println("=== multireader example ===")
	multiReaderExample()

	fmt.Println("=== pipe example ===")
	pipeEx()

	fmt.Println("=== writeing to file example ===")
	filepath := "io.txt"
	writeToFile(filepath, "new text")

	resource := &myResource{name: "myfile"}
	closeResource(resource)

}

type myResource struct {
	name string
}

func (m myResource) Close() error {
	fmt.Println("closing file", m.name)
	return nil
}

/*
use bufio when we need buffering to improve performance or line by line reading or want to reduce system calls
use io package for basic things
*/
