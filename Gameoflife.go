package main

import (
	"fmt"
	"strings"
)

type Grid struct {
	size uint
	data [][]bool
}

func newGrid(size uint, live ...int) Grid {

	grid := make([][]bool, size)

	for i := 0; i < int(size); i++ {
		grid[i] = make([]bool, size)
	}

	for i := 0; i < len(live); i = i + 2 {
		x := live[i]
		y := live[i+1]
		grid[x][y] = true
	}

	return Grid{
		size: size,
		data: grid,
	}
}

func displayGrid(grid Grid) string {
	var display strings.Builder

	for _, row := range grid.data {
		for _, val := range row {
			if val {
				display.WriteString("*")
			} else {
				display.WriteString("-")
			}
			display.WriteString(" ")
		}
		display.WriteString("\n")
	}
	return display.String()

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
	grid := newGrid(uint(size), 0, 9, 4, 4, 9, 0, 0, 0, 9, 9)

	fmt.Println(displayGrid(grid))

	// for {

	// 	fmt.Println(displayGrid(grid))
	// 	grid = applyRules(grid)
	// 	grid = newGeneration(grid)

	// }

}
