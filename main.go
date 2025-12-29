package main

import (
	"fmt"
	"net/http"
)

func main() {

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "hello from the server, improved")
	})

	fmt.Println("server started on port 3000")
	http.ListenAndServe(":3000", nil)

}
