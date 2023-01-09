package app

import (
	"errors"
	"fmt"
	"github.com/gdamore/tcell/v2"
	"github.com/sirupsen/logrus"
	"time"
)

type Game struct {
	log     *logrus.Logger
	sheet   *Sheet
	display *Display
	player  *Player
	line    *Line
	lost    bool
}

func (g *Game) Run() {
	g.player.show()
	for i := 0; i < Rows-1; i++ {
		if i%8 == 0 {
			go g.line.createRest(i)
		} else {
			go g.line.create(i)
		}
	}

	for {
		if g.hit() {
			break
		}
		g.display.Screen.Show()
		time.Sleep(time.Microsecond * 1000)
	}
	g.display.Screen.PostEvent(tcell.NewEventError(errors.New(fmt.Sprintf("Your Point: %d", g.player.point))))
}

func (g *Game) HandleEvents() {
	for {
		// Poll event
		ev := g.display.Screen.PollEvent()

		// Process event
		switch ev := ev.(type) {
		case *tcell.EventError:
			g.display.Screen.Clear()
			g.display.draw((Cols-2)/2, (Rows-1)/2, g.display.defStyle, ev.Error())
			g.lost = true
			g.display.Screen.Sync()
			//return
		case *tcell.EventResize:
			g.display.Screen.Sync()
		case *tcell.EventKey:
			if g.lost {
				return
				//InitGame()
			}
			if ev.Key() == tcell.KeyEscape || ev.Key() == tcell.KeyCtrlC {
				return
			}

			switch ev.Key() {
			case tcell.KeyUp:
				finished := g.player.moveUp()
				if finished {
					g.line.resetRests()
				}
			case tcell.KeyDown:
				g.player.moveDown()
			case tcell.KeyRight:
				g.player.moveRight()
			case tcell.KeyLeft:
				g.player.moveLeft()
			}
		}
	}
}

func (g *Game) hit() bool {
	return g.sheet.cells[g.player.currentRow][g.player.currentCol]
}

func NewGame(log *logrus.Logger, sheet *Sheet, display *Display) *Game {
	return &Game{log: log, sheet: sheet, display: display, player: NewPlayer(display), line: NewLine(sheet, display)}
}
