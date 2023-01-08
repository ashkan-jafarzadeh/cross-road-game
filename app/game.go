package app

import (
	"github.com/gdamore/tcell/v2"
	"github.com/sirupsen/logrus"
	"math/rand"
	"strings"
	"time"
)

type Game struct {
	log     *logrus.Logger
	sheet   *Sheet
	display *Display
}

func drawText(s tcell.Screen, x1, y1, x2, y2 int, style tcell.Style, text string) {
	row := y1
	col := x1
	for _, r := range []rune(text) {
		s.SetContent(col, row, r, nil, style)
		col++
		if col >= x2 {
			s.SetContent(col-1, row, rune(' '), nil, style)
		}
		if row > y2 {
			break
		}
	}
}

func (g Game) Run() {
	//defStyle := tcell.StyleDefault.Background(tcell.ColorReset).Foreground(tcell.ColorReset)

	for i := 0; i < Rows; i++ {
		if i%8 == 0 {
			go g.createRestLine(i)
		} else {
			go g.createLine(i)
		}
	}

	for {
		//Log.Info(rowFirst[35])
		g.display.Screen.Show()
		time.Sleep(time.Microsecond * 1000)
	}
}

func (g Game) createLine(row int) {
	for {
		g.createVehicle(row, time.Duration(0), 0)
	}

	//duration := time.Microsecond * time.Duration(rand.Intn(10-3+1)+3) * 10000
	//for {
	//	vehicleLength := rand.Intn(25)
	//	waitDuration := duration * time.Duration(vehicleLength+rand.Intn(8-3+1)+3)
	//
	//	select {
	//	case <-time.After(waitDuration):
	//		go g.createVehicle(row, duration, vehicleLength)
	//	}
	//}
}

func (g Game) createVehicle(row int, duration time.Duration, vehicleLength int) {
	defStyle := tcell.StyleDefault.Background(tcell.ColorReset).Foreground(tcell.ColorReset)
	duration = time.Microsecond * time.Duration(rand.Intn(10)) * 10000
	vehicleLength = rand.Intn(25)

	c := 0
	for c < Cols+vehicleLength {
		l := vehicleLength
		x := c - l
		if c < l {
			l = c
			x = 0
		}
		g.sheet.cells[row][x+l] = true
		drawText(g.display.Screen, x, row, Cols, Rows, defStyle, strings.Repeat("=", l))
		if x != 0 {
			drawText(g.display.Screen, x-1, row, Cols, Rows, defStyle, " ")
			g.sheet.cells[row][x-1] = false
		}
		time.Sleep(duration)
		c++
	}
}

func (g Game) createRestLine(row int) {
	defStyle := tcell.StyleDefault.Background(tcell.ColorGray).Foreground(tcell.ColorReset)
	str := ""
	for i := 0; i < Cols; i++ {
		if rand.Intn(4) == 1 {
			g.sheet.cells[row][i] = true
			str += "#"
		} else {
			str += " "
		}
	}
	drawText(g.display.Screen, 0, row, Cols, Rows, defStyle, str)
	//drawText(g.display.Screen, 0, row, Cols, Rows, defStyle, strings.Repeat(" ", Cols))
}

func NewGame(log *logrus.Logger, sheet *Sheet, display *Display) *Game {
	return &Game{log: log, sheet: sheet, display: display}
}
