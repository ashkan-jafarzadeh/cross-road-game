package main

import (
	"github.com/gdamore/tcell/v2"
	"github.com/sirupsen/logrus"
	"log"
	"os"
	"time"
)

var rowFirst map[int]bool

func main() {
	logger()
	s := getScreen()

	rowFirst = make(map[int]bool)
	for i := 0; i < 100; i++ {
		rowFirst[i] = false
	}

	defer quit(s)

	go runGame(s)
	handleEvents(s)
}

func remove(slice []int, s int) []int {
	return append(slice[:s], slice[s+1:]...)
}

func runGame(s tcell.Screen) {
	defStyle := tcell.StyleDefault.Background(tcell.ColorReset).Foreground(tcell.ColorReset)
	s.SetSize(100, 100)

	go func() {
		t := 0
		for t < 100 {
			rowFirst[t] = true
			drawText(s, t, 1, 100, 100, defStyle, "====")
			if t != 0 {
				drawText(s, t-1, 1, 100, 100, defStyle, " ")
				rowFirst[t-1] = false
			}
			time.Sleep(time.Microsecond * 10000)
			t++
		}
	}()
	go func() {
		time.Sleep(time.Microsecond * 10000 * (2 + 9))
		t := 0
		for t < 100 {
			rowFirst[t] = true
			drawText(s, t, 1, 100, 100, defStyle, "##")
			if t != 0 {
				drawText(s, t-1, 1, 100, 100, defStyle, " ")
				rowFirst[t-1] = false
			}
			time.Sleep(time.Microsecond * 10000)
			t++
		}
	}()
	go func() {
		t := 0
		for t < 100 {
			t++
			drawText(s, t, 2, 100, 100, defStyle, "000000000000000000")
			drawText(s, t-1, 2, 100, 100, defStyle, " ")
			time.Sleep(time.Microsecond * 30000)
		}
	}()

	go func() {
		t := 0
		for t < 100 {
			t++
			drawText(s, t, 3, 100, 100, defStyle, "---")
			drawText(s, t-1, 3, 100, 100, defStyle, " ")
			time.Sleep(time.Microsecond * 15000)
		}
	}()

	for {
		//Log.Info(rowFirst[35])
		s.Show()
		time.Sleep(time.Microsecond * 1000)
	}
}

var Log = logrus.New()

// Init the actions.log file
func logger() error {
	file, err := os.OpenFile("log.txt", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err == nil {
		//Log.Formatter = &logrus.JSONFormatter{}
		Log.Out = file
	} else {
		Log.Info("Failed to log to file, using default stderr for now")
	}

	return nil
}

func getScreen() tcell.Screen {
	s, err := tcell.NewScreen()
	if err != nil {
		log.Fatalf("%+v", err)
	}
	if err := s.Init(); err != nil {
		log.Fatalf("%+v", err)
	}

	// Set default text style
	defStyle := tcell.StyleDefault.Background(tcell.ColorReset).Foreground(tcell.ColorReset)
	s.SetStyle(defStyle)

	// Clear screen
	s.Clear()

	return s
}

func quit(s tcell.Screen) {
	// You have to catch panics in a defer, clean up, and
	// re-raise them - otherwise your application can
	// die without leaving any diagnostic trace.
	maybePanic := recover()
	s.Fini()
	if maybePanic != nil {
		panic(maybePanic)
	}
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

func handleEvents(s tcell.Screen) {
	for {
		// Poll event
		ev := s.PollEvent()

		// Process event
		switch ev := ev.(type) {
		case *tcell.EventResize:
			s.Sync()
		case *tcell.EventKey:
			if ev.Key() == tcell.KeyEscape || ev.Key() == tcell.KeyCtrlC {
				return
			} else if ev.Key() == tcell.KeyCtrlL {
				s.Sync()
			} else if ev.Rune() == 'C' || ev.Rune() == 'c' {
				//s.Clear()
			}
		}
	}
}
