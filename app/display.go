package app

import (
	"github.com/gdamore/tcell/v2"
)

type Display struct {
	Screen tcell.Screen
}

func NewDisplay() (*Display, error) {
	d := &Display{}
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

	// Set default text style
	defStyle := tcell.StyleDefault.Background(tcell.ColorReset).Foreground(tcell.ColorReset)
	s.SetStyle(defStyle)

	// Clear Screen
	s.Clear()

	//s.SetSize(Cols, Rows)

	d.Screen = s

	return nil
}

func (d *Display) Quit() {
	// You have to catch panics in a defer, clean up, and
	// re-raise them - otherwise your application can
	// die without leaving any diagnostic trace.
	maybePanic := recover()
	d.Screen.Fini()
	if maybePanic != nil {
		panic(maybePanic)
	}
}

func (d *Display) HandleEvents() {
	for {
		// Poll event
		ev := d.Screen.PollEvent()

		// Process event
		switch ev := ev.(type) {
		case *tcell.EventResize:
			d.Screen.Sync()
		case *tcell.EventKey:
			if ev.Key() == tcell.KeyEscape || ev.Key() == tcell.KeyCtrlC {
				return
			} else if ev.Key() == tcell.KeyCtrlL {
				d.Screen.Sync()
			} else if ev.Rune() == 'C' || ev.Rune() == 'c' {
				//s.Clear()
			}
		}
	}
}
