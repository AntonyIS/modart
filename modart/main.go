package main

import (
	"fmt"

	routes "github.com/AntonyIS/modart/api/http"
)

func main() {
	fmt.Println("Welcome to the Modular Art API ...")

	r := routes.SetupGinRouter()

	r.Run(":5000")

}
