package service

import (
	"log"
	"net/http"

	"github.com/mayankshah1607/Enigma-Microservices/auth/model"

	"github.com/gorilla/context"
	"github.com/mayankshah1607/Enigma-Microservices/auth/iohandlers"
)

//SignInHandler handles the /sign-in request
func SignInHandler(w http.ResponseWriter, r *http.Request) {

	// req := context.Get(r, "req")
	// c := make(chan iohandlers.AuthResponse)
	// go model.CreateUser(req.(iohandlers.SignUpRequest), c)

	// resp, err := iohandlers.EncodeResponse(<-c)
	// if err != nil {
	// 	w.WriteHeader(http.StatusInternalServerError)
	// 	return
	// }

	// // Write response
	// w.Write(resp)

}

//SignUpHandler handles the /sign-up route
func SignUpHandler(w http.ResponseWriter, r *http.Request) {
	req := context.Get(r, "req")
	log.Println("Request parsed :", req)
	c := make(chan iohandlers.AuthResponse)
	go model.CreateUser(req.(iohandlers.SignUpRequest), c)

	resp, err := iohandlers.EncodeResponse(<-c)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	// Write response
	w.Write(resp)
}
