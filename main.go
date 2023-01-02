package main

import (
	"fmt"

	routes "example.com/modart/api"
)

func main() {
	fmt.Println("Welcome to the Modular Art API ...")
	r := routes.SetupGinRouter()
	r.Run(":5000")

}
