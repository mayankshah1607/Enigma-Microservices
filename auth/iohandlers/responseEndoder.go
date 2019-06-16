package iohandlers

import "encoding/json"

//AuthResponse defines the struct of the response
type AuthResponse struct {
	Status  bool
	Message string
}

//EncodeResponse encodes the AuthResponse
func EncodeResponse(a AuthResponse) (data []byte, e error) {
	b, err := json.Marshal(a)

	if err != nil {
		return nil, err
	}

	return b, nil
}
