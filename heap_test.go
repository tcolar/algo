// History: Nov 25 13 tcolar Creation

package algo

import (
	"log"
	"testing"
)

func TestHeap(t *testing.T) {
	h := Heap{heap: []int{}}
	validateMaxHeap(h.heap, 0, 9999999999, "")
	h.Push(3)
	h.Push(5)
	h.Push(1)
	validateMaxHeap(h.heap, 0, 9999999999, "")
	h.Push(4)
	h.Push(2)
	h.Push(9)
	validateMaxHeap(h.heap, 0, 9999999999, "")
	log.Print(h.heap)
}

func validateMaxHeap(heap []int, index int, maxVal int, depth string) {
	if len(heap) > index {
		val := heap[index]
		//log.Printf("%s %d", depth, val)
		if val > maxVal {
			log.Fatal("Heap not properly ordered")
			return
		}
		maxVal = val
		validateMaxHeap(heap, (index+1)*2-1, maxVal, depth+"-")
		validateMaxHeap(heap, (index+1)*2, maxVal, depth+"-")
	}
}
