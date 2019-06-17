package service

import (
	"log"
	"net/http"
)

//Run starts the service
func Run(port string) {

	r := NewRouter()
	r.Use(loggingMiddleware)
	r.Use(jsonBodyParser)
	http.Handle("/", r)
	log.Println("Starting HTTP service for auth on PORT: ", port)
	err := http.ListenAndServe(":"+port, nil)

	if err != nil {
		log.Println("Failed to start HTTP service for authService, PORT: ", port)
		log.Println("Error: ", err.Error())
	}
}
