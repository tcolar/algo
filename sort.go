// History: Nov 25 13 tcolar Creation

package algo

import (
	"log"
)

// Heap sort an array (in place)
func HeapSort(slice []int) {
	h := Heap{heap: slice}
	h.Heapify()
	last := len(slice) - 1
	for last >= 0 {
		slice[0], slice[last] = slice[last], slice[0]
		percolateDown(&h, 0, last-1)
		last--
	}
}

// Naive quick sort (not in place)
// Recursive
func Quicksort(slice []int) []int {
	if len(slice) <= 1 { // Can't be sorted any further
		return slice
	}
	pivot := slice[0] // default but can be lame, should find better partition
	low, high := []int{}, []int{}
	for i, v := range slice {
		if i == 0 {
			continue // pivot
		}
		if v <= pivot {
			low = append(low, v)
		} else {
			high = append(high, v)
		}
	}
	l := append(Quicksort(low), pivot)
	l = append(l, Quicksort(high)...)
	return l
}
