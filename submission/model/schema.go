package model

//Submission is the struct for submissions
type Submission struct {
	QNo       int
	Answer    string
	TimeStamp string
}

//Responses is the struct for responses
type Responses struct {
	User        string
	Submissions []Submission
}

//Question is the struct of a question
type Question struct {
	Text     string
	Image    string
	Answer   string
	SolvedBy int
}
