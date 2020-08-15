package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/gorilla/mux"
)

var path string = "data.json"

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

	http.ListenAndServe(":8080", r)
}

func loadUsers(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")

	data, err := ioutil.ReadFile("./" + path)
	if err != nil {
		fmt.Println(err)
	}

	var users []User

	err = json.Unmarshal(data, &users)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(users)
}

func info() {
	fmt.Println("The server is active")
	fmt.Println(" * IP: localhost")
	fmt.Println(" * Port: 8080")
}
