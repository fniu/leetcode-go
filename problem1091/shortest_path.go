package problem1091

type Cell struct {
	x, y int
}

func shortestPathBinaryMatrix(grid [][]int) int {

	l := len(grid)

	if grid[0][0] == 1 || grid[l-1][l-1] == 1 {
		return -1
	}

	if l == 1 && grid[0][0] == 0 {
		return 1
	}

	nextSteps := [8][2]int{{1, 1}, {1, 0}, {0, 1}, {-1, 0}, {0, -1}, {1, -1}, {-1, 1}, {-1, -1}}

	qf := make([]Cell, 0)
	cf := Cell{0, 0}
	qf = append(qf, cf)
	steps := 1
	var x1, y1 int
	var head Cell
	grid[0][0] = 1
	for {
		newQ := make([]Cell, 0)
		for len(qf) > 0 {
			head = qf[0]
			qf = qf[1:]
			for i := 0; i < 8; i++ {
				x1 = head.x + nextSteps[i][0]
				y1 = head.y + nextSteps[i][1]
				if x1 == l-1 && y1 == l-1 {
					return steps + 1
				}
				if x1 >= 0 && x1 < l && y1 >= 0 && y1 < l && grid[x1][y1] == 0 {
					newQ = append(newQ, Cell{x1, y1})
					grid[x1][y1] = 1
				}
			}
		}
		if len(newQ) == 0 {
			break
		}
		qf = newQ
		steps++
	}
	return -1
}

func shortestPathBinaryMatrixHash(grid [][]int) int {

	l := len(grid)

	if grid[0][0] == 1 || grid[l-1][l-1] == 1 {
		return -1
	}

	if l == 1 && grid[0][0] == 0 {
		return 1
	}

	nextSteps := [8][2]int{{1, 1}, {1, 0}, {0, 1}, {-1, 0}, {0, -1}, {1, -1}, {-1, 1}, {-1, -1}}

	q := make(map[Cell]bool)
	q[Cell{0, 0}] = true

	steps := 1
	var x1, y1 int
	grid[0][0] = 1
	grid[l-1][l-1] = 1
	for {
		newQ := make(map[Cell]bool)
		for k := range q {
			for i := 0; i < 8; i++ {
				x1 = k.x + nextSteps[i][0]
				y1 = k.y + nextSteps[i][1]
				if x1 == l-1 && y1 == l-1 {
					return steps + 1
				}
				if x1 >= 0 && x1 < l && y1 >= 0 && y1 < l && grid[x1][y1] == 0 {
					newQ[Cell{x1, y1}] = true
					grid[x1][y1] = 1
				}
			}
		}
		if len(newQ) == 0 {
			break
		}
		q = newQ
		steps++
	}
	return -1
}

func shortestPathBinaryMatrixBothWays(grid [][]int) int {

	l := len(grid)

	if grid[0][0] == 1 || grid[l-1][l-1] == 1 {
		return -1
	}

	if l == 1 && grid[0][0] == 0 {
		return 1
	}

	nextSteps := [8][2]int{{1, 1}, {1, 0}, {0, 1}, {-1, 0}, {0, -1}, {1, -1}, {-1, 1}, {-1, -1}}

	qf := make(map[Cell]bool)
	qf[Cell{0, 0}] = true

	qb := make(map[Cell]bool)
	qb[Cell{l - 1, l - 1}] = true

	steps := 1
	var x1, y1 int
	var c Cell
	grid[0][0] = 1
	grid[l-1][l-1] = 1
	for {
		newQ := make(map[Cell]bool)
		for k := range qf {
			for i := 0; i < 8; i++ {
				x1 = k.x + nextSteps[i][0]
				y1 = k.y + nextSteps[i][1]
				c = Cell{x1, y1}
				if _, ok := qb[c]; ok {
					return 2 * steps
				}
				if x1 >= 0 && x1 < l && y1 >= 0 && y1 < l && grid[x1][y1] == 0 {
					newQ[c] = true
					grid[x1][y1] = 1
				}
			}
		}
		if len(newQ) == 0 {
			break
		}
		qf = newQ

		newQ = make(map[Cell]bool)
		for k := range qb {
			for i := 0; i < 8; i++ {
				x1 = k.x + nextSteps[i][0]
				y1 = k.y + nextSteps[i][1]
				c = Cell{x1, y1}
				if _, ok := qf[c]; ok {
					return 2*steps + 1
				}
				if x1 >= 0 && x1 < l && y1 >= 0 && y1 < l && grid[x1][y1] == 0 {
					newQ[c] = true
					grid[x1][y1] = 1
				}
			}
		}
		if len(newQ) == 0 {
			break
		}
		qb = newQ

		steps++
	}
	return -1
}
