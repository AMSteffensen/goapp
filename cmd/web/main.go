package main

import (
	"fmt"
	"net/http"

	"github.com/AMSteffensen/goapp/pkg/handlers"
)

const portNumber = ":8080"

func main() {
	http.HandleFunc("/", handlers.Home)
	http.HandleFunc("/about", handlers.About)

	fmt.Println(fmt.Sprintf("Starting app on port%s", portNumber))
	if err := http.ListenAndServe(portNumber, nil); err != nil {
		fmt.Println("Error starting server:", err)
	}
}
