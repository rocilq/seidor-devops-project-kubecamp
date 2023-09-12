package main

import (
	"authService/server"

	"os"
)

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		panic("fail to get PORT from env ")
	}
	server := server.Server{
		Port: port,
	}

	server.Start()
}
