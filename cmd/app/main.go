package main

import (
	"log"

	"github.com/OpenTabletopHub/TabletopWiki/internal/app"
)

func main() {
	server := app.NewAppServer(":8000")
	if err := server.Run(); err != nil {
		log.Fatal(err)
	}
}
