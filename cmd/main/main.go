package main

import (
	"log"
	"net/http"

	"github.com/Deatsilence/go-inf-dictionary/pkg/routes"
	"github.com/gorilla/mux"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

func main() {
	r := mux.NewRouter()
	routes.InfdicRoutes(r)

	http.Handle("/", r)
	log.Fatal(http.ListenAndServe(":8080", r))
}
