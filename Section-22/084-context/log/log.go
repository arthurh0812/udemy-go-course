package log

import (
	"context"
	"log"
	"math/rand"
	"net/http"
)

type key int

var requestIDKey = key(42)

// Println function logs a message specific to the incoming context of the request.
func Println(ctx context.Context, msg string) {
	id, ok := ctx.Value(requestIDKey).(int64)
	if !ok {
		log.Println("could not find request ID in this context.")
		return
	}
	log.Printf("[%d] %s", id, msg)
}

// Decorate function uses the incoming http.HandlerFunc to return a http-decorated http.HandlerFunc
func Decorate(f http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		ctx := req.Context()
		id := rand.Int63()
		ctx = context.WithValue(ctx, requestIDKey, id)
		f(w, req.WithContext(ctx))
	}
}
