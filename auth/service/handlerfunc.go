package service

import (
	"enigma_microservices/auth/iohandlers"
	"io/ioutil"
	"log"
	"net/http"
)

//SignInHandler handles the /sign-in request
func SignInHandler(w http.ResponseWriter, r *http.Request) {

	//Get []byte from r.Body
	b, err := ioutil.ReadAll(r.Body)
	defer r.Body.Close()
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	//Decode request
	req, err := iohandlers.DecodeSignInRequest(b)
	log.Println("Request body :", req)

	//Encode response to []byte
	resp, err := iohandlers.EncodeResponse(
		iohandlers.AuthResponse{
			Status:  true,
			Message: "Done",
		},
	)
	// Write response
	w.Write(resp)

}
