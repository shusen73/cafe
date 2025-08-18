package main

import (
	"log"
	"os"

	"cafe/db"
	"cafe/router"
)

func main() {
	// Ensure DB is ready (migrated + seeded)
	if _, err := db.Connect(); err != nil {
		log.Fatal(err)
	}

	r := router.Setup()

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	log.Printf("Server listening on :%s", port)
	if err := r.Run(":" + port); err != nil {
		log.Fatal(err)
	}
}
