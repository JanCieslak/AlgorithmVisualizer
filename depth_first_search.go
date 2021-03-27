package main

import (
	"math/rand"
	"time"
)

type Coords struct {
	row int
	col int
}

func getUnvisitedNeighbours(coords Coords, grid *Grid, visited *map[Coords]bool) []Coords {
	var result []Coords

	if coords.row - 1 >= 0 && !(*visited)[Coords { coords.row - 1, coords.col }] {
		result = append(result, Coords { coords.row - 1, coords.col })
	}

	if coords.row + 1 < grid.rows && !(*visited)[Coords { coords.row + 1, coords.col }] {
		result = append(result, Coords { coords.row + 1, coords.col })
	}

	if coords.col - 1 >= 0 && !(*visited)[Coords { coords.row, coords.col - 1 }] {
		result = append(result, Coords { coords.row, coords.col - 1 })
	}

	if coords.col + 1 < grid.cols && !(*visited)[Coords { coords.row, coords.col + 1 }] {
		result = append(result, Coords { coords.row, coords.col + 1 })
	}

	return result
}

func getDirection(base, neighbour Coords) int8 {
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

func generateMaze(grid *Grid) {
	visited := make(map[Coords]bool)
	for row := 0; row < grid.rows; row++ {
		for col := 0; col < grid.cols; col++ {
			visited[Coords{ row, col }] = false
		}
	}

	currentCoords := Coords {0, 0 }
	visited[currentCoords] = true

	stack := make([]Coords, 0)
	stack = append(stack, currentCoords)

	for len(stack) != 0 {
		currentCoords = stack[len(stack) - 1]
		stack = stack[:len(stack) - 1]
		neighbours := getUnvisitedNeighbours(currentCoords, grid, &visited)

		for len(neighbours) != 0 {
			randomNeighbour := neighbours[rand.Intn(len(neighbours))]
			stack = append(stack, randomNeighbour)

			grid.cells[currentCoords.row][currentCoords.col] += getDirection(currentCoords, randomNeighbour)
			grid.cells[randomNeighbour.row][randomNeighbour.col] += getDirection(randomNeighbour, currentCoords)

			currentCoords = randomNeighbour
			visited[currentCoords] = true
			neighbours = getUnvisitedNeighbours(currentCoords, grid, &visited)
			time.Sleep(20 * time.Millisecond)
		}
	}
}

