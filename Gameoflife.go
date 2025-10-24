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

	return b2i(grid.data[x-1][y]) + b2i(grid.data[x+1][y]) + b2i(grid.data[x][y+1]) + b2i(grid.data[x][y-1]) + b2i(grid.data[x-1][y+1]) + b2i(grid.data[x+1][y+1]) + b2i(grid.data[x-1][y-1]) + b2i(grid.data[x+1][y-1])
}

func main() {
	size := 40
	grid := newGrid(uint(size),
		12, 10, 13, 10, 14, 10, 18, 10, 19, 10, 20, 10,
		10, 12, 15, 12, 17, 12, 22, 12,
		10, 13, 15, 13, 17, 13, 22, 13,
		10, 14, 15, 14, 17, 14, 22, 14,
		12, 15, 13, 15, 14, 15, 18, 15, 19, 15, 20, 15,
		12, 17, 13, 17, 14, 17, 18, 17, 19, 17, 20, 17,
		10, 18, 15, 18, 17, 18, 22, 18,
		10, 19, 15, 19, 17, 19, 22, 19,
		10, 20, 15, 20, 17, 20, 22, 20,
		12, 22, 13, 22, 14, 22, 18, 22, 19, 22, 20, 22,
	)

	for {
		fmt.Print("\033[H\033[2J\033[3J")
		fmt.Print(displayGrid(grid))
		time.Sleep(1 * time.Second)
		grid = runGeneration(grid)
	}

}
