package service

import (
	"net/http"

	"github.com/gorilla/context"
	"github.com/mayankshah1607/Enigma-Microservices/submission/iohandlers"
	"github.com/mayankshah1607/Enigma-Microservices/submission/model"
)

//SubmitHandler handles the /submit route
func SubmitHandler(w http.ResponseWriter, r *http.Request) {
	req := context.Get(r, "req")
	c := make(chan iohandlers.SubmissionResponse)
	go model.MakeSubmission(req.(iohandlers.SubmissionRequest), c)

	resp, err := iohandlers.EncodeResponse(<-c)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.Write(resp)
}
