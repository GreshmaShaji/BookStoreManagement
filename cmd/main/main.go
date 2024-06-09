package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"

	// "gorm.io/driver/mysql"
	"github.com/golangtutorial/BookStoreManagement/pkg/routes"
)

func main() {
	r := mux.NewRouter()
	routes.RegisterBookStoreRoutes(r)
	http.Handle("/", r)
	log.Fatal(http.ListenAndServe("localhost:9010", r))
}
