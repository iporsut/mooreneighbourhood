package grid

type Grid [][]int

func (g Grid) RowLen() int {
	return len(g)
}

func (g Grid) ColLen() int {
	return len(g[0])
}

func (g Grid) Cell(x, y int) int {
	if x < 0 || y < 0 || x >= g.RowLen() || y >= g.ColLen() {
		return 0
	}

	return g[x][y]
}

func CountNeighbours(g Grid, x, y int) int {
	northNeghbour := g.Cell(x-1, y)
	northEastNeghbour := g.Cell(x-1, y+1)
	eastNeghbour := g.Cell(x, y+1)
	southEastNeghbour := g.Cell(x+1, y+1)
	southNeghbour := g.Cell(x+1, y)
	southWestNeghbour := g.Cell(x+1, y-1)
	westNeghbour := g.Cell(x, y-1)
	northWestNeghbour := g.Cell(x-1, y-1)

	return northNeghbour + northEastNeghbour +
		eastNeghbour + southEastNeghbour +
		southNeghbour + southWestNeghbour +
		westNeghbour + northWestNeghbour
}
