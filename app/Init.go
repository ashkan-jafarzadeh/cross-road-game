package app

import (
	"log"
	"os"
	"syscall"
)

func InitGame() {
	redirectStderr()

	d, err := NewDisplay()
	if err != nil {
		log.Fatal(err)
	}

	sheet := NewSheet()
	game := NewGame(sheet, d)

	defer d.Quit()

	go game.Run()
	game.HandleEvents()
}

// redirectStderr to the file passed in
func redirectStderr() {
	f, _ := os.OpenFile("error.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	err := syscall.Dup2(int(f.Fd()), int(os.Stderr.Fd()))
	if err != nil {
		log.Fatalf("Failed to redirect stderr to file: %v", err)
	}
}
