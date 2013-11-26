// History: Nov 25 13 tcolar Creation

package algo

import ()

// Heap built upon a slice
type Heap struct {
	heap    []int
	maxHeap bool // otherwise minHeap
}

// Heapify the whole heap array
func (h *Heap) Heapify() {
	size := len(h.heap)
	start := (size - 2) / 2
	for start >= 0 {
		percolateDown(h, start, size)
		start -= 1
	}
}

func (h *Heap) Push(item int) {
	h.heap = append(h.heap, item)
	// Percolate up
	cur := len(h.heap) - 1
	parent := (cur - 1) / 2
	for cur >= 0 && h.heap[cur] > h.heap[parent] {
		h.heap[cur], h.heap[parent] = h.heap[parent], h.heap[cur]
		cur = parent
		parent = (cur - 1) / 2
	}
}

// Sift down an item to create a sorted heap array
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
