package model

import (
	"context"
	"log"
	"os"
	"time"

	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/x/bsonx"
)

var client *mongo.Client
var db *mongo.Database
var dbURI string
var dbName string

func init() {

	if err := godotenv.Load(); err != nil {
		log.Println("No .env found!")
	}
	dbURI, _ = os.LookupEnv("DB_URI")

	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(dbURI))

	if err != nil {
		log.Println("Error connecting to database :", err.Error())
		return
	}
	dbName, _ = os.LookupEnv("DB_NAME")
	db = client.Database(dbName)

	//Declaring Unique fields
	_, err = db.Collection("users").Indexes().CreateOne(
		context.Background(),
		mongo.IndexModel{
			Keys:    bsonx.Doc{{"email", bsonx.Int32(1)}},
			Options: options.Index().SetUnique(true),
		},
	)
	if err != nil {
		log.Println("Failed to set unique field !")
	}
}
