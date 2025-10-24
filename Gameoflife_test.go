package main

import (
	"reflect"
	"testing"
)

func TestNewGrid(t *testing.T) {
	size := 5
	actual := newGrid(uint(size))
	expected := Grid{
		size: 5,
		data: [][]bool{
			{false, false, false, false, false},
			{false, false, false, false, false},
			{false, false, false, false, false},
			{false, false, false, false, false},
			{false, false, false, false, false},
		},
	}
	if expected.size != actual.size {
		t.Errorf("EXPECTED :%v BUT GOT %d", expected.size, actual.size)

	}
	if !reflect.DeepEqual(actual, expected) {
		t.Errorf("EXPECTED :%v BUT GOT %v", expected.data, actual.data)
	}
}

func TestDisplayGrid(t *testing.T) {

	grid := Grid{
		size: 3,
		data: [][]bool{
			{true, false, false},
			{false, true, false},
			{false, false, true},
		},
	}
	actual := displayGrid(grid)
	expected := "* - - \n- * - \n- - * \n"

	if actual != expected {
		t.Errorf("Expected : \n%s ,but got : %s", expected, actual)
	}

}

func TestCountAliveNeighbourTopLeftCorner(t *testing.T) {
	grid := newGrid(5, 0, 1)
	actual := countAliveNeighbours(grid, 0, 0)

	if actual != 1 {
		t.Errorf("Expected neighbour count to be 1, but got %d", actual)
	}
	grid = newGrid(5, 0, 1, 1, 1)
	actual = countAliveNeighbours(grid, 0, 0)
	if actual != 2 {
		t.Errorf("Expected neighbour count to be 2, but got %d", actual)
	}

	grid = newGrid(5, 0, 1, 1, 1, 1, 0)
	actual = countAliveNeighbours(grid, 0, 0)
	if actual != 3 {
		t.Errorf("Expected neighbour count to be 3, but got %d", actual)
	}
}

func TestCountAliveNeighbourTopRightCorner(t *testing.T) {
	grid := newGrid(5, 0, 3)
	actual := countAliveNeighbours(grid, 0, 4)

	if actual != 1 {
		t.Errorf("Expected neighbour count to be 1, but got %d", actual)
	}
	grid = newGrid(5, 0, 3, 1, 3)
	actual = countAliveNeighbours(grid, 0, 4)
	if actual != 2 {
		t.Errorf("Expected neighbour count to be 2, but got %d", actual)
	}

	grid = newGrid(5, 0, 3, 1, 3, 1, 4)
	actual = countAliveNeighbours(grid, 0, 4)
	if actual != 3 {
		t.Errorf("Expected neighbour count to be 3, but got %d", actual)
	}
}

func TestCountAliveNeighbourBottomLeft(t *testing.T) {
	grid := newGrid(5, 3, 0)
	actual := countAliveNeighbours(grid, 4, 0)

	if actual != 1 {
		t.Errorf("Expected neighbour count to be 1, but got %d", actual)
	}
	grid = newGrid(5, 3, 0, 3, 1)
	actual = countAliveNeighbours(grid, 4, 0)
	if actual != 2 {
		t.Errorf("Expected neighbour count to be 2, but got %d", actual)
	}

	grid = newGrid(5, 3, 0, 3, 1, 4, 1)
	actual = countAliveNeighbours(grid, 4, 0)
	if actual != 3 {
		t.Errorf("Expected neighbour count to be 3, but got %d", actual)
	}
}

func TestCountAliveNeighbourBottomRight(t *testing.T) {
	grid := newGrid(5, 4, 3)
	actual := countAliveNeighbours(grid, 4, 4)

	if actual != 1 {
		t.Errorf("Expected neighbour count to be 1, but got %d", actual)
	}
	grid = newGrid(5, 4, 3, 3, 3)
	actual = countAliveNeighbours(grid, 4, 4)
	if actual != 2 {
		t.Errorf("Expected neighbour count to be 2, but got %d", actual)
	}

	grid = newGrid(5, 4, 3, 3, 3, 3, 4)
	actual = countAliveNeighbours(grid, 4, 4)
	if actual != 3 {
		t.Errorf("Expected neighbour count to be 3, but got %d", actual)
	}
}

