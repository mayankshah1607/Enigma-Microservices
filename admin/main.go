package main

import (
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/mayankshah1607/Enigma-Microservices/admin/service"
)

var serviceName string
var servicePort string

func init() {
	if err := godotenv.Load(); err != nil {
		log.Println("No .env found!")
	}

	serviceName, _ = os.LookupEnv("ADMIN_SERVICE_NAME")
	servicePort, _ = os.LookupEnv("ADMIN_SERVICE_PORT")
}

func main() {
	log.Println("Starting service..", serviceName)
	service.Run(servicePort)
}
