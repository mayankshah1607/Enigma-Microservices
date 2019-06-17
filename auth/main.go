package main

import (
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/mayankshah1607/Enigma-Microservices/auth/service"
)

var serviceName string
var servicePort string

func init() {
	if err := godotenv.Load(); err != nil {
		log.Println("No .env found!")
	}

	serviceName, _ = os.LookupEnv("AUTH_SERVICE_NAME")
	servicePort, _ = os.LookupEnv("AUTH_SERVICE_PORT")
}

func main() {
	log.Println("Starting service..", serviceName)
	service.Run(servicePort)
}
