package main

import (
	"os"

	"github.com/subosito/gotenv"

	server "github.com/ovieokeh/venni-api/server"
)

func main() {
	gotenv.Load()

	environment := os.Getenv("GO_ENV")
	port := os.Getenv("port")
	dbConnectionString := os.Getenv("DB_CONNECTION")

	if environment != "production" {
		dbConnectionString += "?sslmode=disable"
	}

	app := server.Server{}

	app.Init(dbConnectionString)
	app.Run(port)
}
