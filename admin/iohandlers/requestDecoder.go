package iohandlers

import "encoding/json"

//SubmitRequest is a struct for /submit request
type SubmitRequest struct {
	Text   string `json:"text"`
	Image  string `json:"image_url"`
	Answer string `json:"answer"`
}

//DeleteRequest is a struct for /delete request
type DeleteRequest struct {
	ID string `json:"id"`
}

//DecodeSubmitRequest decodes the incoming data as JSON to SubmitRequest struct
func DecodeSubmitRequest(data []byte) (r SubmitRequest, e error) {
	var req SubmitRequest
	err := json.Unmarshal(data, &req)

	if err != nil {
		return SubmitRequest{}, err
	}

	return req, nil
}

//DecodeDeleteRequest decodes the incoming data as JSON to SignUpRequest struct
func DecodeDeleteRequest(data []byte) (r DeleteRequest, e error) {
	var req DeleteRequest
	err := json.Unmarshal(data, &req)

	if err != nil {
		return DeleteRequest{}, err
	}

	return req, nil

}
