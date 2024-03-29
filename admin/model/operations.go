package model

import (
	"context"
	"log"

	"github.com/mayankshah1607/Enigma-Microservices/admin/iohandlers"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
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

//DeleteQuestion is used by /delete
func DeleteQuestion(id string, c chan iohandlers.AdminResponse) {
	objID, _ := primitive.ObjectIDFromHex(id)
	_, err := db.Collection("questions").DeleteOne(context.TODO(), bson.M{"_id": objID})
	if err != nil {
		log.Println("Failed to delete question : ", err.Error())
		c <- iohandlers.AdminResponse{
			Status:  false,
			Message: "Failed to delete question",
		}
		return
	}
	log.Println("Successfully deleted question")
	c <- iohandlers.AdminResponse{
		Status:  true,
		Message: "Successfully deleted question",
	}
}

//CheckAdmin checks if a given email is admin
func CheckAdmin(e string) bool {
	var user User
	err := db.Collection("users").FindOne(context.TODO(), bson.M{"email": e, "admin": true}).Decode(&user)
	if err != nil {
		return false
	}
	return true
}
