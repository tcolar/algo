// History: Nov 25 13 tcolar Creation

package algo

import ()

// Heap built upon a slice
// left node is cur * 2
// right node is cur * 2 + 1
// parent node index is cur/2
type Heap struct {
	heap    []int
	maxHeap bool // otherwise minHeap
}

// Heapify the whole heap array
// (Enforce the heap properties)
// TODO: This currently does more than Heapify as it creates a sorted heap.
func (h *Heap) Heapify() {
	size := len(h.heap)
	start := (size - 2) / 2
	for start >= 0 { // all non-leaf nodes
		percolateDown(h, start, size)
		start -= 1
	}
}

func (h *Heap) Push(item int) {
	h.heap = append(h.heap, item)
	// Percolate up the new item by swapping with parent if larger than it.
	cur := len(h.heap) - 1
	parent := (cur - 1) / 2
	for cur >= 0 && h.heap[cur] > h.heap[parent] {
		h.heap[cur], h.heap[parent] = h.heap[parent], h.heap[cur]
		cur = parent
		parent = (cur - 1) / 2
	}
}

// Send down an item until it's in place in the heap
// Swap with sibling if necessary to have an "ordered" heap
func percolateDown(h *Heap, cur int, last int) {
	heap := h.heap
	for cur*2 < last {
		child := cur*2 + 1
		if (child < last) && (heap[child+1] > heap[child]) {
			child++
		}
		if heap[cur] < heap[child] {
			heap[child], heap[cur] = heap[cur], heap[child]
			cur = child
		} else {
			return
		}
	}
}
