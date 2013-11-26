// History: Nov 25 13 tcolar Creation

package algo

import ()

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
