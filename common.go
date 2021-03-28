package main

type Coords struct {
	row int
	col int
}

func GetUnvisitedNeighbours(coords *Coords, grid *Grid, visited *map[Coords]bool) []Coords {
	var result []Coords

	if coords.row - 1 >= 0 && !(*visited)[Coords{coords.row - 1, coords.col }] {
		result = append(result, Coords{coords.row - 1, coords.col })
	}

	if coords.row + 1 < grid.rows && !(*visited)[Coords{coords.row + 1, coords.col }] {
		result = append(result, Coords{coords.row + 1, coords.col })
	}

	if coords.col - 1 >= 0 && !(*visited)[Coords{coords.row, coords.col - 1 }] {
		result = append(result, Coords{coords.row, coords.col - 1 })
	}

	if coords.col + 1 < grid.cols && !(*visited)[Coords{coords.row, coords.col + 1 }] {
		result = append(result, Coords{coords.row, coords.col + 1 })
	}

	return result
}

func GetDirection(base, neighbour Coords) int8 {
	// Check North
	if base.row - neighbour.row < 0 {
		return 4
	}

	// Check South
	if base.row - neighbour.row > 0 {
		return 1
	}

	// Check West
	if base.col - neighbour.col > 0 {
		return 8
	}

	// Check East
	if base.col - neighbour.col < 0 {
		return 2
	}

	panic("Not known direction")
}

