package iohandlers

import "encoding/json"

//SubmissionRequest defines the request structure
type SubmissionRequest struct {
	QNo    string `json:"q_no"`
	Answer string `json:"answer"`
}

//DecodeSubmissionRequest decodes the incoming data as JSON to SignInRequest struct
func DecodeSubmissionRequest(data []byte) (r SubmissionRequest, e error) {
	var req SubmissionRequest
	err := json.Unmarshal(data, &req)

	if err != nil {
		return SubmissionRequest{}, err
	}

	return req, nil
}
