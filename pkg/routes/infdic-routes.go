package routes

import (
	"github.com/gorilla/mux"

	"github.com/Deatsilence/go-infdic-service/pkg/controllers"
)

var InfdicRoutes = func(router *mux.Router) {
	router.HandleFunc("/dictionary/detail/{word}", controllers.HandleGetWordDefinition).Methods("GET")
	router.HandleFunc("/dictionary/{word}", controllers.GetDictionary).Methods("GET")
}
