package main

import (
	"evento_microservices/auth/service"
	"log"
)

var appName = "authService"

func main() {
	log.Println("Starting service..", appName)
	service.Run("5000")
}
