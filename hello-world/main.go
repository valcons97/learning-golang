package main

import (
	"fmt"
	"net/http"
)

func main() {
	/// HandleFunc use for registration of route/ endpoint of the handler
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		n, err := fmt.Fprintf(w, "Hello, world!")

		if err != nil {
			fmt.Println(err)
		}

		fmt.Println(fmt.Sprintf("Number of bytes written: %d", n))
	})

	/// _ whenever theres error the program wont care
	/// Listen and serve is use to start webserver
	_ = http.ListenAndServe(":8080", nil)
}
