package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

func main() {
	http.HandleFunc("/", foo)
	http.HandleFunc("/encode", encode)
	http.ListenAndServe(":8080", nil)
}

func readFile(path string) (int, []byte, error) {
	file, err := os.Open(path)
	if err != nil {
		return 0, []byte{}, err
	}
	defer file.Close()

	b, err := ioutil.ReadAll(file)
	if err != nil {
		return len(b), b, err
	}

	return len(b), b, nil
}

func foo(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "text/html")

	_, b, err := readFile("foo.html")
	if err != nil {
		log.Fatalf("Error: %v", err)
	}

	fmt.Println(b)

	w.Write(b)
}

func encode(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "application/json")
}
