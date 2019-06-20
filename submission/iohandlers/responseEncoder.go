package iohandlers

import "encoding/json"

//SubmissionResponse defines the struct of the response
type SubmissionResponse struct {
	Status  bool
	Message string
}

//EncodeResponse encodes the AuthResponse
func EncodeResponse(s SubmissionResponse) (data []byte, e error) {
	b, err := json.Marshal(s)

	if err != nil {
		return nil, err
	}

	return b, nil
}
