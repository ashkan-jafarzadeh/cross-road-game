package main

import (
	"math/rand"
	"projects/cross-road-game/app"
	"time"
)

func init()  {
	rand.Seed(time.Now().UnixNano())
}

func main() {
	app.InitGame()
}
