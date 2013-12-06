// History: Dec 05 13 tcolar Creation

package algo

import (
	"fmt"
	"log"
)

// Hard coded to 10x10 for testing
type Matrix struct {
	data [10][10]int
}

func (m *Matrix) Get(col int, row int) int {
	return m.data[col][row]
}

func (m *Matrix) Set(col int, row int, val int) {
	m.data[col][row] = val
}

// Fill in the matrix with vaues from 0 to 99
func (m *Matrix) FillIn() {
	for i := 0; i != 10; i++ {
		for j := 0; j != 10; j++ {
			m.data[j][i] = i*10 + j
		}
	}
}

func (m *Matrix) Print() {
	for i := 0; i != 10; i++ {
		str := ""
		for j := 0; j != 10; j++ {
			str += fmt.Sprintf(" %2d", m.data[j][i])
		}
		log.Print(str)
	}
}
