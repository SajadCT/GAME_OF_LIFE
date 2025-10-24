package main

import (
	"fmt"
	"strings"
	"time"
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

func runGeneration(grid Grid) Grid {
	newGen := newGrid(grid.size)
	for i := 0; i < int(grid.size); i++ {
		for j := 0; j < int(grid.size); j++ {
			cell := grid.data[i][j]
			count := countAliveNeighbours(grid, uint(i), uint(j))

			if cell && count < 2 { //underpopulation
				newGen.data[i][j] = false
			}
			if cell && (count == 2 || count == 3) { //survival
				newGen.data[i][j] = true
			}
			if cell && count > 3 { //over population
				newGen.data[i][j] = false
			}
			if !cell && count == 3 {
				newGen.data[i][j] = true
			}
		}
	}

	return newGen
}

func b2i(b bool) uint {
	if b {
		return 1
	}
	return 0
}

func countAliveNeighbours(grid Grid, x uint, y uint) uint {
	s := grid.size - 1
	if grid.size == 1 {
		return uint(0)
	}
	if x == 0 && y == 0 { // top left corner
		return b2i(grid.data[0][1]) + b2i(grid.data[1][1]) + b2i(grid.data[1][0])
	}
	if x == 0 && y == s { //top right corner
		return b2i(grid.data[x][y-1]) + b2i(grid.data[x+1][y-1]) + b2i(grid.data[x+1][y])
	}
	if x == s && y == 0 { //bottom left corner
		return b2i(grid.data[x-1][y]) + b2i(grid.data[x-1][y+1]) + b2i(grid.data[x][y+1])
	}
	if x == s && y == s { //bottom right corner
		return b2i(grid.data[x-1][y]) + b2i(grid.data[x-1][y-1]) + b2i(grid.data[x][y-1])
	}
	if x == 0 { //top edge
		return b2i(grid.data[x][y-1]) + b2i(grid.data[x][y+1]) + b2i(grid.data[x+1][y-1]) + b2i(grid.data[x+1][y]) + b2i(grid.data[x+1][y+1])
	}
	if x == s { //bottom edge
		return b2i(grid.data[x][y-1]) + b2i(grid.data[x][y+1]) + b2i(grid.data[x-1][y-1]) + b2i(grid.data[x-1][y]) + b2i(grid.data[x-1][y+1])
	}
	if y == 0 { //left edge
		return b2i(grid.data[x-1][y]) + b2i(grid.data[x+1][y]) + b2i(grid.data[x-1][y+1]) + b2i(grid.data[x][y+1]) + b2i(grid.data[x+1][y+1])
	}
	if y == s { //right edge
		return b2i(grid.data[x-1][y]) + b2i(grid.data[x+1][y]) + b2i(grid.data[x-1][y-1]) + b2i(grid.data[x][y-1]) + b2i(grid.data[x+1][y-1])
	}
	if x > 0 && x < s && y > 0 && y < s { //middle
		return b2i(grid.data[x-1][y]) + b2i(grid.data[x+1][y]) + b2i(grid.data[x][y+1]) + b2i(grid.data[x][y-1]) + b2i(grid.data[x-1][y+1]) + b2i(grid.data[x+1][y+1]) + b2i(grid.data[x-1][y-1]) + b2i(grid.data[x+1][y-1])

	}

	return 0
}

func applyRules(grid Grid) Grid {
	return grid
}

func main() {
	size := 30
	grid := newGrid(uint(size), 3, 4, 4, 4, 5, 4)

	for {
		fmt.Print("\033[H\033[2J\033[3J")
		fmt.Print(displayGrid(grid))
		time.Sleep(1 * time.Second)
		grid = runGeneration(grid)
	}

}
