package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
	"github.com/aloksinghal/redi-shard/handlers"
)

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hi there, I love %s!", r.URL.Path[1:])
}

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/ping/", handler)
	r.HandleFunc("/get/{key}",handlers.KeyGetHandler)

}