package app

import (
	"context"
	"errors"
	"fmt"
	"github.com/gdamore/tcell/v2"
	"time"
)

var (
	vehicleSpeedMax  = 15
	vehicleSpeedMin  = 2
	vehiclesDelayMax = 90
	vehiclesDelayMin = 40
)

type Game struct {
	sheet    *Sheet
	display  *Display
	player   *Player
	line     *Line
	lost     bool
	quitLine context.CancelFunc
	mode     int
}

func NewGame(sheet *Sheet, display *Display) *Game {
	return &Game{sheet: sheet, display: display, player: NewPlayer(display), line: NewLine(sheet, display)}
}

func (g *Game) Run() {
	g.player.show()
	g.createRoad()

	for {
		if g.hit() {
			break
		}
		g.display.Screen.Show()
		time.Sleep(time.Microsecond * 1000)
	}
	g.display.Screen.PostEvent(tcell.NewEventError(errors.New(fmt.Sprintf("Your Point: %d", g.player.point))))
}

func (g *Game) hit() bool {
	return g.sheet.cells[g.player.currentRow][g.player.currentCol]
}

func (g *Game) createRoad() {
	ctx, cancel := context.WithCancel(context.Background())
	g.quitLine = cancel
	for i := 0; i < Rows-1; i++ {
		if i%8 == 0 {
			go g.line.createRest(i)
		} else {
			go g.line.create(i, ctx)
		}
	}
}

func (g *Game) HandleEvents() {
	for {
		// Poll event
		ev := g.display.Screen.PollEvent()

		// Process event
		switch ev := ev.(type) {
		case *tcell.EventError:
			g.showPoint(ev.Error())
			//return
		case *tcell.EventResize:
			g.display.Screen.Sync()
		case *tcell.EventKey:
			if g.lost && ev.Key() == tcell.KeyEnter {
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
					g.refreshGame()
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

func (g *Game) showPoint(str string) {
	g.quitLine()
	g.display.Screen.Clear()
	style := tcell.StyleDefault.Background(tcell.ColorReset).Foreground(tcell.ColorRed)
	g.display.draw((Cols-10)/2, (Rows-1)/2, style, str)
	g.lost = true
	g.display.Screen.Sync()
}

func (g *Game) refreshGame() {
	g.quitLine()
	g.display.Screen.Clear()
	g.sheet.refresh()
	g.player.show()
	g.createRoad()
	//g.line.resetRests()
}
