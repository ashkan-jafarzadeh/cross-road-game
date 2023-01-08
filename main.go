package main

import (
	"log"
	"math/rand"
	"projects/cross-road-game/app"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())
	logger := app.GetLogger()

	d, err := app.NewDisplay()
	if err != nil {
		log.Fatal(err)
	}

	sheet := app.NewSheet(logger)
	game := app.NewGame(logger, sheet, d)

	defer d.Quit()

	go game.Run()
	d.HandleEvents()
}
