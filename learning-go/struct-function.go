package main

import (
	"log"
)

type myStruct struct {
	FirstName string
}

func (m *myStruct) printFirstName() string {
	/// We can do some bussines logic as part of
	/// myStruct
	return m.FirstName
}

func main() {

	/// Struct with function
	var myVar myStruct
	myVar.FirstName = "John"

	myVar2 := myStruct{
		FirstName: "Mary",
	}

	log.Println("myVar is set to", myVar.printFirstName())
	log.Println("myVar2 is set to", myVar2.printFirstName())
}
