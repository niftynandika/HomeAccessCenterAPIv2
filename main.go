package main

import (
	"log"
	"os"

	"github.com/niftynandika/HomeAccessCenterAPIv2/api"
)

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080" // default if not running on Render or a cloud platform
	}

	app := api.SetupRouter() // Call a function from your api package to get the router
	log.Printf("Starting server on port %s...", port)
	err := app.Run(":" + port)
	if err != nil {
		log.Fatal(err)
	}
}
