package main

import "fmt"

type BubbleSort struct{}

func (s BubbleSort) Name() string { return "Bubble" }

func (s BubbleSort) Sort(data []int, debug bool) {
	end := len(data)
	for {
		swapped := 0
		for i := 1; i < end; i++ {
			if data[i-1] > data[i] {
				data[i-1], data[i] = data[i], data[i-1]
				if debug {
					fmt.Printf("bb %d %d %v\n", i-1, i, data)
				}
				swapped = i
			}
		}
		end = swapped
		if end == 0 {
			return
		}
	}
}
