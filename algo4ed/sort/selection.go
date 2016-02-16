package main

import "fmt"

type SelectionSort struct{}

func (s SelectionSort) Name() string { return "Selection" }

func (s SelectionSort) Sort(data []int, debug bool) {
	for i := 0; i < len(data); i++ {
		var smallest int
		for j := i; j < len(data); j++ {
			if j == i || data[j] < data[smallest] {
				smallest = j
			}
		}
		data[i], data[smallest] = data[smallest], data[i]
		if debug {
			fmt.Printf("sel %d %d %v\n", i, smallest, data)
		}
	}
}
