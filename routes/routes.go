package routes

import (
	"net/http"

	"github.com/fajrizulfikar/ecommerce-api/controllers"
	"github.com/gorilla/mux"
)

func Routes() http.Handler {
	r := mux.NewRouter()

	r.HandleFunc("/register", controllers.RegisterUser).Methods("POST")

	r.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello world!\n"))
	})

	return r
}
