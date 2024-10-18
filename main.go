package main

import (
	"fmt"

	"github.com/ramony/gostore/controllers"
)

func main() {
	fmt.Println("Server is running!")
	server := controllers.Router()
	server.Run(":8090")
}
