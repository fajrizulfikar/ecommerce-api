package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	route := mux.NewRouter()
	route.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello world!\n"))
	})
	log.Println("Server started at port 3000")
	log.Fatal(http.ListenAndServe(":3000", route))
}
