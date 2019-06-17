package model

import (
	"context"
	"log"
	"os"
	"time"

	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
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

}
