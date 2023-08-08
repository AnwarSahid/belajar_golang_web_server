package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
)

var PORT = ":8000"

type Person struct {
	ID      int
	Name    string
	Age     int
	Address string
}

var person = []Person{
	{1, "anwar Sahid", 23, "Sinar Jaya"},
	{2, "anwar2", 24, "Sinar"},
	{3, "anwar3", 25, "Jaya"},
}

func main() {
	http.HandleFunc("/", greet)
	http.HandleFunc("/person", getperson)
	http.HandleFunc("/addperson", addperson)
	http.ListenAndServe(PORT, nil)
}

func greet(w http.ResponseWriter, r *http.Request) {
	msg := "halo word"
	fmt.Fprint(w, msg)
}

func getperson(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	if r.Method == "GET" {
		json.NewEncoder(w).Encode(person)
		return
	}
	http.Error(w, "invalid Method", http.StatusBadRequest)
}

func addperson(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-type", "application/json")

	if r.Method == "POST" {
		name := r.FormValue("name")
		age := r.FormValue("age")
		address := r.FormValue("address")
		converage, err := strconv.Atoi(age)

		if err != nil {
			http.Error(w, "Invalid Age", http.StatusBadRequest)
			return
		}

		newPerson := Person{len(person) + 1, name, converage, address}
		person = append(person, newPerson)

		json.NewEncoder(w).Encode(newPerson)
		return
	}
	http.Error(w, "invalid method", http.StatusBadRequest)
}
