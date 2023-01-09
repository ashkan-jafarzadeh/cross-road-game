package app

import (
	"log"
)

func InitGame() {
	logger := GetLogger()

	d, err := NewDisplay()
	if err != nil {
		log.Fatal(err)
	}

	sheet := NewSheet(logger)
	game := NewGame(logger, sheet, d)

	defer d.Quit()

	go game.Run()
	game.HandleEvents()
}
