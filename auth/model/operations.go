package model

import (
	"context"
	"log"

	"github.com/mayankshah1607/Enigma-Microservices/auth/iohandlers"
)

//CreateUser operation adds a new user in the database
func CreateUser(u iohandlers.SignUpRequest, c chan iohandlers.AuthResponse) {
	obj := User{
		Name:        u.Name,
		Email:       u.Email,
		University:  u.University,
		Password:    u.Password,
		CurQuestion: 1,
	}
	insertRes, err := db.Collection("users").InsertOne(context.TODO(), obj)

	if err != nil {
		c <- iohandlers.AuthResponse{
			Status:  false,
			Message: "Could not create a new user :" + err.Error(),
		}
		return
	}

	log.Println("Successfully inserted document : ", insertRes.InsertedID)
	c <- iohandlers.AuthResponse{
		Status:  true,
		Message: "Successfully created a new user",
	}
}
