package main

import (
	"log"

	"github.com/valcons97/learning-golang/helpers"
)

func main() {
	log.Println("hello")
	var myVar helpers.SomeType

	myVar.TypeName = "Some Name"

	log.Println(myVar.TypeName)
}
