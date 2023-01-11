package app

var Rows = 20
var Cols = 60

type cells map[int]map[int]bool

type Sheet struct {
	cells cells
}

func NewSheet() *Sheet {
	s := &Sheet{}
	s.createCells()
	return s
}

func (s *Sheet) createCells() {
	s.cells = make(cells)
	for i := 0; i < Rows; i++ {
		s.cells[i] = make(map[int]bool)
		for j := 0; j < Cols; j++ {
			s.cells[i][j] = false
		}
	}
}

func (s *Sheet) refresh() {
	for i := 0; i < Rows; i++ {
		for j := 0; j < Cols; j++ {
			s.cells[i][j] = false
		}
	}
}
