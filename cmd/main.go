package main

import (
	"TgAiBot/internal/app"
	"log"
)

func main() {
	a := app.New()
	err := a.Start()
	if err != nil {
		log.Fatal(err)
	}

}
