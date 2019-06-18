package service

import (
	"log"

	"github.com/gorilla/mux"
)

//NewRouter initializes a new gorilla mux router
func NewRouter() *mux.Router {
	router := mux.NewRouter().StrictSlash(true)

	log.Println("Setting up routes for adminService..")
	for _, route := range routes {
		router.Methods(route.Method).
			Path(route.Pattern).
			Name(route.Name).
			Handler(route.HandlerFunc)
	}

	return router
}
