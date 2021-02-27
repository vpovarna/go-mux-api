package main

import (
	"fmt"
	"log"
	"os"

	"github.com/vpovarna/go-mux-api/server"
)

const port = 18010

func main() {

	a := server.App{}
	a.Initialize(
		os.Getenv("APP_DB_USERNAME"),
		os.Getenv("APP_DB_PASSWORD"),
		os.Getenv("APP_DB_NAME"))

	log.Printf("Starting server on port: %d \n", port)
	a.Run(fmt.Sprintf(":%d", port))
}
