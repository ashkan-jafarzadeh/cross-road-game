package app

import (
	"math/rand"
	"strings"
	"time"
)

type Vehicle struct {
	sheet   *Sheet
	display *Display
	char    string
}

func NewVehicle(sheet *Sheet, display *Display) *Vehicle {
	return &Vehicle{display: display, sheet: sheet, char: "🚋"}
}

func (v *Vehicle) create(row int) {
	duration := time.Microsecond * time.Duration(rand.Intn(10)) * 10000
	vehicleLength := rand.Intn(25)

	c := 0
	for c < Cols+vehicleLength {
		l := vehicleLength
		x := c - l
		if c < l {
			l = c
			x = 0
		}
		v.sheet.cells[row][x+l-1] = true

		v.display.draw(x, row, v.display.defStyle, strings.Repeat(v.char, l))
		if x != 0 {
			v.display.draw(x-1, row, v.display.defStyle, " ")

			v.sheet.cells[row][x-1] = false
		}
		time.Sleep(duration)
		c++
	}
}
