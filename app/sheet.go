package app

import (
	"github.com/sirupsen/logrus"
)

const Rows = 40
const Cols = 70

type col map[int]bool
type cells map[int]col

type Sheet struct {
	log   *logrus.Logger
	cells cells
}

func NewSheet(log *logrus.Logger) *Sheet {
	s := &Sheet{log: log}

	s.cells = make(cells)
	for i := 0; i < Rows; i++ {
		s.cells[i] = make(col)
		for j := 0; j < Cols; j++ {
			s.cells[i][j] = false
		}
	}

	return s
}