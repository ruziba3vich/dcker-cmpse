package main

import (
	"log"

	"github.com/ruziba3vich/dckr-cmpse/app"
	"github.com/ruziba3vich/dckr-cmpse/app/handler"
	"github.com/ruziba3vich/dckr-cmpse/internal/storage"
)

func main() {
	app := app.New(
		handler.New(
			storage.New(),
		),
	)
	log.Fatal(app.RUN("9000"))
}
