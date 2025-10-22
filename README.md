# GAME OF LIFE

The Game of Life, also known as Conway's Game of Life or simply Life, is a cellular automaton devised by the British mathematician John Horton Conway in 1970.
This project is a Go implementation of Game of Life, a zero-player cellular automaton.
Each cell on a grid is either alive or dead, and its state in the next generation depends on its neighbors according to these rules:

Any live cell with fewer than two live neighbors dies (underpopulation).

Any live cell with two or three live neighbors lives on.

Any live cell with more than three live neighbors dies (overpopulation).

Any dead cell with exactly three live neighbors becomes alive (reproduction).

This Go implementation simulates the grid evolution step-by-step in the terminal.
