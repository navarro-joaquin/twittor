package main

import (
	"log"
	"github.com/navarro-joaquin/twittor/handlers"
	"github.com/navarro-joaquin/twittor/bd"
)

func main() {
	if bd.ChequeoConnection() == 0 {
		log.Fatal("Sin conexión a la BD")
		return
	}
	handlers.Manejadores()
}
