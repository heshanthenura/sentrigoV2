package main

import (
	"log"
	"sentrigoV2/engine/internal/api"
)

func main() {
	// go func() {
	// 	if err := engine.StartEngine(); err != nil {
	// 		log.Fatal(err)
	// 	}
	// }()

	if err := api.StartAPI(); err != nil {
		log.Fatal(err)
	}
}
