package main

import (
	"fmt"
	"time"
)

func timeSort(s Sorter, data []int, debug bool) time.Duration {
	d := dup(data, 0, len(data))
	t1 := time.Now()
	s.Sort(d, debug)
	t2 := time.Now()
	return t2.Sub(t1)
}

func checkSort(s Sorter, data []int, debug bool) error {
	d := dup(data, 0, len(data))
	s.Sort(d, debug)
	if len(d) != len(data) {
		return fmt.Errorf("Legth mismatch sorted:%d origina:%d", len(d), len(data))
	}
	for i := 0; i != len(d); i++ {
		if i > 0 && d[i-1] > d[i] {
			return fmt.Errorf("Sort error, found %d@%d and %d@%d", d[i-1], i-1, d[i], i)
		}
	}
	return nil
}

func dup(data []int, from, to int) []int {
	dup := make([]int, to-from)
	copy(dup[:], data[from:to])
	return dup
}

func format(s string, n int) string {
	if len(s) > n {
		return s[:n]
	}
	for len(s) < n {
		s += " "
	}
	return s
}

func formatDur(d time.Duration, n int) string {
	u := uint64(d)
	unit := ""
	switch {
	case u == 0:
		return "0"
	case u < uint64(time.Microsecond):
		unit = "ns"
	case u < uint64(time.Millisecond):
		u /= 1000
		unit = "µs"
	case u < uint64(time.Second):
		u /= 1000000
		unit = "ms"
	case u < uint64(time.Minute):
		u /= 1000000000
		unit = "s"
	default:
		u /= 600000000000
		unit = "m"
	}
	return format(fmt.Sprintf("%d%s", u, unit), n)
}
