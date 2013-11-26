// History: Nov 26 13 tcolar Creation

package algo

import (
	//  "log"
	"testing"
)

// Test various list aspects
func TestSearch(t *testing.T) {
	slice := []int{}
	check(BinarySearch(slice, 3), -1)
	slice = []int{1}
	check(BinarySearch(slice, 1), 0)
	slice = []int{1, 2, 3, 4, 5, 7, 8, 9, 10}
	check(BinarySearch(slice, 9), 7)
	check(BinarySearch(slice, 1), 0)
	check(BinarySearch(slice, 10), 8)
	check(BinarySearch(slice, 6), -1)
}
