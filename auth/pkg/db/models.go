package models

import (
	"context"
	"log"
	"time"

	"golang.org/x/crypto/bcrypt"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/x/bsonx"
)

// User is the Schema
type User struct {
	Email    string
	Name     string
	Password string
}

//CONNECTIONSTRING is the URI of the DG
const CONNECTIONSTRING = "mongodb://localhost:27017"

//DBNAME is the name of DB
const DBNAME = "evento"

//COLLNAME is the name of collection
const COLLNAME = "users"

var db *mongo.Database

func init() {
	client, err := mongo.NewClient(options.Client().ApplyURI(CONNECTIONSTRING))
	if err != nil {
		log.Println(err)
	}
	ctx, cancel := context.WithTimeout(context.Background(), 13*time.Second)
	defer cancel()
	err = client.Connect(ctx)
	if err != nil {
		log.Println("Could not connect to client...")
	}
	db = client.Database("evento")

	_, err = db.Collection("users").Indexes().CreateOne(
		context.Background(),
		mongo.IndexModel{
			Keys:    bsonx.Doc{{"email", bsonx.Int32(1)}},
			Options: options.Index().SetUnique(true),
		},
	)
}

//CreateNew creates a new User in the database
func CreateNew(u *User, c chan error) {
	hash, hashErr := bcrypt.GenerateFromPassword([]byte(u.Password), bcrypt.MinCost)
	if hashErr != nil {
		log.Println("Failed to hash the password..returning")
		c <- hashErr
		return
	}
	log.Println("Password hashed..")
	hashString := string(hash)
	u.Password = hashString
	log.Println("Attempting to insert :", u.Email)
	id, err := db.Collection("users").InsertOne(context.Background(), u)
	if err != nil {
		log.Println("Failed to create new user :", err)
		c <- err
		return
	}
	log.Println("Created new user: ", id.InsertedID)
	c <- err
}

// Authorize logs a user in
func Authorize(email string, plainPwd string, c chan error) {
	var user User
	err := db.Collection("users").FindOne(context.Background(), primitive.D{{"email", email}}).Decode(&user)
	if err != nil {
		log.Println("Failed to read DB :", err)
		c <- err
		return
	}
	log.Println("Result :", user)

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(plainPwd))
	if err != nil {
		log.Println("Error checking hash for ", user.Email, ": ", err)
		c <- err
		return
	}
	log.Println("User ", user.Email, " logged in!")
	c <- nil

}
