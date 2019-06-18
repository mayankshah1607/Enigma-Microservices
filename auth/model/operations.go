package model

import (
	"context"
	"log"
	"os"
	"time"

	"github.com/dgrijalva/jwt-go"

	"go.mongodb.org/mongo-driver/bson"
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
		Admin:       false,
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

//AuthenticateUser is used by /sign-in route
func AuthenticateUser(r iohandlers.SignInRequest, c chan iohandlers.AuthResponse, tk chan string) {
	var user User
	email := r.Email
	plainText := r.Password
	filter := bson.D{{"email", email}}

	err := db.Collection("users").FindOne(context.TODO(), filter).Decode(&user)
	if err != nil {
		log.Println("Failed to find user : ", email)
		c <- iohandlers.AuthResponse{
			Status:  false,
			Message: "Authentication failed",
		}
		return
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(plainText))
	if err != nil {
		log.Println("Authentication failed for : ", email)
		c <- iohandlers.AuthResponse{
			Status:  false,
			Message: "Authentication failed",
		}
		return
	}

	log.Println(user.Email, " logged in")
	//If successful authentication
	c <- iohandlers.AuthResponse{
		Status:  true,
		Message: "Authentication successful",
	}

	//Create a new JWT Token
	jwtKey, exists := os.LookupEnv("JWT_KEY")
	if !exists {
		log.Println("Failed to find a JWT Key!")
		tk <- ""
		return
	}
	expirationTime := time.Now().Add(12 * time.Minute)
	claims := &Claims{
		Email: email,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString([]byte(jwtKey))
	if err != nil {
		log.Println("Failed to create new token for :", email)
		log.Println(err.Error())
		tk <- ""
		return
	}
	tk <- tokenString
}
