package main

import "fmt"

func main() {
	fmt.Println("Hello, World.")

	var whatToSay string

	var i int

	whatToSay = "Goodbye, cruel world"

	fmt.Println(whatToSay)

	i = 7

	fmt.Println("i is set to", i)

	whatWasSaid := saySomething()
	fmt.Println("The function returned", whatWasSaid)
}

func saySomething() string {
	return "something"
}
