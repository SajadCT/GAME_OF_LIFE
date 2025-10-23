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
