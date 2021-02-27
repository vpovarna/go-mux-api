package main

import (
	"os"

	"github.com/vpovarna/go-mux-api/server"
)

func main() {

	a := server.App{}
	a.Initialize(
		os.Getenv("APP_DB_USERNAME"),
		os.Getenv("APP_DB_PASSWORD"),
		os.Getenv("APP_DB_NAME"))

	a.Run(":18010")
}
