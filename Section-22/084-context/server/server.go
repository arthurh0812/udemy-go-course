package main

import (
	"flag"
	"fmt"
	"github.com/Arthur1208/go-pogramming/Section-22/084-context/log"
	"net/http"
	"time"
)

func main() {
	flag.Parse()
	http.HandleFunc("/", log.Decorate(handler))
	panic(http.ListenAndServe(":8080", nil))
}

func handler(w http.ResponseWriter, req *http.Request) {
	ctx := req.Context()

	log.Println(ctx, "Handler started.")
	defer log.Println(ctx, "Handler ended.")

	// fmt.Printf("foo: %v\n", ctx.Value("foo"))

	select {
	case <-time.After(5 * time.Second):
		fmt.Fprintln(w, "hello")
	case <-ctx.Done():
		err := ctx.Err().Error()
		log.Println(ctx, err)
		http.Error(w, err, http.StatusInternalServerError)
	}
}
