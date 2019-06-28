package model

import (
	"context"
	"log"

	"github.com/mayankshah1607/Enigma-Microservices/submission/iohandlers"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

//MakeSubmission handles the /submit route
func MakeSubmission(r iohandlers.SubmissionRequest, c chan iohandlers.SubmissionResponse) {
	var q Question
	id, _ := primitive.ObjectIDFromHex(r.QId)
	UserAnswer := r.Answer

	filter := bson.M{"_id": id}
	err := db.Collection("questions").FindOne(context.TODO(), filter).Decode(&q)
	if err != nil {
		log.Println("Failed to find user : ", err.Error())
		c <- iohandlers.SubmissionResponse{
			Status:  false,
			Message: "Invalid Question",
		}
		return
	}

	RealAnswer := q.Answer

	if UserAnswer == RealAnswer {
		log.Println("Answer cracked")
		c <- iohandlers.SubmissionResponse{
			Status:  true,
			Message: "Correct Answer",
		}
		return
	} else {
		log.Println("Could not solve")
		c <- iohandlers.SubmissionResponse{
			Status:  false,
			Message: "Wrong Answer",
		}
		return
	}
}
