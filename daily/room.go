/*
Given an array of time intervals (start, end) for classroom lectures (possibly overlapping), find the minimum number of rooms required.

For example, given [(30, 75), (0, 50), (60, 150)], you should return 2.
*/

package main

import (
	"fmt"
)

func main() {
	intervals := []interval{{30, 75}, {0, 50}, {60, 150}}
	fmt.Println(minRooms(intervals))
}

func minRooms(intervals []interval) int {
	data := map[int]int{}
	rooms := 0
	for _, interval := range intervals {
		for i := interval.start; i <= interval.end; i++ {
			if v, found := data[i]; found {
				data[i] = v + 1
			} else {
				data[i] = 1
			}
			if data[i] > rooms {
				rooms = data[i]
			}
		}
	}
	return rooms
}

type interval struct {
	start int
	end   int
}
