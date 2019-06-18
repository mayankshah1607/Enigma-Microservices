package service

import (
	"net/http"

	"github.com/gorilla/context"
	"github.com/mayankshah1607/Enigma-Microservices/admin/iohandlers"
	"github.com/mayankshah1607/Enigma-Microservices/admin/model"
)

//SubmitHandler handles the /submit request
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

//DeleteHandler handles the /delete route
func DeleteHandler(w http.ResponseWriter, r *http.Request) {
	req := context.Get(r, "req")
	c := make(chan iohandlers.AdminResponse)

	id := req.(iohandlers.DeleteRequest).ID
	go model.DeleteQuestion(id, c)

	resp, err := iohandlers.EncodeResponse(<-c)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.Write(resp)
}
