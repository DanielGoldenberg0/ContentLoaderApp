package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/gorilla/mux"
)

var pathToData string = "data.json"

// User represents a user
type User struct {
	Username string `json:"username"`
	Email    string `json:"email"`
	ID       string `json:"id"`
}

func main() {
	info()

	r := mux.NewRouter()
	r.HandleFunc("/api/load", loadUsers).Methods("GET")

	buildHandler := http.FileServer(http.Dir("../client/build"))
	r.PathPrefix("/").Handler(buildHandler)

	staticHandler := http.StripPrefix("/static/", http.FileServer(http.Dir("../client/build/static")))
	r.PathPrefix("/static/").Handler(staticHandler)

	http.ListenAndServe(":8080", r)
}

func loadUsers(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")

	data, err := ioutil.ReadFile("./" + pathToData)
	if err != nil {
		fmt.Println(err)
	}

	var users []User

	err = json.Unmarshal(data, &users)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("GET recieved")

	json.NewEncoder(w).Encode(users)
}

func info() {
	fmt.Println("The server is active")
	fmt.Println(" * IP: localhost")
	fmt.Println(" * Port: 8080")
}
