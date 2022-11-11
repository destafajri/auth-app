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

	r.HandleFunc("/login", authentications.Login).Methods(http.MethodPost)
	r.HandleFunc("/register", authentications.Register).Methods(http.MethodPost)

	api := r.PathPrefix("/api").Subrouter()
	api.Use(middlewares.JWTMiddleware)
	api.HandleFunc("/products", products.Index).Methods(http.MethodGet)

	log.Fatal(http.ListenAndServe(":9000", r))
}