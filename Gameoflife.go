package main

import (
	"fmt"
)

type Grid struct {
	size uint
	data [][]bool
}

func newGrid(size uint) Grid {

	grid := make([][]bool, size)

	for i := 0; i < int(size); i++ {
		grid[i] = make([]bool, size)
	}

	return Grid{
		size: size,
		data: grid,
	}

}

func displayGrid(grid Grid) string {
	return ""
}

func newGeneration(grid Grid) Grid {
	return grid
}

func countNeighbours(grid Grid, x uint, y uint) uint {
	return 0
}

func countAliveNeighbours(grid Grid, x uint, y uint) uint {
	return 0
}

func applyRules(grid Grid) Grid {
	return grid
}

func main() {
	size := 10
	grid := newGrid(uint(size))
	//fmt.Println(grid)

	for {

		fmt.Println(displayGrid(grid))
		grid = applyRules(grid)
		grid = newGeneration(grid)

	}

}
