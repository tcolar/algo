package main

import (
	"fmt"
	"sort"
)

// Retrieve tickets from train windows
// Ticket price is equal to remaining number of tickets
// Input:
//	Line 1 : n m
//	Line 2 : a1 a2 a3 ...
// n->windows
// m->Total tickets to retrieve
// a->tickets of the window (by n index)
// Output : Print max possile revenue
func main() {
	var n, m, a int
	fmt.Scan(&n, &m)
	windows := make([]int, n)
	for i := 0; i != n; i++ {
		fmt.Scan(&a)
		windows[i] = a
	}
	sort.Sort(sort.Reverse(sort.IntSlice(windows)))
	total := revenue(m, windows)
	fmt.Println(total)
}

func revenue(m int, windows []int) int {
	total := 0
	tickets := 0
	msg := ""
outer:
	for {
		for i := 0; i != len(windows)-1; i++ {
			if i > 0 && windows[i] <= windows[i-1] {
				continue outer
			}
			total += windows[i]
			msg = fmt.Sprintf("%s+%d", msg, windows[i])
			windows[i]--
			tickets++
			//fmt.Printf("%d %v : %s = %d\n", windows, tickets, msg, total)
			if tickets == m {
				return total
			}
		}
	}
	return -1
}
