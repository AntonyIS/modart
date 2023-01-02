package main

import (
	routes "example.com/modart/api"
)

func main() {
	r := routes.InitGinRoute()
	r.Run(":5000")

}
