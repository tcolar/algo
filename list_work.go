// History: Nov 24 13 tcolar Creation

package algo

import (
	"fmt"
	"log"
)

// Reverse the list "in place"
func (l *List) Reverse() {
	e := l.Root
	var prev *elem // nil to start
	for e != nil {
		next := e.Next // temp save "next"
		e.Next = prev
		prev = e
		e = next
	}
	l.Root = prev
}

// print the list for testing / debugging
func (l *List) print() {
	s := ""
	e := l.Root
	for e != nil {
		s += fmt.Sprintf("%d ,", e.Val) // assume klist of ints
		e = e.Next
	}
	log.Print(s)
}
