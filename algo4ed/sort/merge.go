package main

import "fmt"

type MergeSort struct{}

func (s MergeSort) Name() string { return "Merge" }

func (s MergeSort) Sort(data []int, debug bool) {
	copy(data, s.sort(data, debug))
}

func (s MergeSort) sort(data []int, debug bool) []int {
	if len(data) <= 1 {
		return data
	}
	cut := len(data) / 2
	left := s.sort(data[:cut], debug)
	right := s.sort(data[cut:], debug)
	merged := merge(left, right)
	if debug {
		fmt.Printf("mg %v\n", merged)
	}
	return merged
}

func merge(left, right []int) []int {
	result := make([]int, 0, len(left)+len(right))
	l := 0
	r := 0
	for l < len(left) || r < len(right) {
		switch {
		case l >= len(left):
			result = append(result, right[r])
			r++
		case r >= len(right):
			result = append(result, left[l])
			l++
		default:
			if left[l] <= right[r] {
				result = append(result, left[l])
				l++
			} else {
				result = append(result, right[r])
				r++
			}
		}
	}
	return result
}
