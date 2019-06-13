package models

import (
	"context"
	"fmt"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// User is the Schema
type User struct {
	Email    string
	name     string
	password string
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
		log.Fatalln(err)
	}
	ctx, _ := context.WithTimeout(context.Background(), 13*time.Second)
	err = client.Connect(ctx)
	db = client.Database("evento")
}

//CreateNew creates a new User in the database
func CreateNew(u User) error {
	fmt.Println("Attempting to insert ", u)
	_, err := db.Collection("users").InsertOne(context.Background(), u)
	if err != nil {
		return err
	}
	return nil
}
