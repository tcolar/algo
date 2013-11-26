// History: Nov 26 13 tcolar Creation

package algo

import (
//"log"
)

// Search for target in (sorted) slice and return index where found
// or -1 if not found
func BinarySearch(slice []int, target int) (pos int) {
	start := 0
	end := len(slice) - 1
	for start <= end {
		mid := start + (end-start)/2
		if slice[mid] == target {
			return mid // found
		}
		if target > slice[mid] {
			start = mid + 1
		} else {
			end = mid - 1
		}
	}
	// Not found
	return -1
}
