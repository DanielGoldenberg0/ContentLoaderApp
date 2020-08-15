package main

import (
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/gorilla/mux"
)

var path string = "data.json"

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/api/load", loadUsers).Methods("GET")

	http.ListenAndServe(":8080", r)

	info()
}

func loadUsers(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")

	data, err := ioutil.ReadFile("./" + path)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(data[0])
}

func info() {
	fmt.Println("The server is active")
	fmt.Println(" * IP: localhost")
	fmt.Println(" * Port: 8080")
}
