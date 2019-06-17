package model

import (
	"context"
	"log"

	"golang.org/x/crypto/bcrypt"

	"github.com/mayankshah1607/Enigma-Microservices/auth/iohandlers"
)

//CreateUser operation adds a new user in the database
func CreateUser(u iohandlers.SignUpRequest, c chan iohandlers.AuthResponse) {
	//Hash password
	hash, err := bcrypt.GenerateFromPassword([]byte(u.Password), bcrypt.MinCost)
	if err != nil {
		log.Println("Error hashing password")
		c <- iohandlers.AuthResponse{
			Status:  false,
			Message: "Error hashing the password : " + err.Error(),
		}
		return
	}

	obj := User{
		Name:        u.Name,
		Email:       u.Email,
		University:  u.University,
		Password:    string(hash),
		CurQuestion: 1,
	}
	insertRes, err := db.Collection("users").InsertOne(context.TODO(), obj)

	if err != nil {
		log.Println("Error creating new user : ", err.Error())
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
