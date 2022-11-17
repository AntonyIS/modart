package main

import (
	"fmt"

	routes "github.com/AntonyIS/modart/api/http"
)

func main() {
	fmt.Println("Welcome to the Modular Art API ...")

	router := routes.SetupGinRouter()

	router.Run(":5000")

}
