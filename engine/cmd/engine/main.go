package main

import (
	"log"
	"sentrigoV2/engine/internal/engine"
)

func main() {
	if err := engine.StartEngine(); err != nil {
		log.Fatal(err)
	}
}
