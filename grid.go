package main

type Grid struct {
	cols int
	rows int
	cells [][]int8
	cellWidth int
	cellHeight int
	resources []Texture
	resourceCount int8
}

func newGrid(rows, cols, cellWidth, cellHeight int) *Grid {
	grid := Grid {
		cols,
		rows,
		make([][]int8, 0),
		cellWidth,
		cellHeight,
		make([]Texture, 0),
		0,
	}

	for i := 0; i < rows; i++ {
		grid.cells = append(grid.cells, []int8 {})

		for j := 0; j < cols; j++ {
			grid.cells[i] = append(grid.cells[i], 0)
		}
	}

	return &grid
}

