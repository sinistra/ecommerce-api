package main

import (
	"log"

	"github.com/sinistra/ecommerce-api/app"
)

func main() {

	log.SetFlags(log.LstdFlags | log.Lshortfile)
	app.StartApplication()
}
