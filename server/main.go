package main

import (
	"log"
	"os"

	"cafe/router"
)

func main() {
	publicDir := "server/public"
	if v := os.Getenv("PUBLIC_DIR"); v != "" {
		publicDir = v
	}

	r := router.NewRouter(publicDir)

	addr := ":8080"
	if v := os.Getenv("PORT"); v != "" {
		addr = v
	}

	log.Printf("serving frontend from %q on %s\n", publicDir, addr)
	if err := r.Run(addr); err != nil {
		log.Fatal(err)
	}
}
