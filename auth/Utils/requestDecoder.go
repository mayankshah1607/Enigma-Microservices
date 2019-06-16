package utils

import "encoding/json"

//SignInRequest defines the struct of sign-in request
type SignInRequest struct {
	Email    string
	Password string
}

//SignUpRequest defines the struct of sign-up request
type SignUpRequest struct {
	Name       string
	Email      string
	University string
	Password   string
}

func decodeSignInRequest(data []byte) (r SignInRequest, e error) {
	var req SignInRequest
	err := json.Unmarshal(data, &req)

	if err != nil {
		return SignInRequest{}, err
	}

	return req, nil
}

func decodeSignUpRequest(data []byte) (r SignUpRequest, e error) {
	var req SignUpRequest
	err := json.Unmarshal(data, &req)

	if err != nil {
		return SignUpRequest{}, err
	}

	return req, nil

}