func TestCountAliveNeighboursTopEdge(t *testing.T) {
	grid := newGrid(5, 0, 1)
	actual := countAliveNeighbours(grid, 0, 2)

	if actual != 1 {
		t.Errorf("Expected neighbour count to be 1, but got %d", actual)
	}
	grid = newGrid(5, 0, 1, 1, 1)
	actual = countAliveNeighbours(grid, 0, 2)
	if actual != 2 {
		t.Errorf("Expected neighbour count to be 2, but got %d", actual)
	}

	grid = newGrid(5, 0, 1, 1, 1, 1, 2)
	actual = countAliveNeighbours(grid, 0, 2)
	if actual != 3 {
		t.Errorf("Expected neighbour count to be 3, but got %d", actual)
	}

	grid = newGrid(5, 0, 1, 1, 1, 1, 2, 1, 3)
	actual = countAliveNeighbours(grid, 0, 2)
	if actual != 4 {
		t.Errorf("Expected neighbour count to be 3, but got %d", actual)
	}

	grid = newGrid(5, 0, 1, 1, 1, 1, 2, 1, 3, 0, 3)
	actual = countAliveNeighbours(grid, 0, 2)
	if actual != 5 {
		t.Errorf("Expected neighbour count to be 3, but got %d", actual)
	}
}

func TestCountAliveNeighboursBottomEdge(t *testing.T) {
	grid := newGrid(5, 4, 1)
	actual := countAliveNeighbours(grid, 4, 2)

	if actual != 1 {
		t.Errorf("Expected neighbour count to be 1, but got %d", actual)
	}
	grid = newGrid(5, 4, 1, 3, 1)
	actual = countAliveNeighbours(grid, 4, 2)
	if actual != 2 {
		t.Errorf("Expected neighbour count to be 2, but got %d", actual)
	}

	grid = newGrid(5, 4, 1, 3, 1, 3, 2)
	actual = countAliveNeighbours(grid, 4, 2)
	if actual != 3 {
		t.Errorf("Expected neighbour count to be 3, but got %d", actual)
	}

	grid = newGrid(5, 4, 1, 3, 1, 3, 2, 3, 3)
	actual = countAliveNeighbours(grid, 4, 2)
	if actual != 4 {
		t.Errorf("Expected neighbour count to be 4, but got %d", actual)
	}

	grid = newGrid(5, 4, 1, 3, 1, 3, 2, 3, 3, 4, 3)
	actual = countAliveNeighbours(grid, 4, 2)
	if actual != 5 {
		t.Errorf("Expected neighbour count to be 5, but got %d", actual)
	}
}

func TestCountAliveNeighboursLeftEdge(t *testing.T) {
	grid := newGrid(5, 1, 0)
	actual := countAliveNeighbours(grid, 2, 0)

	if actual != 1 {
		t.Errorf("Expected neighbour count to be 1, but got %d", actual)
	}
	grid = newGrid(5, 1, 0, 1, 1)
	actual = countAliveNeighbours(grid, 2, 0)
	if actual != 2 {
		t.Errorf("Expected neighbour count to be 2, but got %d", actual)
	}

	grid = newGrid(5, 1, 0, 1, 1, 2, 1)
	actual = countAliveNeighbours(grid, 2, 0)
	if actual != 3 {
		t.Errorf("Expected neighbour count to be 3, but got %d", actual)
	}

	grid = newGrid(5, 1, 0, 1, 1, 2, 1, 3, 1)
	actual = countAliveNeighbours(grid, 2, 0)
	if actual != 4 {
		t.Errorf("Expected neighbour count to be 4, but got %d", actual)
	}

	grid = newGrid(5, 1, 0, 1, 1, 2, 1, 3, 1, 3, 0)
	actual = countAliveNeighbours(grid, 2, 0)
	if actual != 5 {
		t.Errorf("Expected neighbour count to be 5, but got %d", actual)
	}
}
func TestCountAliveNeighboursRightEdge(t *testing.T) {
	grid := newGrid(5, 1, 4)
	actual := countAliveNeighbours(grid, 2, 4)

	if actual != 1 {
		t.Errorf("Expected neighbour count to be 1, but got %d", actual)
	}
	grid = newGrid(5, 1, 4, 1, 3)
	actual = countAliveNeighbours(grid, 2, 4)
	if actual != 2 {
		t.Errorf("Expected neighbour count to be 2, but got %d", actual)
	}

	grid = newGrid(5, 1, 4, 1, 3, 2, 3)
	actual = countAliveNeighbours(grid, 2, 4)
	if actual != 3 {
		t.Errorf("Expected neighbour count to be 3, but got %d", actual)
	}

	grid = newGrid(5, 1, 4, 1, 3, 2, 3, 3, 3)
	actual = countAliveNeighbours(grid, 2, 4)
	if actual != 4 {
		t.Errorf("Expected neighbour count to be 4, but got %d", actual)
	}

	grid = newGrid(5, 1, 4, 1, 3, 2, 3, 3, 3, 3, 4)
	actual = countAliveNeighbours(grid, 2, 4)
	if actual != 5 {
		t.Errorf("Expected neighbour count to be 5, but got %d", actual)
	}
}

