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

func xmll() {

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

type book struct {
	XMLName xml.Name `xml:"book"`
	ISBN    string   `xml:"isbn,attr"` // attribute of parent element
	Title   string   `xml:"title,attr"`
}

// <book isbn="asdfasd23rf2f2r" color="blue"> asdf </blue>

// best practices - use xml tags, validate xml, handle nested structrues, handle errors, custom marshalling, unmarshalling
// use cases -
// 1. web services and apis(sprinf framework, microsoft .net applications),
// 2. data interchange and storage(rss and atom feeds, electronic data interchange),
// 3. industry standards(healthcare(HL7) and finance(FIXML))
