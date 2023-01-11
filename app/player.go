package app

import (
	"fmt"
)

type Player struct {
	display    *Display
	char       string
	currentCol int
	currentRow int
	point      int
}

func NewPlayer(display *Display) *Player {
	return &Player{
		display: display,
		//char:       "🕴",
		char:       "🐐",
		currentRow: Rows - 1,
		currentCol: (Cols - 2) / 2}
}

func (p *Player) show() {
	p.display.draw(p.currentCol, p.currentRow, p.display.defStyle, p.char)
	p.showPoint()
}

func (p *Player) moveUp() bool {
	if p.currentRow == 0 {
		p.reset()
		p.point += 1
		p.showPoint()
		return true
	}
	p.move(0, -1)
	return false
}

func (p *Player) moveLeft() {
	if p.currentCol == 0 {
		return
	}
	p.move(-1, 0)
}

func (p *Player) moveDown() {
	if p.currentRow == Rows-1 {
		return
	}
	p.move(0, 1)
}

func (p *Player) moveRight() {
	if p.currentCol == Cols-2 {
		return
	}
	p.move(1, 0)
}

func (p *Player) move(col int, row int) {
	p.display.draw(p.currentCol, p.currentRow, p.display.defStyle, " ")
	p.currentCol = p.currentCol + col
	p.currentRow = p.currentRow + row
	p.display.draw(p.currentCol, p.currentRow, p.display.defStyle, p.char)
}

func (p *Player) reset() {
	p.display.draw(p.currentCol, p.currentRow, p.display.defStyle, " ")
	p.currentRow = Rows - 1
	p.currentCol = (Cols - 2) / 2
	p.display.draw(p.currentCol, p.currentRow, p.display.defStyle, p.char)
}

func (p *Player) showPoint() {
	p.display.draw(0, Rows-1, p.display.defStyle, fmt.Sprintf("%d", p.point))
}
