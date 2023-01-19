package main

import (
	routes "example.com/server/api"
)

func main() {
	routes.InitGinRoute().Run(":5000")

}
