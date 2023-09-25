package main

import "fmt"
import "log"

func main() {
	fmt.Println("Hello, World.")

	var whatToSay string

	var i int

	whatToSay = "Goodbye, cruel world"

	fmt.Println(whatToSay)

	i = 7

	fmt.Println("i is set to", i)

	whatWasSaid, theOtherThingThatWasSaid := saySomething()
	fmt.Println("The function returned", whatWasSaid, theOtherThingThatWasSaid)

	/// Pointer
	var myString string
	myString = "Green"

	log.Println("myString is set to", myString)
	changeUsingPointer(&myString)
	log.Println("after func call myString is set to", myString)

}

func saySomething() (string, string) {
	return "something", "else"
}

func changeUsingPointer(s *string) {
	log.Println("s is set to", s)
	newValue := "Red"
	*s = newValue
}
