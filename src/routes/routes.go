package routes

import (
	"net/http"

	"github.com/gorilla/mux"
)

func Routes() http.Handler {
	route := mux.NewRouter()

	route.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello world!\n"))
	})

	return route
}
