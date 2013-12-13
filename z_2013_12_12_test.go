// History: Dec 13 13 tcolar Creation

package algo

import (
	"log"
	"testing"
)

type Pos struct {
	Row int
	Col int
}

func TestNQueens(t *testing.T) {
	queens := []Pos{}
	n := 9
	result := NQueens(n, 0, queens)
	PrintQueens(n, result)
}

// The N queen problem is that we have to place N queens on a N*N chess baord such
// as they don't attack one  other (This is guaranteed to works for N>=4)
// Goal of exercise it to find the first solution for a given N
// I Figured out that for a n*n matrix with n queen their has to be exactly one queen
// per row and per column (because of queen property of "capturing" whole row and col)
func NQueens(n int, row int, queens []Pos) []Pos {
	found := false
	if n < 4 {
		log.Fatal("Can't be solved for n<4")
	}
	// For any given row we must find a place to place a queen (since one per row/col)
	// We need to try placing the first Queen at different location in the row
	// until it works out
	for colOffset := 0; colOffset != n; colOffset++ {
		for col := colOffset; col != n; col++ {
			if IsOpen(queens, row, col) {
				// found an open position for our queen, add it
				q := Pos{
					Row: row,
					Col: col,
				}
				queens = append(queens, q)
				//log.Printf("Added queen %d at row:%d, col:%d", len(*queens), row, col)
				if len(queens) == n {
					// Found a solution, we are done
					found = true
					return queens
				} else if row+1 < n {
					//Recurse to the next row only if we found a queen in the current row
					// otherwise we already failed
					queens = NQueens(n, row+1, queens)
					// if solution found during recusrion we are done
					if found {
						return queens
					}
				}
				break // since we found this row's queen, we can reak out of this loop
			}
		}
		// If we failed to find enough queens we will try starting from a differrent
		// position(colOffset) in the base row
		// But first we need to clear the "failed" queens at or bellow the current row
		if len(queens) != n {
			for i, q := range queens {
				if q.Row >= row {
					queens = append(queens[:i], queens[i+1:]...)
				}
			}
		}
	}
	return queens
}

// Check whether a position is safe from all queens so far
func IsOpen(queens []Pos, row int, col int) bool {
	for _, q := range queens {
		// If col already owned by another queen
		if q.Col == col {
			return false
		}
		// check if "diagonal" is owned by another queen
		if abs(q.Row-row) == abs(q.Col-col) {
			return false
		}
	}
	return true
}

func abs(nb int) int {
	if nb >= 0 {
		return nb
	} else {
		return -nb
	}
}

// Utility method to pretty print the result as an ASCII chessboard
func PrintQueens(n int, queens []Pos) {
	log.Println()
	board := [15][15]bool{} // Hardoded size to test
	for _, q := range queens {
		board[q.Row][q.Col] = true
	}
	for r := 0; r != n; r++ {
		ln := "|"
		for c := 0; c != n; c++ {
			if board[r][c] {
				ln += "Q|"
			} else {
				ln += " |"
			}
		}
		log.Println(ln)
	}
}
