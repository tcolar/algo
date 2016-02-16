package main

import "fmt"

type HeapSort struct{}

func (s HeapSort) Name() string { return "Heap" }

// Heap sort an array (in place)
func (s HeapSort) Sort(data []int, debug bool) {
	h := Heap{data: data}
	n := len(data)
	// Heapify
	for k := n / 2; k >= 0; k-- {
		sink(h.data, k)
	}
	if debug {
		fmt.Printf("hp: heapified: %v\n", data)
	}
	// Sort, by putting max to right, then fixing the heap
	for n > 1 {
		data[0], data[n-1] = data[n-1], data[0]
		if debug {
			fmt.Printf("hp: s1 %v\n", data[:n-1])
		}
		sink(h.data[:n-1], 0)
		if debug {
			fmt.Printf("hp: s2 %v\n", data[:n-1])
		}
		n--
	}
	if debug {
		fmt.Printf("hp: sorted %v\n", data)
	}
}

// Heap built upon a slice
// left node is cur * 2
// right node is cur * 2 + 1
// parent node index is cur/2
type Heap struct {
	data    []int
	maxHeap bool // otherwise minHeap
}

func (h *Heap) Size() int {
	return len(h.data)
}

// Send down an item until it's in satisfies heap properties
func sink(d []int, k int) {
	for k*2+1 < len(d) {
		child := 2*k + 1 // 2*k+1 & 2*k+2 are the 2 children nodes
		if (child < len(d)-1) && (d[child] < d[child+1]) {
			child++ // which sibling
		}
		if d[k] < d[child] { // swap
			d[child], d[k] = d[k], d[child]
			k = child
		} else {
			return // in the right place
		}
	}
}
