package main

import (
	"math/rand"
	"time"
)

type Wall struct {
	from Coords
	to Coords
}

func getUnvisitedWalls(coords Coords, visited *map[Coords]bool, grid *Grid) []Wall {
	var result []Wall

	if coords.row - 1 >= 0 && !(*visited)[Coords{coords.row - 1, coords.col }] {
		result = append(result, Wall { coords, Coords{coords.row - 1, coords.col }})
	}

	if coords.row + 1 < grid.rows && !(*visited)[Coords{coords.row + 1, coords.col }] {
		result = append(result, Wall { coords, Coords{coords.row + 1, coords.col }})
	}

	if coords.col - 1 >= 0 && !(*visited)[Coords{coords.row, coords.col - 1 }] {
		result = append(result, Wall { coords, Coords{coords.row, coords.col - 1 }})
	}

	if coords.col + 1 < grid.cols && !(*visited)[Coords{coords.row, coords.col + 1 }] {
		result = append(result, Wall { coords, Coords{coords.row, coords.col + 1 }})
	}

	return result
}

func GeneratePrimsMaze(grid *Grid) {
	visited := make(map[Coords]bool)
	for row := 0; row < grid.rows; row++ {
		for col := 0; col < grid.cols; col++ {
			visited[Coords{row, col }] = false
		}
	}

	currentCoords := Coords{0, 0 }
	visited[currentCoords] = true
	walls := make([]Wall, 0)
	walls = append(walls, getUnvisitedWalls(currentCoords, &visited, grid)...)

	for len(walls) > 0 {
		randomIndex := rand.Intn(len(walls))
		randomWall := walls[randomIndex]

		if visited[randomWall.from] && !visited[randomWall.to] || !visited[randomWall.from] && visited[randomWall.to] {
			grid.cells[randomWall.from.row][randomWall.from.col] += GetDirection(randomWall.from, randomWall.to)
			grid.cells[randomWall.to.row][randomWall.to.col] += GetDirection(randomWall.to, randomWall.from)

			if visited[randomWall.from] {
				visited[randomWall.to] = true
			} else {
				visited[randomWall.from] = true
			}

			walls = append(walls[:randomIndex], walls[randomIndex + 1:]...)
			walls = append(walls, getUnvisitedWalls(randomWall.from, &visited, grid)...)
			walls = append(walls, getUnvisitedWalls(randomWall.to, &visited, grid)...)
			time.Sleep(50 * time.Millisecond)
		}
	}
}