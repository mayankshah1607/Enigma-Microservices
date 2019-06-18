package model

import (
	"context"
	"log"

	"github.com/mayankshah1607/Enigma-Microservices/admin/iohandlers"
)

//CreateQuestion is used by /submit
func CreateQuestion(q iohandlers.SubmitRequest, c chan iohandlers.AdminResponse) {
	obj := Question{
		Text:     q.Text,
		Answer:   q.Answer,
		Image:    q.Image,
		SolvedBy: 0,
	}

	insertRes, err := db.Collection("questions").InsertOne(context.TODO(), obj)
	if err != nil {
		log.Println("Failed to insert new question :", err.Error())
		c <- iohandlers.AdminResponse{
			Status:  false,
			Message: "Failed to insert new question : " + err.Error(),
		}
		return
	}
	log.Println("Successfully inserted question : ", insertRes.InsertedID)
	c <- iohandlers.AdminResponse{
		Status:  true,
		Message: "Successfully inserted question!",
	}
}

//GetAllQuestions fetches all the questions
func GetAllQuestions() {}

//DeleteQuestion is used by /delete
func DeleteQuestion() {}
