package app

import (
	"github.com/gdamore/tcell/v2"
)

type Display struct {
	Screen   tcell.Screen
	defStyle tcell.Style
}

func NewDisplay() (*Display, error) {
	d := &Display{}
	d.defStyle = tcell.StyleDefault.Background(tcell.ColorReset).Foreground(tcell.ColorReset)
	err := d.InitScreen()

	return d, err
}

func (d *Display) InitScreen() error {
	s, err := tcell.NewScreen()
	if err != nil {
		return err
	}
	if err := s.Init(); err != nil {
		return err
	}

	s.SetStyle(d.defStyle)
	s.Clear()

	Cols, Rows = s.Size()
	d.Screen = s

	return nil
}

func (d *Display) Quit() {
	maybePanic := recover()
	d.Screen.Fini()
	if maybePanic != nil {
		panic(maybePanic)
	}
}

func (d *Display) draw(x1, y1 int, style tcell.Style, text string) {
	x2 := Cols
	y2 := Rows
	row := y1
	col := x1
	for _, r := range []rune(text) {
		d.Screen.SetContent(col, row, r, nil, style)
		col++
		if col >= x2 {
			d.Screen.SetContent(col-1, row, rune(' '), nil, style)
		}
		if row > y2 {
			break
		}
	}
}
