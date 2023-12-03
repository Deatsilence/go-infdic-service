package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/Deatsilence/go-infdic-service/pkg/routes"
	"github.com/gorilla/mux"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

func main() {
	r := mux.NewRouter()
	routes.InfdicRoutes(r)

	http.Handle("/", r)

	fmt.Println("Server is running on port 8080")
	log.Fatal(http.ListenAndServe(":8080", r))

}
