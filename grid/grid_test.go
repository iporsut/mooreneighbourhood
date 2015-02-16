package grid

import "testing"

func TestGetValueOfCellInGrid(t *testing.T) {
	g := Grid{
		{0, 1, 1},
		{0, 1, 0},
		{1, 1, 0},
	}
	v := g.Cell(1, 1)

	if v != 1 {
		t.Errorf("Expect v is 1 but was %d", v)
	}
}

func TestCountNeighbours(t *testing.T) {
	g := Grid{
		{0, 1, 1},
		{0, 1, 0},
		{1, 1, 0},
	}

	c := CountNeighbours(g, 1, 1)

	if c != 4 {
		t.Errorf("Expect count neighbours of (1,1) are 4 but was %d", c)
	}
}
