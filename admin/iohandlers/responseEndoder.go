package iohandlers

import (
	"encoding/json"
)

//AdminResponse is the structure of response
type AdminResponse struct {
	Status  bool
	Message string
}

//EncodeResponse encodes the AuthResponse
func EncodeResponse(a AdminResponse) (data []byte, e error) {
	b, err := json.Marshal(a)

	if err != nil {
		return nil, err
	}

	return b, nil
}
