package main

import (
	"fmt"

	routes "example.com/server/api/http"
)

func main() {
	fmt.Println("Welcome to the Modular Art API ...")

	router := routes.SetupGinRouter()

	router.Run(":5000")

}
