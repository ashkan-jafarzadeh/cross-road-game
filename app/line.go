package app

import (
	"context"
	"github.com/gdamore/tcell/v2"
	"math/rand"
	"time"
)

type Line struct {
	sheet    *Sheet
	display  *Display
	vehicle  *Vehicle
	treeChar string
	rests    []int
}

func NewLine(sheet *Sheet, display *Display) *Line {
	return &Line{display: display, sheet: sheet, vehicle: NewVehicle(sheet, display), treeChar: "🎄"}
}

func (l *Line) create(row int, ctx context.Context) {
	speed := time.Microsecond * time.Duration(rand.Intn(vehicleSpeedMax-vehicleSpeedMin+1)+vehicleSpeedMin) * 10000
	go l.vehicle.create(row, speed, rand.Intn(Cols/3), rand.Intn(Cols), ctx)

	for {
		select {
		case <-ctx.Done():
			return
		case <-time.After(speed * time.Duration(rand.Intn(vehiclesDelayMax-vehiclesDelayMin+1)+vehiclesDelayMin)):
			go l.vehicle.create(row, speed, rand.Intn(Cols/3), 0, ctx)
		}
	}
}

func (l *Line) resetRests() {
	for _, rest := range l.rests {
		l.createRest(rest)
	}
}

func (l *Line) createRest(row int) {
	l.rests = append(l.rests, row)
	style := tcell.StyleDefault.Background(tcell.ColorReset).Foreground(tcell.ColorReset)
	str := ""
	for i := 0; i < Cols; i++ {
		if rand.Intn(5) == 1 {
			l.sheet.cells[row][i] = true
			str += l.treeChar
		} else {
			l.sheet.cells[row][i] = false
			str += " "
		}
	}

	l.display.draw(0, row, style, str)
}
