package main

import (
	"encoding/xml"
	"fmt"
	"log"
)

type person2 struct {
	XMLName xml.Name `xml:"person"` // wrapper element tag field
	Name    string   `xml:"name"`
	Age     int      `xml:"age"`
	City    string   `xml:"city"`
	Email   string   `xml:"email"`
}

func main() {

	p := person2{Name: "lakshya", Age: 22, City: "blr", Email: "as@asdf.com"}
	xmlData, err := xml.MarshalIndent(p, "", "\t")

	if err != nil {
		log.Fatalln("error marshalling xml: ", err)
	}
	fmt.Println(string(xmlData))

	xmlRaw := `<person><name>lakshya</name><age>22</age></person>`
	var xmlData2 person2
	err = xml.Unmarshal([]byte(xmlRaw), &xmlData2)

	if err != nil {
		log.Fatalln("error marshalling xml: ", err)
	}
	fmt.Println(xmlData2)
}
