// History: Nov 24 13 tcolar Creation

package algo

import (
	"strings"
)

// Exercise: Check whether the string is of the form "abcdeXedcba", using a stack
// Meaning a "Palindrome" separated with 'X' in the middle
func Palindromic(s string) bool {
	l := List{} // My test stack is implemented on a list
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
