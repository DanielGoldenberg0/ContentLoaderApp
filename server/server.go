package main

import (
	"fmt"
	// "github.com/gorilla/mux"
)

func info() {
	fmt.Println("The server is active")
	fmt.Println(" * IP: localhost")
	fmt.Println(" * Port: 8080")
}

func main() {
	info()
}
