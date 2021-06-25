/*
“ Given a matrix of 1s and 0s, return the number of “islands” in the matrix.
A 1 represents land and 0 represents water, so an island is a group of 1s that are neighboring whose perimeter is surrounded by water.
For example, this matrix has 4 islands. ”
1 0 0 0 0
0 0 1 1 0
0 1 1 0 0
0 0 0 0 0
1 1 0 0 1
1 1 0 0 1
*/

package main

import "fmt"

func main() {
	b := board{
		{1, 0, 0, 0, 0},
		{0, 0, 1, 1, 0},
		{0, 1, 1, 0, 0},
		{0, 0, 0, 0, 0},
		{1, 1, 0, 0, 1},
		{1, 1, 0, 0, 1},
	}
	islands := 0
	for rowIdx, row := range b {
		for colIdx := range row {
			if b[rowIdx][colIdx] == 1 {
				islands++
				findIsland(b, rowIdx, colIdx, islands)
			}
		}
	}
	fmt.Println(b)
	fmt.Println(islands)
}

func findIsland(b board, row, col, count int) {
	if row < 0 || col < 0 || row >= len(b) || col >= len(b[row]) {
		return
	}

	// not land
	if b[row][col] != 1 {
		return
	}

	// land
	b[row][col] = count + 1 // mark the current location as seen
	for i := row - 1; i <= row+1; i++ {
		for j := col - 1; j <= col+1; j++ {
			if i != row || j != col {
				findIsland(b, i, j, count)
			}
		}
	}
}

type board [][]int
