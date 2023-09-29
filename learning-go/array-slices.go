package main

import (
	"log"
	"sort"
)

func main() {
	// Array, almost never used in go

	var mySlice []int
	mySlice = append(mySlice, 2)
	mySlice = append(mySlice, 1)
	mySlice = append(mySlice, 3)

	sort.Ints((mySlice))

	log.Println(mySlice)

	numbers := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	log.Println(numbers)

	log.Println(numbers[6:9])

	names := []string{"one", "seven", "fish", "cat"}

	log.Println(names)
}
