// History: Nov 25 13 tcolar Creation

package algo

import (
	"log"
	"testing"
)

func TestHeapSort(t *testing.T) {
	slice := []int{7, 2, 5, 6, 4, 9, 3, 1, 8}
	HeapSort(slice)
	for i, v := range slice {
		if i != v-1 {
			log.Fatal("Failed heap sort")
		}
	}
}

func TestQuicksort(t *testing.T) {
	slice := []int{7, 2, 5, 6, 4, 9, 3, 1, 8}
	slice = Quicksort(slice)
	for i, v := range slice {
		if i != v-1 {
			log.Fatal("Failed heap sort")
		}
	}
}
