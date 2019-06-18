package service

import (
	"net/http"

	"github.com/gorilla/context"
	"github.com/mayankshah1607/Enigma-Microservices/admin/iohandlers"
	"github.com/mayankshah1607/Enigma-Microservices/admin/model"
)

//SubmitHandler handles the /sign-in request
func SubmitHandler(w http.ResponseWriter, r *http.Request) {
	req := context.Get(r, "req")
	c := make(chan iohandlers.AdminResponse)
	go model.CreateQuestion(req.(iohandlers.SubmitRequest), c)

	resp, err := iohandlers.EncodeResponse(<-c)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.Write(resp)
}
