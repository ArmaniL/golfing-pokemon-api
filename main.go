package main

import (
	"gen/api"
	"log"
	"os"

	"github.com/labstack/echo/v4"
)

func main() {
	// create a type that satisfies the `api.ServerInterface`, which contains an implementation of every operation from the generated code
	server := api.NewServer()

	e := echo.New()

	api.RegisterHandlers(e, server)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	log.Fatal(e.Start(":" + port))

}
