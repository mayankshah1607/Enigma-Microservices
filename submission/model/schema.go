package model

type Submission struct {
	QNo       int
	Answer    string
	TimeStamp string
}

type Responses struct {
	User        string
	Submissions []Submission
}
