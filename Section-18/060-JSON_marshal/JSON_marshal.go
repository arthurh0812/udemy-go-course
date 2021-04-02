package main

import (
	"encoding/json"
	"log"
	"net/http"
)

type person struct {
	Fname string
	Lname string
	Items []string
}

func main() {
	http.HandleFunc("/", foo)
	http.HandleFunc("/marshal", marshal)
	http.HandleFunc("/encode", encode)
	http.ListenAndServe(":8080", nil)
}

func foo(w http.ResponseWriter, req *http.Request) {
	s := `<!DOCTYPE html>
	<html lang="en">
	<head>
	<meta charset="UTF-8">
	<title>FOO</title>
	</head>
	<body>
	You are at foo
	</body>
	</html>`

	w.Write([]byte(s))
}

func marshal(w http.ResponseWriter, req *http.Request) {
	// set headers
	w.Header().Set("Content-Type", "application/json")

	p1 := person{
		Fname: "James",
		Lname: "Bond",
		Items: []string{"Gin", "Suit", "Gun"},
	}
	j, err := json.Marshal(p1)
	if err != nil {
		log.Println(err)
	}
	// write json to response
	w.Write(j)
}

func encode(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	p1 := person{
		Fname: "James",
		Lname: "Bond",
		Items: []string{"Gin", "Suit", "Gun"},
	}

	err := json.NewEncoder(w).Encode(p1)
	if err != nil {
		log.Println(err)
	}
}
