package main

import "log"

type User struct {
	FirstName string
	LastName  string
}

func main() {

	/// Map
	/// Map is not sorted in order / always in random
	myMap := make(map[string]User)

	/// Interface{} == any in typescript
	// myMaps := make(map[string]interface{})

	me := User{
		FirstName: "Trevor",
		LastName:  "Sawler",
	}

	myMap["me"] = me
	log.Println(myMap["me"].FirstName)

}
