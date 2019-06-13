package models

import (
	"context"
	"log"
	"os"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/joho/godotenv"
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

//CookieModel represents structure of cookie
type CookieModel struct {
	email string
	jwt.StandardClaims
}

//CONNECTIONSTRING is the URI of the DG
var CONNECTIONSTRING string

//DBNAME is the name of DB
var DBNAME string

//COLLNAME is the name of collection
var COLLNAME string

var db *mongo.Database

func init() {
	if err := godotenv.Load(); err != nil {
		log.Print("No .env file found. Services may not work!")
	}

	CONNECTIONSTRING, _ = os.LookupEnv("CONNECTIONSTRING")
	DBNAME, _ = os.LookupEnv("DBNAME")
	COLLNAME, _ = os.LookupEnv("COLLNAME")

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

	_, err = db.Collection(COLLNAME).Indexes().CreateOne(
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
	id, err := db.Collection(COLLNAME).InsertOne(context.Background(), u)
	if err != nil {
		log.Println("Failed to create new user :", err)
		c <- err
		return
	}
	log.Println("Created new user: ", id.InsertedID)
	c <- err
}

// Authorize logs a user in
func Authorize(email string, plainPwd string, c chan error, tk chan string) {
	var user User
	err := db.Collection(COLLNAME).FindOne(context.Background(), primitive.D{{"email", email}}).Decode(&user)
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

	claims := &CookieModel{
		email: email,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(10 * time.Minute).Unix(),
		},
	}

	jwtKey, exists := os.LookupEnv("JWT_KEY")
	if !exists {
		log.Println("JWT_KEY does not exist! Can not write token")
		tk <- ""
		return
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString([]byte(jwtKey))

	if err != nil {
		log.Println("Error writing token : ", err)
		tk <- ""
		return
	}

	log.Println("Wrote token for user ", email)
	tk <- tokenString

}
