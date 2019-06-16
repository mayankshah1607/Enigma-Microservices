package main

import (
	"enigma_microservices/auth/service"
	"log"
)

var appName = "authService"

func main() {
	log.Println("Starting service..", appName)
	service.Run("5000")
}
