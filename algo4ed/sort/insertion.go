package main

import "fmt"

type InsertionSort struct{}

func (s InsertionSort) Name() string { return "Insertion" }

func (s InsertionSort) Sort(data []int, debug bool) {
	for i := 1; i < len(data); i++ {
		for j := 0; j < i; j++ {
			if data[i] <= data[j] {
				data[i], data[j] = data[j], data[i]
				if debug {
					fmt.Printf("in %d %d %v\n", i, j, data)
				}
				continue
			}
		}
	}
}
