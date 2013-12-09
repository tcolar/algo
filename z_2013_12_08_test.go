// History: Dec 08 13 tcolar Creation

package algo

import (
	"testing"
)

// 2.2 Implement an algo to find the kth to last element of a singly linked list
func TestK2Last(t *testing.T) {
	l := List{}
	for i := 0; i != 25; i++ { // 0 .. 24
		l.Append(i)
	}
	check(K2Last(&l, 10), 14)
}

// Find it in o(n) - One pass
func K2Last(l *List, k int) int {
	e := l.Root
	ek := e // Pointer to e - k element
	cpt := 0
	for e.Next != nil {
		e = e.Next
		if cpt >= k {
			ek = ek.Next
		}
		cpt++
	}
	// TODO : if cpt < k return an error ?
	return ek.Val.(int)
}
