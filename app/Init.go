package app

import (
	"fmt"
	"log"
	"os"
	"syscall"
)

func InitGame() {
	redirectStderr()
	setMode()

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

func setMode() {
	var mode string
	fmt.Print("Choose mode: (default: Normal) \n  [0]: Easy \n  [1]: Normal \n  [2]: Hard \n\n  : ")
	fmt.Scanf("%s", &mode)

	if mode == "0" {
		vehicleSpeedMax = 20
		vehicleSpeedMin = 5
		vehiclesDelayMax = 120
		vehiclesDelayMin = 70
	}
	if mode == "2" {
		vehiclesDelayMax = 10
		vehiclesDelayMin = 0
		vehiclesDelayMax = 70
		vehiclesDelayMin = 25
	}
}

// redirectStderr to the file passed in
func redirectStderr() {
	f, _ := os.OpenFile("error.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	err := syscall.Dup2(int(f.Fd()), int(os.Stderr.Fd()))
	if err != nil {
		log.Fatalf("Failed to redirect stderr to file: %v", err)
	}
}
