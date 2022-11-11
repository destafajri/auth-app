package routes

import (
	"log"
	"net/http"

	"github.com/destafajri/auth-app/applications/middlewares"
	"github.com/destafajri/auth-app/applications/ucase/authentications"
	"github.com/destafajri/auth-app/applications/ucase/products"
	"github.com/destafajri/auth-app/db"
	"github.com/gorilla/mux"
)

func Router() {
	db.DBcon()
	r := mux.NewRouter()

	r.HandleFunc("/login", authentications.Login).Methods("POST")
	r.HandleFunc("/register", authentications.Register).Methods("POST")

	api := r.PathPrefix("/api").Subrouter()
	api.HandleFunc("/products", products.Index).Methods("GET")
	api.Use(middlewares.JWTMiddleware)

	log.Fatal(http.ListenAndServe(":8080", r))
}