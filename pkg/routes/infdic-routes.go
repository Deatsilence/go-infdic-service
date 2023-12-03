package routes

import (
	"github.com/gorilla/mux"
)

var InfdicRoutes = func(router *mux.Router) {
	router.HandleFunc("/dictionary", controllers.GetDictionary).Methods("GET")
}
