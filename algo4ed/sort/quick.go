package main

import "fmt"

type QuickSort struct{}

func (s QuickSort) Name() string { return "Quick" }

func (s QuickSort) Sort(data []int, debug bool) {
	s.qsort(data, debug)
	if debug {
		fmt.Printf("qs %v\n", data)
	}
}

// quicksort (in place)
func (s QuickSort) qsort(data []int, debug bool) {
	if len(data) <= 1 {
		return
	}
	p := s.partition(data)
	if debug {
		fmt.Printf("qs %d %v\n", data[p], data)
	}
	s.qsort(data[:p], debug)
	s.qsort(data[p+1:], debug)
}

// In place partition
func (s QuickSort) partition(data []int) int {
	l, r := 0, len(data)-1
	pv := data[r]
	data[r] = 99
	for r > l {
		if data[r-1] > pv {
			data[r] = data[r-1]
			r--
		} else {
			data[r-1], data[l] = data[l], data[r-1]
			l++
		}
	}
	data[r] = pv
	return r
}

// medianPivot to avoid worst cases
func medianPivot(data []int) int {
	left := 0
	right := len(data) - 1
	mid := right / 2
	if data[right] < data[left] {
		data[right], data[left] = data[left], data[right]
	}
	if data[mid] < data[left] {
		data[mid], data[left] = data[left], data[mid]
	}
	if data[right] < data[mid] {
		data[right], data[mid] = data[mid], data[right]
	}
	return mid
}