func TestCountAliveNeighboursMiddle(t *testing.T) {
	grid := newGrid(5, 1, 2)
	actual := countAliveNeighbours(grid, 2, 2)

	if actual != 1 {
		t.Errorf("Expected neighbour count to be 1, but got %d", actual)
	}
	grid = newGrid(5, 1, 2, 1, 3)
	actual = countAliveNeighbours(grid, 2, 2)
	if actual != 2 {
		t.Errorf("Expected neighbour count to be 2, but got %d", actual)
	}

	grid = newGrid(5, 1, 2, 1, 3, 2, 3)
	actual = countAliveNeighbours(grid, 2, 2)
	if actual != 3 {
		t.Errorf("Expected neighbour count to be 3, but got %d", actual)
	}

	grid = newGrid(5, 1, 2, 1, 3, 2, 3, 3, 3)
	actual = countAliveNeighbours(grid, 2, 2)
	if actual != 4 {
		t.Errorf("Expected neighbour count to be 4, but got %d", actual)
	}

	grid = newGrid(5, 1, 2, 1, 3, 2, 3, 3, 3, 3, 2)
	actual = countAliveNeighbours(grid, 2, 2)
	if actual != 5 {
		t.Errorf("Expected neighbour count to be 5, but got %d", actual)
	}

	grid = newGrid(5, 1, 2, 1, 3, 2, 3, 3, 3, 3, 2, 3, 1)
	actual = countAliveNeighbours(grid, 2, 2)
	if actual != 6 {
		t.Errorf("Expected neighbour count to be 6, but got %d", actual)
	}

	grid = newGrid(5, 1, 2, 1, 3, 2, 3, 3, 3, 3, 2, 3, 1, 2, 1)
	actual = countAliveNeighbours(grid, 2, 2)
	if actual != 7 {
		t.Errorf("Expected neighbour count to be 7, but got %d", actual)
	}

	grid = newGrid(5, 1, 2, 1, 3, 2, 3, 3, 3, 3, 2, 3, 1, 2, 1, 1, 1)
	actual = countAliveNeighbours(grid, 2, 2)
	if actual != 8 {
		t.Errorf("Expected neighbour count to be 8, but got %d", actual)
	}
}

func TestRule1(t *testing.T) { //underpopulation
	grid := newGrid(3, 0, 0)
	nextGen := runGeneration(grid)

	if nextGen.data[0][0] {
		t.Errorf("Expected cell[0,0] to be dead due to under population, but it is alive")
	}

}
func TestRule2(t *testing.T) { //survival
	grid := newGrid(4, 1, 1, 1, 2, 1, 3)
	nextGen := runGeneration(grid)

	if !nextGen.data[1][2] {
		t.Errorf("Expected cell[1,2] to survive, but its dead")
	}

}
