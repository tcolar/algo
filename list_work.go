// History: Nov 24 13 tcolar Creation

package algo

import (
	"fmt"
	"log"
	"strings"
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

// Exercise: Check whether the list is of the form "abcdeXedcba", using a stack
// "Palindrome" separated with 'X' in the middle
func Palindromic(s string) bool {
	l := List{}
	if len(s) == 0 || !strings.Contains(s, "X") {
		return false
	}
	stacking := true
	for _, c := range s {
		// stacking until 'X'
		if stacking {
			if c != 'X' {
				l.Push(c)
			} else {
				stacking = false
			}
		} else {
			// Popping back in reverse order should match rest of string
			if l.Length == 0 || l.Pop() != c {
				return false
			}
		}
	}
	return l.Length == 0
}
