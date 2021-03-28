package main

import (
	"math/rand"
	"time"
)

func GenerateDepthFirstMaze(grid *Grid) {
	visited := make(map[Coords]bool)
	for row := 0; row < grid.rows; row++ {
		for col := 0; col < grid.cols; col++ {
			visited[Coords{row, col }] = false
		}
	}

	currentCoords := Coords{0, 0 }
	visited[currentCoords] = true

	stack := make([]Coords, 0)
	stack = append(stack, currentCoords)

	for len(stack) != 0 {
		currentCoords = stack[len(stack) - 1]
		stack = stack[:len(stack) - 1]
		neighbours := GetUnvisitedNeighbours(&currentCoords, grid, &visited)

		for len(neighbours) != 0 {
			randomNeighbour := neighbours[rand.Intn(len(neighbours))]
			stack = append(stack, randomNeighbour)

			grid.cells[currentCoords.row][currentCoords.col] += GetDirection(currentCoords, randomNeighbour)
			grid.cells[randomNeighbour.row][randomNeighbour.col] += GetDirection(randomNeighbour, currentCoords)

			currentCoords = randomNeighbour
			visited[currentCoords] = true
			neighbours = GetUnvisitedNeighbours(&currentCoords, grid, &visited)
			time.Sleep(20 * time.Millisecond)
		}
	}
}

