// History: Dec 05 13 tcolar Creation

package algo

import (
	"log"
	"testing"
)

// Write algo such as if an element in N*M matrix is 0 the entire row or col
// are set to 0
func TestMatrixZero(t *testing.T) {
	// Not sure if this is the most efficient way ? best i could think of
	// ~ 0(n*m) complexity
	m := Matrix{}
	m.FillIn()
	m.Set(3, 2, 0)
	m.Set(9, 8, 0)
	// we now have 0 at 0,0; 3,2; 9,8;

	// Pass1 : find the zeroes and set bool flags for zeroed cols/rows
	zeroRows := [10]bool{}
	zeroCols := [10]bool{}
	cpt := 0
	for x := 0; x != 10; x++ { // rows
		for y := 0; y != 10; y++ { //cols
			if zeroCols[y] {
				continue // already zeroed
			}
			cpt++
			v := m.Get(x, y)
			if v == 0 {
				zeroRows[y] = true
				zeroCols[x] = true
				break // no point looking at rest of row
			}
		}
	}
	// Pass 2 zero rows, columns
	for y, b := range zeroRows {
		if b {
			for x := 0; x != 10; x++ {
				m.Set(x, y, 0)
				cpt++
			}
		}
	}
	for x, b := range zeroCols {
		if b {
			for y := 0; y != 10; y++ {
				if !zeroRows[y] { // already cleared "optiization"
					m.Set(x, y, 0)
					cpt++
				}
			}
		}
	}
	m.Print()
	// blanked cols
	check(m.Get(0, 9), 0)
	check(m.Get(3, 5), 0)
	check(m.Get(9, 7), 0)
	// blanked rows
	check(m.Get(2, 0), 0)
	check(m.Get(4, 2), 0)
	check(m.Get(9, 8), 0)
	// Should have been left alone
	check(m.Get(1, 9), 91)
	check(m.Get(8, 1), 18)
	check(m.Get(5, 5), 55)
	log.Printf("Did %d ops", cpt)
}
