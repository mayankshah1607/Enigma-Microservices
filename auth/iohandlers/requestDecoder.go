package iohandlers

import "encoding/json"

//SignInRequest defines the struct of sign-in request
type SignInRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

//SignUpRequest defines the struct of sign-up request
type SignUpRequest struct {
	Name       string `json:"name"`
	Email      string `json:"email"`
	University string `json:"university"`
	Password   string `json:"password"`
}

//DecodeSignInRequest decodes the incoming data as JSON to SignInRequest struct
func DecodeSignInRequest(data []byte) (r SignInRequest, e error) {
	var req SignInRequest
	err := json.Unmarshal(data, &req)

	if err != nil {
		return SignInRequest{}, err
	}

	return req, nil
}

//DecodeSignUpRequest decodes the incoming data as JSON to SignUpRequest struct
func DecodeSignUpRequest(data []byte) (r SignUpRequest, e error) {
	var req SignUpRequest
	err := json.Unmarshal(data, &req)

	if err != nil {
		return SignUpRequest{}, err
	}

	return req, nil

}
