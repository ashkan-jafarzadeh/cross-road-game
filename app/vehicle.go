package app

import (
	"context"
	"strings"
	"sync"
	"time"
)

var mapMutex = sync.RWMutex{}

type Vehicle struct {
	sheet   *Sheet
	display *Display
	char    string
	queue   chan struct{}
}

func NewVehicle(sheet *Sheet, display *Display) *Vehicle {
	return &Vehicle{display: display, sheet: sheet, char: "🚋"}
}

func (v *Vehicle) create(row int, speed time.Duration, vehicleLength int, startPoint int, ctx context.Context) {
	c := startPoint
	v.fillCellsForFirstInit(row, c, vehicleLength)

	for c < Cols+vehicleLength {
		select {
		case <-ctx.Done():
			return
		case <-time.After(speed):
			l := vehicleLength
			x := c - l
			if c < l {
				l = c
				x = 0
			}
			mapMutex.Lock()
			v.sheet.cells[row][x+l-1] = true

			v.display.draw(x, row, v.display.defStyle, strings.Repeat(v.char, l))
			if x != 0 {
				v.display.draw(x-1, row, v.display.defStyle, " ")

				v.sheet.cells[row][x-1] = false
			}
			mapMutex.Unlock()
			c++
		}

	}
}

func (v *Vehicle) fillCellsForFirstInit(row, c, vehicleLength int) {
	if c != 0 {
		p := c - vehicleLength
		if p < 0 {
			p = 0
		}
		for i := p; i < c; i++ {
			v.sheet.cells[row][i] = true
		}
	}
}
