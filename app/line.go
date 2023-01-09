package app

import (
	"github.com/gdamore/tcell/v2"
	"math/rand"
)

type Line struct {
	sheet    *Sheet
	display  *Display
	treeChar string
	rests    []int
}

func NewLine(sheet *Sheet, display *Display) *Line {
	return &Line{display: display, sheet: sheet, treeChar: "🎄"}
}

func (l *Line) create(row int) {
	v := NewVehicle(l.sheet, l.display)
	for {
		v.create(row)
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
		if rand.Intn(4) == 1 {
			l.sheet.cells[row][i] = true
			str += l.treeChar
		} else {
			l.sheet.cells[row][i] = false
			str += " "
		}
	}

	l.display.draw(0, row, style, str)
}
